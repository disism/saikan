package api_server

import (
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func AddImageHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add images new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
	if err := server.NewServer(c, client.Client).AddImage(); err != nil {
		slog.Error("add images error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}
