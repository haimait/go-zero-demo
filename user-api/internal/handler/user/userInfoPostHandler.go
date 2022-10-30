package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/user-api/internal/logic/user"
	"go-zero-demo/user-api/internal/svc"
	"go-zero-demo/user-api/internal/types"
)

func UserInfoPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoPostRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserInfoPostLogic(r.Context(), svcCtx)
		resp, err := l.UserInfoPost(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
