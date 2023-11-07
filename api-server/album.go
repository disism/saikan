package api_server

import (
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func CreateAlbumHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("create albums new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).CreateAlbum(); err != nil {
		slog.Error("create albums error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func GetAlbumsHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("get albums new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).GetAlbums(); err != nil {
		slog.Error("get albums error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func GetAlbumHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("get album new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).GetAlbum(); err != nil {
		slog.Error("get album error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func AddMusicsToAlbumHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add musics to album new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).AddMusicsToAlbum(); err != nil {
		slog.Error("add musics to album error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func RemoveMusicsFromAlbumHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("remove musics from album new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).RemoveMusicsFromAlbum(); err != nil {
		slog.Error("remove musics from album error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func EditAlbumHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("edit album new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).EditAlbum(); err != nil {
		slog.Error("edit album error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func DelAlbumHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("delete album new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).DelAlbum(); err != nil {
		slog.Error("delete album error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

}
