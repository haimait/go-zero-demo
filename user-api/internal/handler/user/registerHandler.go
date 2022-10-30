package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/common/errorx"   //④
	"go-zero-demo/common/response" // ①
	"go-zero-demo/pkg/validator"   // ③
	"go-zero-demo/user-api/internal/logic/user"
	"go-zero-demo/user-api/internal/svc"
	"go-zero-demo/user-api/internal/types"
	"log" // ⑤
	"net/http"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if errMsg, errCode := validator.Validate(req); errCode != 0 {
			log.Printf("request params validate failed, err: %s, params: %+v", errMsg, req)
			httpx.Error(w, errorx.NewDefaultError(errMsg))
			return
		} // ⑥

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		response.Response(w, resp, err) //②
	}
}
