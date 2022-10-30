package {{.PkgName}}

import (
	"net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/common/response" // ①
    "go-zero-demo/pkg/validator"   // ③
	"go-zero-demo/common/errorx" //④
	"log" // ⑤
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}{{end}}

		if errMsg, errCode := validator.Validate(req); errCode != 0 {
			log.Printf("request params validate failed, err: %s, params: %+v", errMsg, req)
			httpx.Error(w, errorx.NewDefaultError(errMsg))
			return
		} // ⑥

		l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		{{if .HasResp}}response.Response(w, resp, err){{else}}response.Response(w, nil, err){{end}}//②
	}
}
