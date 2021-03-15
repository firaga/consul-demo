# consul 注册发现用例
## client consul客户端
## server 服务端
## dnsc 使用dns协议获取服务地址
## todo
实现http server 通过 dns调用http服务

启动consul
consul agent -dev -client 0.0.0.0 -config-dir=/etc/consul.d
启动server
cd grpcServer
go run cmd/main.go
go run 
启动反向代理
go run cmd/http_entrypoint.go
请求:
curl http://localhost:8081/echo/me
curl -X POST http://localhost:8081/echo -H "Content-Type: application/json" -d'{"value": "foo"}'
执行客户端
go run cmd/main.go
执行通过consul获取地址的客户端:
go run cmd/main_consul_client.go


#部分代码来自:
client server
[janlely/consul-go-grpc-demo](https://github.com/janlely/consul-go-grpc-demo)
dnsc
[consul分布式服务注册和发现](https://blog.51cto.com/tianshili/1758566)