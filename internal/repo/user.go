package repo

import (
	"database/sql"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
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

func (r UserRepo) Verify(token string) *errs.AppError {
	jwtToken, err := domain.JwtTokenFromString(token)
	if err != nil {
		return errs.NewAuthorizationError(err.Error())
	}

	if !jwtToken.Valid {
		return errs.NewAuthorizationError("Invalid token")
	}
	return nil
}

func (r UserRepo) GenerateAccessTokenAndRefreshToken(data *domain.User) (string, string, *errs.AppError) {

	claims := data.ClaimsForAccessToken()

	authToken := domain.NewAuthToken(claims)

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
