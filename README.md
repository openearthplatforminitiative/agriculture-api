# agriculture-api
API for agriculture data. This API aggregates data from our other APIs. It gives a summary of the most
useful data for agriculture.


## Running locally
To run the API locally either run with the Dockerfile or follow the steps below.

Building:
```bash
go build -v ./...
```

Running:
```bash
go run ./...
```

You can then access the API at `http://localhost:8080`:

```bash
curl -i -X GET http://0.0.0.0:8080/agriculture/summary?lat=-1.9441&lon=30.0619
```

## Generating openapi spec
First generate the swagger spec using the following command:
```bash
swag init -g ./cmd/app/main.go -o cmd/docs
```


Then run the following command to convert it to openapi spec:

```bash
# call swaggers api to convert the spec
curl -X POST -H "Content-Type: application/json" --data-binary "@cmd/docs/swagger.json" https://converter.swagger.io/api/convert -o "cmd/docs/swagger.json"
```