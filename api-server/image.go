package api_server

import (
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/image"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func AddImages(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add images new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := image.NewServer(c, client.Client).AddImages(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}
