package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/common/response" // ①
	"go-zero-demo/user-api/internal/logic/user"
	"go-zero-demo/user-api/internal/svc"
	"go-zero-demo/user-api/internal/types"
	"net/http"
)

func UserInfoRpcGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoGetRpcRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserInfoRpcGetLogic(r.Context(), svcCtx)
		resp, err := l.UserInfoRpcGet(&req)
		response.Response(w, resp, err) //②
	}
}
