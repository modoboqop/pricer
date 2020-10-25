# pricer

To start server use docker-compose

```bash
docker-compose up -d
```
MongoDB, 2 replicas of app and nginx will be deployed.

GRPC: localhost:8443

HTTP: localhost:8000
```curl
GET localhost:8000/alive
GET localhost:8000/products - for test
```

For GRPC method UpdatePrices you can use product prices provided itself just send request with payload
```json
{ "url": "http://localhost:8080/products" }
```

GRPC:
UpdatePrices: in response return products parsed from external file. Not from DB

GetProductPrices in response return products from db
