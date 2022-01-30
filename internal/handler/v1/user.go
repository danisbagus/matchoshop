package v1

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/dgrijalva/jwt-go"
)

type AuthHandler struct {
	Service port.IUserService
}

func (rc AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		logger.Error("Error while decoding login request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed to login")
		return
	}

	token, appErr := rc.Service.Login(loginRequest)
	if appErr != nil {
		response.Write(w, appErr.Code, appErr.AsMessage())
		return
	}

	response.Write(w, http.StatusOK, *token)
}

func GetClaimData(r *http.Request) (*domain.AccessTokenClaims, *errs.AppError) {
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer")
	var token string

	if len(splitToken) == 2 {
		token = strings.TrimSpace(splitToken[1])
	} else {
		logger.Error("Error while split token")
		return nil, errs.NewAuthorizationError("Invalid token")
	}

	jwtToken, err := jwt.ParseWithClaims(token, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		logger.Error("Error while parsing token: " + err.Error())
		return nil, errs.NewAuthorizationError(err.Error())
	}

	if !jwtToken.Valid {
		return nil, errs.NewAuthorizationError("Invalid token")
	}

	claims := jwtToken.Claims.(*domain.AccessTokenClaims)
	return claims, nil
}
