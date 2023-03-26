package middleware

import (
	"net/http"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/utils/auth"

	"github.com/labstack/echo/v4"
)

func AuthorizationHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := auth.VerifyToken(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": err.Error()})
			}
			return next(c)
		}
	}
}

func ACL(permission map[int]bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userInfo := auth.GetClaimData(c)
			if !permission[int(userInfo.RoleID)] {
				appErr := errs.NewAuthorizationError("Not authorized")
				return c.JSON(appErr.Code, appErr.AsMessage())

			}
			return next(c)
		}
	}
}
