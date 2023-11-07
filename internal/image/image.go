package image

import (
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/file"
	"github.com/disism/saikan/ent/image"
	"github.com/disism/saikan/internal/saved"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
)

type Server struct {
	ctx    *gin.Context
	client *ent.Client
}

func NewServer(ctx *gin.Context, client *ent.Client) *Server {
	return &Server{ctx: ctx, client: client}
}

func (s *Server) QueryImage(coverID uint64) (*ent.Image, error) {
	return s.client.Image.Query().Where(image.IDEQ(coverID)).WithFile().Only(s.ctx)
}

type Image struct {
	ID     string      `json:"id"`
	Width  int32       `json:"width"`
	Height int32       `json:"height"`
	File   *saved.File `json:"file"`
}

func (s *Server) AddImages() error {
	defer s.client.Close()

	params := map[string]string{
		"hash":   "file hash is required",
		"name":   "file name is required",
		"size":   "file size is required",
		"height": "image height is required",
		"width":  "image width is required",
	}

	for key, msg := range params {
		if s.ctx.PostForm(key) == "" {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf(msg))
		}
	}

	hash := s.ctx.PostForm("hash")
	name := s.ctx.PostForm("name")

	height, err := strconv.ParseInt(s.ctx.PostForm("height"), 10, 32)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidHeight"))
	}

	width, err := strconv.ParseInt(s.ctx.PostForm("width"), 10, 32)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidWidth"))
	}

	size, err := strconv.ParseUint(s.ctx.PostForm("size"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidImageSize"))
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	f, err := func() (*ent.File, error) {
		query, err := tx.File.Query().Where(file.HashEQ(hash)).Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				create, err := tx.File.
					Create().
					SetName(name).
					SetHash(hash).
					SetSize(size).
					Save(s.ctx)
				if err != nil {
					return nil, fmt.Errorf("create file error: %w", err)
				}
				return create, nil
			}
			return nil, fmt.Errorf("query file error: %w", err)
		}
		return query, nil
	}()
	if err != nil {
		slog.Error("query file error: ", err.Error())
		return err
	}
	i, err := func() (*ent.Image, error) {
		query, err := tx.Image.Query().Where(image.HasFileWith(file.IDEQ(f.ID))).Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				create, err := tx.Image.Create().SetFile(f).SetHeight(int32(height)).SetWidth(int32(width)).Save(s.ctx)
				if err != nil {
					return nil, fmt.Errorf("create image error: %w", err)
				}
				return create, nil
			}
			return nil, fmt.Errorf("query image error: %w", err)
		}
		return query, nil
	}()
	if err != nil {
		slog.Error("query image error: ", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	s.ctx.JSON(http.StatusCreated, &Image{
		ID:     strconv.FormatUint(i.ID, 10),
		Width:  int32(width),
		Height: int32(height),
		File: &saved.File{
			ID:   strconv.FormatUint(f.ID, 10),
			Hash: f.Hash,
			Name: f.Name,
			Size: strconv.FormatUint(f.Size, 10),
		},
	})
	return nil
}
