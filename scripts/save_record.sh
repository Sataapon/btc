# save record
curl -X POST localhost:8080 -d '{
    "datetime": "2019-10-05T14:00:00+00:00",
    "amount": 10.25
}'
curl -X POST localhost:8080 -d '{
    "datetime": "2019-10-05T15:00:00+00:00",
    "amount": 10.25
}'
curl -X POST localhost:8080 -d '{
    "datetime": "2019-10-05T16:00:00+00:00",
    "amount": 10.25
}'