package repo

import (
	"database/sql"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

const ACCESS_TOKEN_DURATION = time.Hour
const dbTSLayout = "2006-01-02 15:04:05"

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) port.IUserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r UserRepo) FindOne(email string) (*domain.User, *errs.AppError) {
	var login domain.User
	sqlVerify := `SELECT user_id, email, password, role_id FROM users WHERE email = $1`

	err := r.db.QueryRow(sqlVerify, email).Scan(&login.UserID, &login.Email, &login.Password, &login.RoleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewAuthenticationError("invalid credentials")
		} else {
			logger.Error("Error while verifying login request from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &login, nil
}

func (r UserRepo) GenerateAndSaveRefreshTokenToStore(authToken *domain.AuthToken) (string, *errs.AppError) {
	// generate ther refresh token
	refreshToken, appErr := generateRefreshToken(authToken)
	if appErr != nil {
		return "", appErr
	}

	// store it to stroe
	sqlInsert := `INSERT INTO refresh_token_stores(refresh_token, created_at) 
	VALUES($1, $2)`

	currentTime := time.Now().Format(dbTSLayout)

	_, err := r.db.Exec(sqlInsert, refreshToken, currentTime)
	if err != nil {
		logger.Error("Error while insert refresh token: " + err.Error())
		return "", errs.NewUnexpectedError("Unexpected database error")
	}

	return refreshToken, nil
}

func (r UserRepo) Verify(token string) *errs.AppError {
	jwtToken, err := jwtTokenFromString(token)
	if err != nil {
		return errs.NewAuthorizationError(err.Error())
	}

	if !jwtToken.Valid {
		return errs.NewAuthorizationError("Invalid token")
	}
	return nil
}

func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		logger.Error("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}

func generateRefreshToken(authToken *domain.AuthToken) (string, *errs.AppError) {

	refreshToken, appErr := authToken.NewRefreshToken()
	if appErr != nil {
		return "", appErr
	}

	return refreshToken, nil
}
