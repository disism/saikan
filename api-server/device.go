package api_server

import (
	"github.com/disism/saikan/internal/authx"
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func GetDevices(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("authn get devices new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}

	if err := authx.NewServer(c, client.Client).GetDevices(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func DeleteDevice(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("authn delete device new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}

	if err := authx.NewServer(c, client.Client).DeleteDevice(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}
