package api_server

import (
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/nodeinfo"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func GetNodeInfo(c *gin.Context) {

	client, err := database.New(c)
	if err != nil {
		slog.Error("get saved new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}

	oidc, err := nodeinfo.NewServer(c, client.Client).GetAllOIDC()
	if err != nil {
		server.NewContext(c).ErrorInternalServer()
	}

	c.JSON(200, gin.H{
		"saikan":                  "https://saikan.disism.com",
		"oidc_provider_supported": oidc,
	})
}
