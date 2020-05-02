# Rest API Call

To run the sample code:

    go run main.go --symbol=${SYMBOL} --period=${TIME_IN_MINUTE}


Crypto.com Exchange's API returns 500 instead 404 when symbol does not exist.
Period flag has to be in minutes and possible values are `[1, 5, 15, 30, 60, 1440, 10080, 43200]` which corresponds to `1min, 5min, 15min, 30min, 1hour, 1day, 1week, 1month`

To obtain the keys, you have to generate it in [API Management](https://crypto.com/exchange/personal/api-management)