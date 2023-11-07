package api_server

import (
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func ListDevicesHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("list devices new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).ListDevices(); err != nil {
		slog.Error("list devices error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func DeleteDeviceHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("delete device new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).DeleteDevice(); err != nil {
		slog.Error("delete device error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}
