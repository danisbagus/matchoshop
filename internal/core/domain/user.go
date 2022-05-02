package domain

type User struct {
	UserID    int64  `db:"user_id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Name      string `db:"name"`
	RoleID    int64  `db:"role_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type UserDetail struct {
	User
	RoleName string `db:"role_name"`
}
