package api_server

import (
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

//func AuthN(c *gin.Context) {
//	client, err := database.New(c)
//	if err != nil {
//		slog.Error("authn new database error: ", err.Error())
//		server.ErrorInternalServerError(c)
//	}
//	if err := server.NewServer(c, client.Client).AuthN(); err != nil {
//		server.ErrorUnauthorized(c, err.Error())
//	}
//}

func LoginHandler(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("login new database error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

	if err := server.NewServer(c, client.Client).Login(); err != nil {
		slog.Error("login error: ", err.Error())
		server.ErrorInternalServerError(c)
	}

}
