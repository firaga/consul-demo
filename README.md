# consul 注册发现用例

## 文件结构

```
consul-demo/
├── dnsc       // client 通过dns方式服务发现
├── google     // grpc http options 编译依赖
├── grpcClient //client
├── grpcServer //server
├── proto      //protobuff定义文件
```

## 使用

### 编译ptotobuffer文件

```
cd grpcClient
go gernerate ./...
cd ../grpcServer
go gernerate ./...
```

### 启动consul

```
consul agent -dev -client 0.0.0.0 -config-dir=/etc/consul.d
```

### 启动server

```
cd grpcServer
go run cmd/main.go
go run 
go run cmd/http_entrypoint.go  //grpc代理
```

访问

```
curl http://localhost:8081/echo/me
// 配置不一样有不同的访问方式,如果配置为post方式:
// curl -X POST http://localhost:8081/echo -H "Content-Type: application/json" -d'{"value": "foo"}'
```

### 启动client

```
cd grpcClient
go run cmd/main.go //直接访问
go run cmd/main_consul_client.go  //通过consul服务发现访问
```

### dns方式访问(暂未成功)

代码/dnsc可以取到,curl访问未成功,可能需要配置下dns服务器

```
dig @127.0.0.1 -p 8600 echoService.service.consul SRV

这行并没有走dns,只是做了到本地的映射
~~ curl --resolve echoService.service.consul:8081:127.0.0.1  http://echoService.service.consul:8081/echo/hello ~~
```

## 参考

实现分别参考自:

* server && client:  [janlely/consul-go-grpc-demo](https://github.com/janlely/consul-go-grpc-demo)
* dnsc: [consul分布式服务注册和发现](https://blog.51cto.com/tianshili/1758566)