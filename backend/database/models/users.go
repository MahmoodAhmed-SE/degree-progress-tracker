package models

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/MahmoodAhmed-SE/degree-progress-tracker/engine"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          int      `json:"id" db:"id"`
	Username    string   `json:"username" db:"username"`
	Phone       string   `json:"phone" db:"phone"`
	Password    string   `json:"password" db:"password"`
	Roles       []string `json:"roles" db:"roles"`
	AuthGroupID int64    `json:"auth_group_id" db:"auth_group_id"`
}

func setTokens(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":       user.ID,
		"username":      user.Username,
		"phone":         user.Phone,
		"iss":           "degree-progress-tracker",
		"aud":           "degree-progress-tracker",
		"roles":         user.Roles,
		"auth_group_id": user.AuthGroupID,
		"exp":           time.Now().Add(72 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, nil
}

func emptyTokens() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(-1 * time.Hour),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, nil
}

func (user *User) validate() error {
	if strings.TrimSpace(user.Username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(user.Password) == "" {
		return errors.New("password is required")
	}
	for _, r := range user.Password {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsSymbol(r) {
			return errors.New("password can only contain letters, numbers and symbols")
		}
	}
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}

var Ctx = context.Background()

func (user *User) create() error {
	if _, err := engine.PgxDB.Exec(Ctx, "INSERT INTO users (username,password,phone) VALUES ($1,$2,$3)", user.Username, user.Password, user.Phone); err != nil {
		return err
	}
	return nil
}

func (user *User) Register() error {
	if err := user.validate(); err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if err := user.create(); err != nil {
		return err
	}

	return nil
}

func (user *User) Login() (string, error) {
	var dbUser User
	err := pgxscan.Get(Ctx, engine.PgxDB, &dbUser, "SELECT id, username, password,phone, auth_group_id FROM users WHERE username=$1", user.Username)
	if err != nil {
		return "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) != nil {
		return "", errors.New("invalid password")
	}
	if err := pgxscan.Select(Ctx, engine.PgxDB, &dbUser.Roles, "SELECT r.name FROM roles_group rg  JOIN roles r ON rg.role_id = r.id WHERE rg.group_id = $1", dbUser.AuthGroupID); err != nil {
		return "", err
	}
	token, err := setTokens(dbUser)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (user *User) GetAll() ([]User, error) {
	var users []User
	if err := pgxscan.Select(Ctx, engine.PgxDB, &users, "SELECT id, username,phone,  auth_group_id FROM users"); err != nil {
		return []User{}, err
	}
	return users, nil
}

func (user *User) Update() error {
	if strings.TrimSpace(user.Username) == "" {
		return errors.New("username is required")
	}
	for _, r := range user.Username {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsSymbol(r) {
			return errors.New("username can only contain letters, numbers and symbols")
		}
	}
	if user.ID == 0 {
		return errors.New("Unauthorized")
	}
	if _, err := engine.PgxDB.Exec(Ctx, "UPDATE users SET username=$1 WHERE id=$2", user.Username, user.ID); err != nil {
		return err
	}

	return nil
}
func (user *User) Logout() (string, error) {
	token, err := emptyTokens()
	if err != nil {
		return "", err
	}
	return token, nil
}
func (user *User) GetByID() error {
	if err := pgxscan.Get(Ctx, engine.PgxDB, user, "SELECT id, username, phone, auth_group_id FROM users WHERE id=$1", user.ID); err != nil {
		return err
	}
	return nil
}
