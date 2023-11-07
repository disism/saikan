package api_server

import (
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func AddMusicsHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add musics new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).AddMusics(); err != nil {
		slog.Error("add musics error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func SaveMusicHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("save music new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).SaveMusic(); err != nil {
		slog.Error("save music error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func ListMusicsHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("list musics new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).ListMusics(); err != nil {
		slog.Error("list musics error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func GetMusicHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("get music new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).GetMusic(); err != nil {
		slog.Error("get music error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func EditMusicHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("edit music new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).EditMusic(); err != nil {
		slog.Error("edit music error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}

func DelMusicHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("del music new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).DelMusic(); err != nil {
		slog.Error("del music error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}
