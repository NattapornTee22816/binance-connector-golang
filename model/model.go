package model

import (
	"github.com/NattapornTee22816/binance-connector-golang/lib"
	"strings"
)

type Parser struct {
	logger lib.BinanceLogger
	ParserOption
}

type ParserOption struct {
	showWarning bool
}

func NewParser(options ...*ParserOption) *Parser {
	var option *ParserOption
	if len(options) == 0 {
		option = &ParserOption{
			showWarning: true,
		}
	} else {
		option = options[0]
	}

	return &Parser{
		*lib.NewLogger("binance-model-parser", lib.LogLevelTrace),
		*option,
	}
}

func (r *Parser) errorParser(err error) error {
	if err == nil {
		return nil
	}
	message := err.Error()
	// error json-parser
	if strings.HasPrefix(message, "Value is not") {
		return err
	}

	if r.showWarning {
		r.logger.Warn(message)
	}
	return nil
}
