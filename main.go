package main

import (
	"log"
	"net/http"
	signinup "webser/signInUp"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Open("postgres", "host=localhost port=5433 user=postgres password=12341 dbname=auth sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	signinup.SetDB(db)

	router := gin.Default()
	router.LoadHTMLGlob("*.html")
	router.StaticFS("/web", http.Dir("web"))
	router.GET("/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signin.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", gin.H{
			"title": "Main website",
		})
	})

	router.POST("/sign-up", signinup.SignUp)
	router.POST("/sign-in", signinup.SignIn)

	router.Run(":8888")
}
