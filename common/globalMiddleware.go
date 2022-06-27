package common

import (
	"fmt"
	"net/http"
)

type GlobalMiddleware struct {
	Secret string
}

func NewCommonGlobalMiddleware(secret string) *GlobalMiddleware {
	return &GlobalMiddleware{
		Secret: secret,
	}
}

func (m *GlobalMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("this is a global middleware begin ...")
		next(writer, request)
		fmt.Println("this is a global middleware end ...")
		return
	}
}
