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

type AuthRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) port.IAuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r AuthRepo) FindOne(email string) (*domain.User, *errs.AppError) {
	var login domain.User
	sqlVerify := `SELECT user_id, email, password FROM users WHERE email = $1`

	err := r.db.QueryRow(sqlVerify, email).Scan(&login.UserID, &login.Email, &login.Password)
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

func (r AuthRepo) Verify(token string) *errs.AppError {
	jwtToken, err := jwtTokenFromString(token)
	if err != nil {
		return errs.NewAuthorizationError(err.Error())
	}

	if !jwtToken.Valid {
		return errs.NewAuthorizationError("Invalid token")
	}
	return nil
}

func (r AuthRepo) CreateUserMerchant(data *domain.UserMerchant) (*domain.UserMerchant, *errs.AppError) {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting create new user merchant " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	sqlInsertUser := `INSERT INTO users(email, password, created_at, updated_at) 
					  VALUES($1, $2, $3, $4)
					  RETURNING user_id`

	var userID int64
	err = tx.QueryRow(sqlInsertUser, data.Email, data.Password, data.CreatedAt, data.UpdatedAt).Scan(&userID)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while insert new user: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	sqlInserMerchant := `INSERT INTO merchants(user_id, name, identifier, created_at, updated_at) 
					     VALUES($1, $2, $3, $4, $5)
						 RETURNING merchant_id`

	var merchantID int64
	err = tx.QueryRow(sqlInserMerchant, userID, data.MerchantName, data.MerchantIdentifier, data.CreatedAt, data.UpdatedAt).Scan(&merchantID)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while insert new merchant: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	data.MerchantID = merchantID
	data.UserID = userID

	return data, nil
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
