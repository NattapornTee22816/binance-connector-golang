package spot

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/NattapornTee22816/binance-connector-golang/lib"
	"github.com/NattapornTee22816/binance-connector-golang/model"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

var (
	VERSION = "v0.0.1-beta"
)

type APIConfig struct {
	// API-key (required)
	key string `validate:"required"`
	// API-secret (required)
	secret string `validate:"required"`
	// base url (optional): the API base url, useful to switch to testnet, etc.
	// By default, it's https://api.binance.com
	baseUrl string
	// the time waiting for server response, number of seconds.
	// By default, no limit
	timeout time.Duration
	// Dictionary mapping protocol to the URL of the proxy. e.g. {'https': 'http://1.2.3.4:8080'}
	proxies map[string]string
}

type API struct {
	APIConfig
	ctx    context.Context
	offset int64
	logger *lib.BinanceLogger
	parser *model.Parser
}

func defaultApiConfig(config *APIConfig) *APIConfig {
	if len(config.baseUrl) == 0 {
		config.baseUrl = "https://api.binance.com"
	}
	if config.timeout < 0 {
		config.timeout = 0
	}
	if &config.proxies == nil {
		config.proxies = make(map[string]string)
	}

	return config
}

func NewBasicAPI(key string, secret string) (*API, error) {
	return NewAPI(
		key,
		secret,
		"https://api.binance.com",
		0,
		nil,
		lib.LogLevelDebug,
		context.Background(),
	)
}

func NewTestnetAPI(key, secret string) (*API, error) {
	return NewAPI(
		key,
		secret,
		"https://testnet.binance.vision",
		0,
		nil,
		lib.LogLevelDebug,
		context.Background(),
	)
}

func NewAPI(
	key string,
	secret string,
	baseUrl string,
	timeout int,
	proxies map[string]string,
	logLevel lib.LogLevel,
	ctx context.Context,
) (*API, error) {
	config := &APIConfig{
		key:     key,
		secret:  secret,
		baseUrl: baseUrl,
		timeout: time.Second * time.Duration(timeout),
		proxies: proxies,
	}
	config = defaultApiConfig(config)

	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		return nil, err
	}

	if ctx == nil {
		ctx = context.Background()
	}

	api := &API{
		*config,
		ctx,
		0,
		lib.NewLogger("binance-connector", logLevel),
		model.NewParser(),
	}

	serverTime, err := api.CheckServerTime()
	if err != nil {
		return nil, err
	}

	api.offset = lib.GetTimestamp(0) - serverTime
	return api, nil
}

// prepareParameters
// check zero value and required field
func (r *API) prepareParameters(params interface{}) url.Values {
	out := url.Values{}

	iVal := reflect.ValueOf(params).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		if f.IsZero() {
			continue
		}

		field := typ.Field(i)
		if keyName, ok := field.Tag.Lookup("param"); ok {
			var v string
			switch f.Interface().(type) {
			case string:
				v = f.String()
			case int, int8, int16, int32, int64:
				v = strconv.FormatInt(f.Int(), 64)
			case float32:
				v = strconv.FormatFloat(f.Float(), 'f', -1, 32)
			case float64:
				v = strconv.FormatFloat(f.Float(), 'f', -1, 64)
			case time.Time:
				v = strconv.FormatInt(f.Interface().(time.Time).UnixMilli()-r.offset, 10)
			default:
				r.logger.Warn(fmt.Sprintf("parameter '%s' type not support", keyName))
			}

			if len(v) > 0 {
				out.Add(keyName, v)
			}
		}
	}

	return out
}

func convertEndpointSecurityType(securityType model.EndpointSecurityType) (bool, bool) {
	apiKey, signature := false, false
	if securityType != model.EndpointSecurityTypeNone {
		apiKey = true
	}
	if securityType == model.EndpointSecurityTypeTrade ||
		securityType == model.EndpointSecurityTypeMargin ||
		securityType == model.EndpointSecurityTypeUserData {
		signature = true
	}
	return apiKey, signature
}

func (r *API) sign(payload string) string {
	mac := hmac.New(sha256.New, []byte(r.secret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}

func (r *API) sendRequest(httpMethod string, urlPath string, payload interface{}, securityType model.EndpointSecurityType) ([]byte, error) {
	defer func() {
		_ = r.logger.Sync()
	}()
	apiKey, signature := convertEndpointSecurityType(securityType)

	queryString := ""
	params := url.Values{}
	if payload != nil && (reflect.ValueOf(payload).Kind() == reflect.Ptr && !reflect.ValueOf(payload).IsNil()) {
		validate := validator.New()
		if err := validate.Struct(payload); err != nil {
			if r.logger.CanDebug() {
				r.logger.Error(err.Error())
			}
			return nil, err
		}

		params = r.prepareParameters(payload)
	}

	if signature {
		params.Add("timestamp", strconv.FormatInt(lib.GetTimestamp(0), 10))
		queryString = params.Encode()
		queryString = fmt.Sprintf("%s&signature=%s", params.Encode(), r.sign(queryString))
	} else if len(params) > 0 {
		queryString = params.Encode()
	}

	header := http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	header.Add("User-Agent", fmt.Sprintf("binance-connector-golang/%s", VERSION))
	if apiKey {
		header.Add("X-MBX-APIKEY", r.key)
	}

	endpoint := fmt.Sprintf("%s%s", r.baseUrl, urlPath)
	response, err := r.dispatchRequest(httpMethod, endpoint, header, queryString)
	if err != nil {
		return nil, err
	}

	byteArray, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if r.logger.CanDebug() {
		r.logger.Debug(fmt.Sprintf("[res] status %v", response.Status))
	}
	if r.logger.CanTrace() {
		r.logger.Debug(fmt.Sprintf("[res] header %v", response.Header))
		r.logger.Debug(fmt.Sprintf("[res] body %v", string(byteArray)))
	}

	if err = r.handleException(response, byteArray); err != nil {
		return nil, err
	}

	return byteArray, nil
}

func (r *API) dispatchRequest(httpMethod string, endpoint string, header http.Header, payload string) (*http.Response, error) {
	transport := &http.Transport{}
	if proxy, ok := r.proxies["https"]; ok {
		urlProxy, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(urlProxy)
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   r.timeout,
	}

	req, err := http.NewRequest(httpMethod, endpoint, nil)
	if err != nil {
		r.logger.Error("unable to request request")
		return nil, err
	}
	req.WithContext(r.ctx)

	if len(header) > 0 {
		req.Header = header
	}
	if len(payload) > 0 {
		req.URL.RawQuery = payload
	}

	if r.logger.CanDebug() {
		r.logger.Debug(fmt.Sprintf("[req] %s %s", httpMethod, req.URL.String()))
	}
	if r.logger.CanTrace() {
		r.logger.Debug(fmt.Sprintf("[req][header] %v", req.Header))
	}

	return client.Do(req)
}

func (r *API) handleException(response *http.Response, body []byte) error {
	statusCode := response.StatusCode
	if statusCode < 400 {
		return nil
	}

	if statusCode >= 400 && statusCode < 500 {
		type BinanceError struct {
			Code int64  `json:"code"`
			Msg  string `json:"msg"`
		}

		binanceError := new(BinanceError)
		if err := json.Unmarshal(body, &binanceError); err != nil {
			return &ClientError{
				StatusCode:   int64(statusCode),
				ErrorCode:    binanceError.Code,
				ErrorMessage: binanceError.Msg,
				Header:       response.Header,
			}
		} else {
			return &ClientError{
				StatusCode:   int64(statusCode),
				ErrorCode:    0,
				ErrorMessage: string(body),
				Header:       response.Header,
			}
		}
	}
	return &ServerError{
		StatusCode: int64(statusCode),
		Message:    string(body),
	}
}
