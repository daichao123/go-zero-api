package middleware

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (m *UserMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		fmt.Println("middleware begin ...")
		logx.Info("middleware begin ...")
		// Passthrough to next handler if need
		next(w, r)
		fmt.Println("middleware end ...")
	}
}
