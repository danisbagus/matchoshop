package domain

import (
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/dgrijalva/jwt-go"
)

const ACCESS_TOKEN_DURATION = time.Hour
const REFRESH_TOKEN_DURATION = time.Hour * 24 * 30

const HMAC_SAMPLE_SECRET = "machoshop-secret"

type User struct {
	UserID    int64  `db:"user_id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	RoleID    int64  `db:"role_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

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

func NewAuthToken(claims AccessTokenClaims) AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return AuthToken{token: token}
}

func JwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		logger.Error("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}

func (r User) ClaimsForAccessToken() AccessTokenClaims {
	return r.claimsForUser()
}

func (r User) claimsForUser() AccessTokenClaims {
	return AccessTokenClaims{
		UserID: r.UserID,
		RoleID: r.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
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

func (r AuthToken) NewAccessToken() (string, *errs.AppError) {
	signedString, err := r.token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		logger.Error("Failed while signing access token: " + err.Error())
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
		logger.Error("Failed while signing refresh token: " + err.Error())
		return "", errs.NewUnexpectedError("cannot generate refresh token")
	}

	return signedString, nil
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
