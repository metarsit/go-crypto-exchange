# [Crypto.com Exchange](https://crypto.com/exchange)
Crypto Exchange is a Crypto Trading platform that has a ready API for consumer. This repository uses Golang to call the provided REST API

## `Supported API` ([Official Docs](https://crypto.com/exchange-doc#endpoint)):
The table of listed APIs that are supported by this package.

### `User API`
| User API | Support |
:---------------- | :---------------- |
/v1/account | :heavy_check_mark:
/v1/order |
/v1/showOrders |
/v1/orders |
/v1/cancelAllOrders |
/v1/openOrders |
/v1/allOrders |
/v1/myTrade |


### `Market API`
| Endpoint API | Support |
:---------------- | :---------------- |
/v1/symbols | :heavy_check_mark:
/v1/ticker |
/v1/ticker |
/v1/klines |
/v1/trades |
/v1/ticker/price |
/v1/depth | :heavy_check_mark:

## `Installation`
To use this package, run:

    go get github.com/metarsit/exchange