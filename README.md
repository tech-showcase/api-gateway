## API GATEWAY

### Description
This repo contains project that act as an **api-gateway**.
This service is part of a big system. 
The whole system will be used to present **micro-services without an orchestrator**.

### Features
- Forward request from a client to corresponding service.
- Translate http request/response to gRPC request/response.

### API
Please refer to all proto file [here](proto) for more detail about the provided API.
You can use gRPC by:
- Installing [protoc](http://google.github.io/proto-lens/installing-protoc.html)
- Installing [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/)
- Generating code by executing `protoc -I./proto --go_out=plugins=grpc:. proto/*/*.proto`

### How to run
#### Docker
- Install docker
- Create following environment variable and fill it with the right value
```shell script
  CONSUL_ADDRESS=http://consul-server-address
  ENTERTAINMENT_SERVICE_ADDRESS=entertainment-service-address
```
`CONSUL_ADDRESS` is currently not used cause the service discovery feature is under maintenance
- Build and run docker image as below
```shell script
$ docker build -t api-gateway .
$ docker run -p 8081:8080 api-gateway
```

### Tech / Dependency
- [Go kit - service](https://github.com/go-kit/kit)
- [gRPC - api](https://grpc.io/)
- [Gorilla mux - api](https://github.com/gorilla/mux)