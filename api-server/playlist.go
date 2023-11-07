package api_server

import (
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func CreatePlaylistHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("create playlist new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).CreatePlaylist(); err != nil {
		slog.Error("create playlist error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func ListPlaylistsHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("list playlists new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).ListPlaylists(); err != nil {
		slog.Error("list playlists error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func GetPlaylistHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("get playlist new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).GetPlaylist(); err != nil {
		slog.Error("get playlist error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func EditPlaylistHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("edit playlist new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).EditPlaylist(); err != nil {
		slog.Error("edit playlist error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func RemovePlaylistImageHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("remove playlist image new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).RemovePlaylistImage(); err != nil {
		slog.Error("remove playlist image error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func AddMusicsToPlaylistHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add musics to playlist new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).AddMusicsToPlaylist(); err != nil {
		slog.Error("add musics to playlist error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func RemoveMusicsFromPlaylistHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("remove musics from playlist new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).RemoveMusicsFromPlaylist(); err != nil {
		slog.Error("remove musics from playlist error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func DelPlaylistHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("del playlist new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).DelPlaylist(); err != nil {
		slog.Error("del playlist error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}
