package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/dgrijalva/jwt-go"
)

const ACCESS_TOKEN_DURATION = time.Hour
const REFRESH_TOKEN_DURATION = time.Hour * 24 * 30
const HMAC_SAMPLE_SECRET = "machoshop-secret"

type AccessTokenClaims struct {
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	TokenType string `json:"token_type"`
	UserID    int64  `json:"user_id"`
	RoleID    int64  `json:"role_id"`
	jwt.StandardClaims
}

type AuthToken struct {
	token *jwt.Token
}

func (r AuthToken) NewAccessToken() (string, *errs.AppError) {
	signedString, err := r.token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		return "", errs.NewUnexpectedError("cannot generate access token")
	}
	return signedString, nil
}

func (r AuthToken) NewRefreshToken() (string, *errs.AppError) {
	claims := r.token.Claims.(AccessTokenClaims)
	refreshClaims := claims.RefreshTokenClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedString, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))

	if err != nil {
		return "", errs.NewUnexpectedError("cannot generate refresh token")
	}
	return signedString, nil
}

func (r AccessTokenClaims) RefreshTokenClaims() RefreshTokenClaims {
	return RefreshTokenClaims{
		TokenType: "refresh_token",
		UserID:    r.UserID,
		RoleID:    r.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(REFRESH_TOKEN_DURATION).Unix(),
		},
	}
}

func (r RefreshTokenClaims) AccessTokenClaims() AccessTokenClaims {
	return AccessTokenClaims{
		UserID: r.UserID,
		RoleID: r.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}

func GenerateAccessTokenAndRefreshToken(userID int64, roleID int64) (string, string, *errs.AppError) {

	claims := AccessTokenClaims{
		UserID: userID,
		RoleID: roleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}

	authToken := NewAuthToken(claims)

	accessToken, appErr := authToken.NewAccessToken()
	if appErr != nil {
		return "", "", appErr
	}

	refreshToken, appErr := authToken.NewRefreshToken()
	if appErr != nil {
		return "", "", appErr
	}

	return accessToken, refreshToken, nil
}

func ValidatedToken(r *http.Request) (*AccessTokenClaims, *errs.AppError) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return nil, errs.NewAuthorizationError("Missing token!")
	}
	tokenString := strings.Split(authHeader, "Bearer ")[1]

	jwtToken, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(HMAC_SAMPLE_SECRET), nil
	})

	if err != nil {
		return nil, errs.NewAuthorizationError(fmt.Sprintf("Error while parsing token: %+v", err.Error()))
	}

	if !jwtToken.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errs.NewAuthorizationError("That's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errs.NewAuthorizationError("Token has expired")
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				return nil, errs.NewAuthorizationError("Invalid token")
			} else {
				return nil, errs.NewAuthorizationError(fmt.Sprintf("Couldn't handle this token: %+v", err.Error()))
			}
		}
	}

	claims := jwtToken.Claims.(*AccessTokenClaims)

	return claims, nil
}

func NewAuthToken(claims AccessTokenClaims) AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return AuthToken{token: token}
}

func NewAccessTokenFromRefreshToken(refreshToken string) (string, *errs.AppError) {
	token, err := jwt.ParseWithClaims(refreshToken, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		return "", errs.NewAuthenticationError("Invalid or expired refresh token")
	}

	claims := token.Claims.(*RefreshTokenClaims)
	accessTokenClaims := claims.AccessTokenClaims()
	authToken := NewAuthToken(accessTokenClaims)

	return authToken.NewAccessToken()
}

func IsTokenValid(token string) *jwt.ValidationError {

	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(HMAC_SAMPLE_SECRET), nil
	})

	if err != nil {
		var validationErr *jwt.ValidationError
		if errors.As(err, &validationErr) {
			return validationErr
		}
	}

	return nil
}
