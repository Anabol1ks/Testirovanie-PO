package signinup

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

func SetDB(database *sqlx.DB) {
	db = database
}

var db *sqlx.DB
var jwtKey = []byte("dpokfqkl-023fcs")

type ClientLog struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	Username     string `db:"username"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
