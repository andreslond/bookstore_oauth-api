package app

import (
	"github.com/andrestor2/bookstore_oauth-api/src/clients/cassandra"
	"github.com/andrestor2/bookstore_oauth-api/src/http"
	"github.com/andrestor2/bookstore_oauth-api/src/repository/db"
	"github.com/andrestor2/bookstore_oauth-api/src/repository/rest"
	"github.com/andrestor2/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session := cassandra.GetSessions()
	session.Close()

	atHandler := http.NewHandler(access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token/", atHandler.Create)

	router.Run(":8080")

}
