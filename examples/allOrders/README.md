# Rest API Call

To run the sample code:

    go run main.go --api=${API_KEY} --secret=${SECRET_KEY} --symbol=${SYMBOL} --startDate=${START_DATE} --endDate=${END_DATE} --page=${PAGE} --pageSize=${PAGE_SIZE}


Date format has to be in `yyyy-MM-dd HH:mm:ss` i.e. `2020-05-09 00:00:00`

To obtain the keys, you have to generate it in [API Management](https://crypto.com/exchange/personal/api-management)