## 安装
```shell
# 安装
go install github.com/zeromicro/go-zero/tools/goctl@latest

# 查看版本号
goctl -v
```

## 创建服务
```shell
# 创建模块
goctl api new core

# 运行
go run core.go -f etc/core-api.yaml

```

## api生成
```shell
# core.api中写入路由
service user-api {
	@handler UserHandler
	get /user/login(LoginRequest) returns (LoginReply)
}

type LoginRequest {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginReply {
	Token string `json:"token"`
}

# 生成api
goctl api go -api core.api -dir . -style go_zero
```


```shell
# 中间件
@server(
	middleware: Auth
)
service core-api {
  
}
```