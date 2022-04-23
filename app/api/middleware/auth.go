package middleware

import (
	"context"
	"net/http"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/matchoshop/utils/auth"
)

func AuthorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, appErr := auth.ValidatedToken(r)
			if appErr != nil {
				response.Write(w, appErr.Code, appErr.AsMessage())
				return
			} else {
				ctx := context.WithValue(r.Context(), "userInfo", claims)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			}
		})
	}
}

func ACL(permission map[int]bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)
			if !permission[int(userInfo.RoleID)] {
				appErr := errs.NewAuthorizationError("Not authorized")
				response.Write(w, appErr.Code, appErr.AsMessage())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
