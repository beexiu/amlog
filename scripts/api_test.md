# API

## API Curl Test
```
curl -i -k -3 -c tmp/cookies "http://localhost:26031/login?namename=brant&password=1234"
curl -i -k -3 -b tmp/cookies "http://localhost:26031/record?miles=12345&cash=12.34"


curl -i -k -3 "http://localhost:26031/record?miles=12354&cash=12.34"
```