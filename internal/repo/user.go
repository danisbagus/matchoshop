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

func NewUserRepo(db *sqlx.DB) port.UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r UserRepo) FindOne(email string) (*domain.User, *errs.AppError) {
	var login domain.User
	sqlVerify := `SELECT user_id, email, password, name, role_id FROM users WHERE email = $1`

	err := r.db.QueryRow(sqlVerify, email).Scan(&login.UserID, &login.Email, &login.Password, &login.Name, &login.RoleID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while verifying login request from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
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

func (r UserRepo) CreateUserCustomer(data *domain.User) (*domain.User, *errs.AppError) {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting create new user customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	sqlInsert := `INSERT INTO users(email, password, name, role_id,  created_at, updated_at)
				  VALUES($1, $2, $3, $4, $5, $6)
				  RETURNING user_id`

	var userID int64
	err = tx.QueryRow(sqlInsert, data.Email, data.Password, data.Name, data.RoleID, data.CreatedAt, data.UpdatedAt).Scan(&userID)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while create new user: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	data.UserID = userID

	return data, nil
}
