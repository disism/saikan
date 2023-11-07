package api_server

import (
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func CreateArtistHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("create artist new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).CreateArtist(); err != nil {
		slog.Error("create artist error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

}

func SearchArtistsHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("search artists new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).SearchArtists(); err != nil {
		slog.Error("search artists error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}
