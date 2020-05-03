# [Crypto.com Exchange](https://crypto.com/exchange)
Crypto Exchange is a Crypto Trading platform that has a ready API for consumer. This repository uses Golang to call the provided REST API

## `Supported API` ([Official Docs](https://crypto.com/exchange-doc#endpoint)):
The table of listed APIs that are supported by this package.

### `User API`
| User API | Support |
:---------------- | :----------------: |
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
:---------------- | :----------------: |
/v1/symbols | :heavy_check_mark:
/v1/ticker | :heavy_check_mark:
/v1/ticker?symbol= | :heavy_exclamation_mark:
/v1/klines | :heavy_check_mark:
/v1/trades | :heavy_check_mark:
/v1/ticker/price | :heavy_check_mark:
/v1/depth | :heavy_check_mark:

### `Remark`
There are more information in `README.md` of each examples.

:heavy_check_mark: - API works as expected

:heavy_exclamation_mark: - API does not fully return expected information

## `Installation`
To use this package, run:

    go get github.com/metarsit/go-crypto-exchange
