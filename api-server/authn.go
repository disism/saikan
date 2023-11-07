package api_server

import (
	"github.com/disism/saikan/internal/authx"
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func AuthN(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("authn new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := authx.NewServer(c, client.Client).AuthN(); err != nil {
		server.NewContext(c).ErrorUnauthorized()
	}
}

func Login(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("login new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := authx.NewServer(c, client.Client).Login(); err != nil {
		server.NewContext(c).ErrorUnauthorized()
	}
}
