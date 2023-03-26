package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
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

func VerifyToken(c echo.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}

	_, err = parseToken(token)
	if err != nil {
		return err
	}

	return nil
}

func GetClaimData(c echo.Context) *AccessTokenClaims {
	token, err := getToken(c)
	if err != nil {
		return nil
	}

	jwtToken, err := parseToken(token)
	if err != nil {
		return nil
	}

	userClaims := jwtToken.Claims.(*AccessTokenClaims)

	return userClaims
}

func getToken(c echo.Context) (string, error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return "", errors.New("token not found")
	}

	splitToken := strings.Split(token, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return "", errors.New("invalid token format")
	}

	return strings.TrimSpace(splitToken[1]), nil
}

func parseToken(token string) (*jwt.Token, error) {
	jwtSecret := []byte(HMAC_SAMPLE_SECRET) // todo: move to config
	jwtToken, err := jwt.ParseWithClaims(token, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error while parsing token: %+v", err)
	}

	if !jwtToken.Valid {
		return nil, errors.New("invalid token")
	}

	return jwtToken, nil
}
