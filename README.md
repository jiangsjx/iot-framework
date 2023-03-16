# iot-framework

基于go-zero的IoT应用框架。

### 编译运行

```shell
go mod tidy

cd device && go build && ./device -f etc/device-api.yml

cd server && go build && ./server -f etc/server-api.yml
```
