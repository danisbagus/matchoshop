package middleware

import (
	"net/http"
	"strings"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/matchoshop/internal/core/port"
)

type AuthMiddleware struct {
	UserRepo port.UserRepo
}

func (rc AuthMiddleware) AuthorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader != "" {
				token := getTokenFromHeader(authHeader)
				err := rc.UserRepo.Verify(token)
				if err != nil {
					response.Write(w, err.Code, err.AsMessage())
				} else {
					next.ServeHTTP(w, r)
				}

			} else {
				err := errs.NewAuthorizationError("Missing token!")
				response.Write(w, err.Code, err.AsMessage())

			}
		})
	}
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
