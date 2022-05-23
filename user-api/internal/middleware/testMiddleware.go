package middleware

import (
	"fmt"
	"net/http"
)

type TestMiddleware struct {
}

func NewTestMiddleware() *TestMiddleware {
	return &TestMiddleware{}
}

func (m *TestMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("TestMiddleware...before")
		next(w, r)
		fmt.Println("TestMiddleware...after")
	}
}
