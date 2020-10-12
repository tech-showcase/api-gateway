docker exec consul-client \
  /bin/sh -c \
  "echo '$(cat deployment/consul/entertainment-service.json)' > /consul/config/entertainment-service.json"
docker exec consul-client \
  /bin/sh -c \
  "echo '$(cat deployment/consul/entertainment-service-2.json)' > /consul/config/entertainment-service-2.json"

docker exec consul-client \
  consul reload