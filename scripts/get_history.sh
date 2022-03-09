# get history
curl -X GET localhost:8080 -d '{
    "startDatetime": "2019-10-05T14:00:00+00:00",
    "endDatetime": "2019-10-05T16:00:00+00:00"
}'
curl -X GET localhost:8080 -d '{
    "startDatetime": "2019-10-05T14:01:00+00:00",
    "endDatetime": "2019-10-05T17:00:01+00:00"
}'