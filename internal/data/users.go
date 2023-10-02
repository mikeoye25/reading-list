package data

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel struct {
	DB *sql.DB
}

func (u UserModel) Insert(user *User) error {
	query := `
		INSERT INTO dbo.users (name, username, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

	args := []interface{}{user.Name, user.Username, user.Email, user.Password}
	// return the auto generated system values to Go object
	return u.DB.QueryRow(query, args...).Scan(&user.ID, &user.CreatedAt)
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
