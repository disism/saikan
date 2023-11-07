package api_server

import (
	"encoding/json"
	"github.com/disism/saikan/conf"
	"github.com/disism/saikan/internal/ipfs"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
)

type FileResponse struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
	Size string `json:"size"`
}

func AddFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	f := form.File["path"]

	local := ipfs.Localhost{
		BaseURL: conf.GetIPFSAddr(),
	}

	var fs []FileResponse
	for _, file := range f {
		dst := "/tmp/" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			slog.Error("failed to save file", err)
			server.NewContext(c).ErrorInternalServer()
		}

		r, err := local.Post(dst)
		if err != nil {
			slog.Error("failed to post ipfs", err)
			server.NewContext(c).ErrorInternalServer()
		}

		if r.IsError() {
			c.JSON(r.StatusCode(), gin.H{"errors": r.Status()})
			return
		}

		var fr FileResponse
		if err := json.Unmarshal(r.Body(), &fr); err != nil {
			slog.Error("failed to parse response body", err)
			server.NewContext(c).ErrorInternalServer()
		}
		fs = append(fs, fr)

		if err := os.Remove(dst); err != nil {
			slog.Error("failed to remove file", err)
			server.NewContext(c).ErrorInternalServer()
		}
	}
	c.JSON(http.StatusOK, fs)
}
