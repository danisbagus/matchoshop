package repository

import (
	"database/sql"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/domain"
	"github.com/jmoiron/sqlx"
)

const ACCESS_TOKEN_DURATION = time.Hour
const dbTSLayout = "2006-01-02 15:04:05"

type IUserRepository interface {
	GetAll() ([]domain.UserDetail, *errs.AppError)
	FindOne(email string) (*domain.UserModel, *errs.AppError)
	FindOneById(userID int64) (*domain.UserModel, *errs.AppError)
	CreateUserCustomer(data *domain.UserModel) (*domain.UserModel, *errs.AppError)
	Update(userID int64, data *domain.UserModel) *errs.AppError
	Delete(userID int64) *errs.AppError
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) GetAll() ([]domain.UserDetail, *errs.AppError) {
	sqlGet := `SELECT u.user_id, u.email, u.name, u.role_id, r.name AS role_name FROM users u
			   INNER JOIN roles r ON r.role_id = u.role_id
			   ORDER BY u.user_id`

	rows, err := r.db.Query(sqlGet)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while fetch all user from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	defer rows.Close()

	users := make([]domain.UserDetail, 0)

	for rows.Next() {
		var user domain.UserDetail
		if err := rows.Scan(&user.UserID, &user.Email, &user.Name, &user.RoleID, &user.RoleName); err != nil {
			logger.Error("Error while scanning product category from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		users = append(users, user)
	}

	return users, nil
}

func (r UserRepository) FindOne(email string) (*domain.UserModel, *errs.AppError) {
	var login domain.UserModel
	sqlVerify := `SELECT user_id, email, password, name, role_id FROM users WHERE email = $1`

	err := r.db.QueryRow(sqlVerify, email).Scan(&login.UserID, &login.Email, &login.Password, &login.Name, &login.RoleID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while verifying login request from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &login, nil
}

func (r UserRepository) FindOneById(userID int64) (*domain.UserModel, *errs.AppError) {
	var login domain.UserModel
	sqlVerify := `SELECT user_id, email, password, name, role_id FROM users WHERE user_id = $1`

	err := r.db.QueryRow(sqlVerify, userID).Scan(&login.UserID, &login.Email, &login.Password, &login.Name, &login.RoleID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while verifying login request from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &login, nil
}

func (r UserRepository) CreateUserCustomer(data *domain.UserModel) (*domain.UserModel, *errs.AppError) {

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

func (r UserRepository) Update(userID int64, data *domain.UserModel) *errs.AppError {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting update user: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlUpdate := `
	UPDATE users 
	SET name = $2, 
		updated_at = $3
	WHERE user_id = $1`

	_, err = tx.Exec(sqlUpdate, userID, data.Name, data.UpdatedAt)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while update user: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}

func (r UserRepository) Delete(userID int64) *errs.AppError {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting delete user: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlDelete := `
	DELETE FROM users 
	WHERE user_id = $1`

	_, err = tx.Exec(sqlDelete, userID)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while delete user: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}
	return nil
}
