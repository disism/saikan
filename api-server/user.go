package api_server

import (
	"github.com/disism/saikan/internal/authx"
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func CreateUser(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("create user new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := authx.NewServer(c, client.Client).CreateUser(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}
