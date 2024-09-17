package signinup

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var newUser ClientLog
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Хз что за ошибка"})
		return
	}

	var existingUser User
	err := db.Get(&existingUser, "SELECT email FROM users WHERE email=$1", newUser.Email)
	if err == nil && existingUser.Email != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "Пользователь с таким email уже зарегистрирован"})
		return
	}

	hashPas, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось хешировать"})
		return
	}
	_, err = db.Exec("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)", newUser.Username, newUser.Email, string(hashPas))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось зарегистрироваться"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь зарегистрирован"})

}
