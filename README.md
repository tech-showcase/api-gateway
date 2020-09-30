## API GATEWAY

### Description
This repo contains project that act as a api-gateway of microservices system.
This service is part of a big system. 
The whole system will be used to present technology show case.

### Features
- Forward request from client to corresponding service.
- Translate http request response to grpc request response.

### How to run
#### Docker
- Install docker
- Create following environment variable and fill it with the right value
```shell script
  CONSUL_ADDRESS=http://consul-server-address,
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