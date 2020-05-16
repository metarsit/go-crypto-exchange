# [Crypto.com Exchange](https://crypto.com/exchange)

Crypto Exchange is a Crypto Trading platform that has a ready API for consumer. This repository uses Golang to call the provided REST API

## `Supported API` ([Official Docs](https://crypto.com/exchange-doc#endpoint)):

The table of listed APIs that are supported by this package.

### `User API`

| User API | Support |
:---------------- | :----------------: |
/v1/account | :heavy_check_mark:
/v1/order |
/v1/showOrders | :heavy_check_mark:
/v1/orders/cancel | :heavy_check_mark:
/v1/cancelAllOrders | :heavy_check_mark:
/v1/openOrders | :heavy_check_mark:
/v1/allOrders | :heavy_check_mark:
/v1/myTrade | :heavy_check_mark:

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

    go get github.com/metarsit/exchange

## Bazel

Below are some useful commands for building and generally fiddling with bazel.

- To build all targets: `bazel build //...`
- To run all tests, building if necessary: `bazel test //...`
- To run all tests, building if necessary, showing verbose output, even when passing: `bazel test --test_output=all --test_arg=-test.v //...`
- To show all known targets: `bazel query ...`
- To automatcially update bazel's go dependencies from the `go.mod` file: `bazel run //:gazelle -- update-repos -from_file=go.mod`
- To automatically generate BUILD.bazel targets for go source code: `bazel run //:gazelle`
