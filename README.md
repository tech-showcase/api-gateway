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
- Create `config-dev.json` under `config` dir which contains following content
```json
{
  "consul_address": "http://consul-server:8500",
  "entertainment_service_address": "localhost:8082"
}
```
`consul_address` is currently not used cause the related feature is under maintenance
- Build and run docker image as below
```shell script
$ docker build -t api-gateway .
$ docker run -p 8081:8080 api-gateway
```
