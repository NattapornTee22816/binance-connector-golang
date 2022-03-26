## Binance Connector for Golang (Spot)
<hr/>

- This project guideline from [binance spot](https://binance-docs.github.io/apidocs/spot/en)
- Important ``this project not released``

### How to use

> go get github.com/NattapornTee22816/binance-connector-golang

### Using Binance Spot
document [binance general info](https://binance-docs.github.io/apidocs/spot/en/#general-info)

> ```
> import "github.com/NattapornTee22816/binance-connector-golang/spot"
> ```
> Testnet
> ```
> client := spot.NewTestnetAPI(<api-key>, <api-secret>)
> ```
> Basic
> ```
> client := spot.NewBasicAPI(<api-key>, <api-secret>)
> ```
> Advance
> ```
> client := spot.NewAPI(
>   <api-key>,
>   <api-secret>,<br>
>   url,
>   timeout, // 0 is infinite
>   proxies,
>   logLevel,
>   ctx,
> )
> ```

### Using Websocket
document [Binance Websocket](https://binance-docs.github.io/apidocs/spot/en/#websocket-market-streams)
- websocket auto reconnect when any error
- can use multiple handler on subscription

> New Websocket
> ```
> ws, err := websocket.NewWsStream()
> if err != nil {
>   panic(err)
> }
> ```
> 
> Example for kline stream<br>
> create stream type<br>
> kline stream is ```<symbol>@kline_<interval>```
> 
> ```
> streams := make([]string, 0)<br>
> stream, _ := NewKlineStreamType(symbol, interval)<br>
> streams = append(streams, stream)
> ```      
> 
> ***Subscribe streams***<br>
> subscription must have handler for do something<br>
> can have multiple handler on subscription<br>
> on next subscribe can not send handler
> 
> Example handler
> ```
> func handlerLogData(steam string, data *KlineStream, err error) {
>   // stream is ```<symbol>@kline_<interval>```
>   do something 
> }
> ```
>
> ```
> func handlerOtherProcess(steam string, data *KlineStream, err error) {
>   // stream is ```<symbol>@kline_<interval>```
>   do something 
> }
> ```
> ```
> err = ws.SubscribeKlineStreams(streams, handlerLogData, handlerOtherProcess)<br>
> if err != nil {
>   panic(err)
> }
> ```
> on receive stream data will call handler step by step
> 
> ***Example*** from above:<br>
> has 2 handler<br>
>  1.call handlerLogData<br>
>  2.call handlerOtherProcess
> 
> ***Note:***<br>
> when has error and cannot process next handler, Please set that error to err for stop next handler
> 
> ***# Stop Websocket***
> ```
> ws.Shutdown()
> ```
