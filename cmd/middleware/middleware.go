package middleware

import (
	"net/http"
)

type IAPIMiddleware interface {
	Authorization() func(http.Handler) http.Handler
	ACL(permission map[int]bool) func(http.Handler) http.Handler
}

type APIMiddleware struct {
}

func NewAPIMiddleware() IAPIMiddleware {
	return &APIMiddleware{}
}
