
### 1. "获取用户信息"

1. 路由定义

- Url: /user/info
- Method: POST
- Request: `UserInfoRequest`
- Response: `UserInfoResponse`

2. 请求定义


```golang
type UserInfoRequest struct {
	UserId int64 `json:"userId"`
}
```


3. 返回定义


```golang
type UserInfoResponse struct {
	UserId int64 `json:"userId"`
	Nickname string `json:"nickname"`
}
```
  

