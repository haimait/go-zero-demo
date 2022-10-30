
## 生成api代码命令
在 `user-api/api` 目录中执行下面的命令
> cd user-api/api
> 
> 默认模板生成
> alias apigen="goctl api go -api *.api -dir ../  --style=goZero"
>
> 指定生成模板
> alias apigen="goctl api go -api *.api -dir ../  --style=goZero --home=../../goctlTpl" 

## 生成rpc代码命令

在 `user-rpc/pb` 目录中执行下面的命令
> cd user-rpc/pb
> 
> 默认模板生成
> alias rpcgen="goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero"
> 
> 指定生成模板 
> goctl template init 会生成在~/.goctl/goctl版本号/handler.tpl 目录里 ，默认使用用户家目标里的模板
> 可以指定模板目标 --home=../goctlTpl  这里写*.proto文件的相对路径
> alias rpcgen="goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero --home=../../goctlTpl"

windows下 `*` 号改成具体的文件名，如user.proto

## 生成model层

在项目根目录：

> 执行 `./genModel`


## 接口文档 apipost

> https://console-docs.apipost.cn/preview/8053007db38f9ca5/d53cd71f5e46e2f4
