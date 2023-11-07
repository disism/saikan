package api_server

import (
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/saved"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func AddSaved(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add saved new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).AddSaved(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func GetSaved(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("get saved new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).GetSaved(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func GetSaves(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("get saves new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).GetSaves(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func EditSaved(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("edit saved new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).EditSaved(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func DeleteSaved(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("delete saved new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).DeleteSaved(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func MKDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add dir new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).MKDir(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func ListDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("list dirs new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).ListDir(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func ListDirs(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("list dirs new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).ListDirs(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func RenameDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("rename dir new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).RenameDir(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func MVDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("mv dir new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).MVDir(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func RMDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("delete dir new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).RMDir(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func LinkDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("link dir new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).LinkDir(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func UnlinkDir(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("unlink dir new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := saved.NewServer(c, client.Client).UnlinkDir(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}
