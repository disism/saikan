package api_server

import (
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func IPFSAddFilesHandler(c *gin.Context) {
	if err := server.NewServer(c, nil).IPFSAddFiles(); err != nil {
		slog.Error("ipfs add files error: ", err.Error())
		server.ErrorInternalServerError(c)
	}
}
