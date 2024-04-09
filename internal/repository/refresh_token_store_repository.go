package repository

import (
	"database/sql"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/jmoiron/sqlx"
)

type IRefreshTokenStoreRepository interface {
	Insert(refreshToken string) *errs.AppError
	CheckRefreshToken(refreshToken string) (bool, *errs.AppError)
}

type RefreshTokenStoreRepository struct {
	db *sqlx.DB
}

func NewRefreshTokenStoreRepository(db *sqlx.DB) *RefreshTokenStoreRepository {
	return &RefreshTokenStoreRepository{
		db: db,
	}
}

func (r RefreshTokenStoreRepository) Insert(refreshToken string) *errs.AppError {

	sqlInsert := `INSERT INTO refresh_token_stores(refresh_token, created_at) 
		VALUES($1, $2)`

	currentTime := time.Now().Format(dbTSLayout)

	_, err := r.db.Exec(sqlInsert, refreshToken, currentTime)
	if err != nil {
		logger.Error("Error while insert refresh token: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}

func (r RefreshTokenStoreRepository) CheckRefreshToken(refreshToken string) (bool, *errs.AppError) {

	sqlCountRefreshToken := `SELECT COUNT(refresh_token) 
	FROM refresh_token_stores 
	WHERE refresh_token = $1`

	var totalData int64
	err := r.db.QueryRow(sqlCountRefreshToken, refreshToken).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count refresh token from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
}
