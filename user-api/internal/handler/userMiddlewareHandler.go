package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-test/user-api/internal/logic"
	"go-zero-test/user-api/internal/svc"
	"go-zero-test/user-api/internal/types"
)

func userMiddlewareHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserMiddlewareLogic(r.Context(), svcCtx)
		resp, err := l.UserMiddleware(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
