# iot-framework

基于go-zero的IoT应用框架。

### 编译运行

```shell
go mod tidy

cd device && go build && ./device -f etc/device-api.yml

cd server && go build && ./server -f etc/server-api.yml
```

### Swagger

```shell
cd device && docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/device.json -v $PWD:/foo swaggerapi/swagger-ui

cd server && docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/server.json -v $PWD:/foo swaggerapi/swagger-ui
```
