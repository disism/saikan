package music

import (
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/cover"
	"github.com/disism/saikan/ent/file"
	"github.com/disism/saikan/internal/server"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
)

func (s *Server) QueryCover(coverID uint64) (*ent.Cover, error) {
	return s.client.Cover.Query().Where(cover.IDEQ(coverID)).WithFile().Only(s.ctx)
}

// AddCover adds a cover to the Music API.
//
// This function requires the following parameters:
// - cid: file cid is required
// - name: file name is required
// - size: file size is required
// - height: cover height is required
// - width: cover width is required
//
// It returns an error.
func (s *Server) AddCover() error {
	defer s.client.Close()

	params := map[string]string{
		"cid":    "file cid is required",
		"name":   "file name is required",
		"size":   "file size is required",
		"height": "cover height is required",
		"width":  "cover width is required",
	}

	for key, msg := range params {
		if s.ctx.PostForm(key) == "" {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf(msg))
		}
	}

	cid := s.ctx.PostForm("cid")
	name := s.ctx.PostForm("name")
	size := s.ctx.PostForm("size")
	height := s.ctx.PostForm("height")
	width := s.ctx.PostForm("width")

	siz, err := strconv.ParseUint(size, 10, 64)
	if err != nil {
		return err
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	create, err := func() (*ent.Cover, error) {
		f, err := func() (*ent.File, error) {
			query, err := tx.File.Query().Where(file.CidEQ(cid)).Only(s.ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					create, err := tx.File.
						Create().
						SetName(name).
						SetCid(cid).
						SetSize(siz).
						Save(s.ctx)
					if err != nil {
						slog.Error("create file error: ", err.Error())
						return nil, err
					}
					return create, nil
				}
				slog.Error("query file error: ", err.Error())
				return nil, err
			}
			return query, nil
		}()
		if err != nil {
			return nil, err
		}
		query, err := tx.
			Cover.
			Query().
			Where(
				cover.HasFileWith(
					file.IDEQ(f.ID),
				),
			).
			Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				h, err := strconv.ParseInt(height, 10, 32)
				if err != nil {
					return nil, err
				}
				w, err := strconv.ParseInt(width, 10, 32)
				if err != nil {
					return nil, err
				}
				create, err := tx.Cover.Create().SetFile(f).SetHeight(int32(h)).SetWidth(int32(w)).Save(s.ctx)
				if err != nil {
					return nil, err
				}
				return create, nil
			}
			return nil, err
		}
		return query, nil
	}()
	if err != nil {
		return server.NewContext(s.ctx).ErrorBadRequest(err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	s.ctx.JSON(http.StatusCreated, create)
	return nil
}
