package saved

import (
	"encoding/json"
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/dir"
	"github.com/disism/saikan/ent/file"
	"github.com/disism/saikan/ent/saved"
	"github.com/disism/saikan/ent/user"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Saved struct {
	ID         string    `json:"id,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
	Name       string    `json:"name,omitempty"`
	Caption    string    `json:"caption,omitempty"`
	File       *File     `json:"file,omitempty"`
}

type File struct {
	ID   string `json:"id,omitempty"`
	Hash string `json:"hash,omitempty"`
	Name string `json:"name,omitempty"`
	Size string `json:"size,omitempty"`
}

const (
	DefaultRootDirName = "root"
)

type Server struct {
	ctx    *gin.Context
	client *ent.Client
}

func NewServer(ctx *gin.Context, client *ent.Client) *Server {
	return &Server{ctx: ctx, client: client}
}

type AddSavedFileData struct {
	Hash    string `json:"hash"`
	Name    string `json:"name"`
	Size    string `json:"size"`
	Caption string `json:"caption"`
}

func (s *Server) AddSaved() error {
	defer s.client.Close()

	id := s.ctx.Query("dir_id")

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var exists []*ent.Saved
	var creates []*ent.Saved

	for _, fs := range s.ctx.PostFormArray("files") {
		var fd *AddSavedFileData
		if err := json.Unmarshal([]byte(fs), &fd); err != nil {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("problems with the expected files"))
		}
		f, err := func() (*ent.File, error) {
			qf, err := tx.File.
				Query().
				Where(
					file.HashEQ(fd.Hash),
				).Only(s.ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					parse, err := strconv.ParseUint(fd.Size, 10, 64)
					if err != nil {
						return nil, err
					}
					cf, err := tx.File.
						Create().
						SetHash(fd.Hash).
						SetName(fd.Name).
						SetSize(parse).
						Save(s.ctx)
					if err != nil {
						slog.Error("create file error: ", err.Error())
						return nil, err
					}
					return cf, nil
				}
				slog.Error("query file error: ", err.Error())
				return nil, err
			}
			return qf, nil
		}()
		if err != nil {
			return err
		}

		if err := func() error {
			qs, err := tx.Saved.
				Query().
				Where(
					saved.HasFileWith(
						file.IDEQ(f.ID),
					),
					saved.HasOwnerWith(
						user.IDEQ(server.NewContext(s.ctx).GetUserID()),
					),
				).
				WithDir().
				Only(s.ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					c, err := tx.
						Saved.
						Create().
						SetFile(f).
						SetName(fd.Name[:strings.LastIndex(fd.Name, ".")]).
						SetCaption(fd.Caption).
						SetOwnerID(
							server.NewContext(s.ctx).GetUserID(),
						).
						Save(s.ctx)
					if err != nil {
						if ent.IsValidationError(err) {
							return err
						}
						slog.Error("create saved error: ", err.Error())
						return err
					}
					l, err := func() (*ent.Dir, error) {
						if strings.TrimSpace(id) == "" {
							qd, err := tx.Dir.
								Query().
								Where(
									dir.HasOwnerWith(
										user.IDEQ(
											server.NewContext(s.ctx).GetUserID(),
										),
									),
									dir.NameEQ(DefaultRootDirName),
								).Only(s.ctx)
							if err != nil {
								if ent.IsNotFound(err) {
									cd, err := tx.Dir.
										Create().
										SetName(DefaultRootDirName).
										SetOwnerID(
											server.NewContext(s.ctx).GetUserID(),
										).
										Save(s.ctx)
									if err != nil {
										slog.Error("create dir error: ", err.Error())
										return nil, err
									}
									return cd, err
								}
								slog.Error("query dir error: ", err.Error())
								return nil, err
							}
							return qd, nil
						}
						parser, err := strconv.ParseUint(id, 10, 64)
						if err != nil {
							return nil, err
						}
						if err != nil {
							return nil, err
						}
						only, err := tx.Dir.Query().Where(
							dir.HasOwnerWith(
								user.IDEQ(
									server.NewContext(s.ctx).GetUserID(),
								),
							),
							dir.IDEQ(parser),
						).Only(s.ctx)
						if err != nil {
							return nil, err
						}
						return only, nil
					}()

					if err := tx.Saved.
						Update().
						Where(
							saved.HasOwnerWith(
								user.IDEQ(
									server.NewContext(s.ctx).GetUserID(),
								),
							),
						).
						AddDir(l).
						Exec(s.ctx); err != nil {
						slog.Error("update saved error: ", err.Error())
						return err
					}
					creates = append(creates, c)
					return nil
				}
				slog.Error("query saved error: ", err.Error())
				return err
			}
			exists = append(exists, qs)
			return nil
		}(); err != nil {
			return server.NewContext(s.ctx).ErrorBadRequest(err)
		}
	}

	if err := tx.Commit(); err != nil {
		slog.Error("saved create commit error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"exists":  exists,
		"creates": creates,
	})
	return nil
}

func (s *Server) GetSaved() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid id"))
	}
	query, err := s.client.Saved.
		Query().
		Where(
			saved.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
			saved.IDEQ(id),
		).
		WithFile().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("get saved query saved error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, query)
	return nil
}

func (s *Server) GetSaves() error {
	defer s.client.Close()
	query, err := s.client.Saved.
		Query().
		Where(
			saved.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		WithFile().
		All(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("get saves query saved error: ", err.Error())
		return err
	}

	r := make([]Saved, len(query))
	for i, q := range query {
		r[i] = Saved{
			ID:         strconv.FormatUint(q.ID, 10),
			CreateTime: q.CreateTime,
			UpdateTime: q.UpdateTime,
			Name:       q.Name,
			Caption:    q.Caption,
			File: &File{
				ID:   strconv.FormatUint(q.Edges.File.ID, 10),
				Hash: q.Edges.File.Hash,
				Name: q.Edges.File.Name,
				Size: strconv.FormatUint(q.Edges.File.Size, 10),
			},
		}
	}
	s.ctx.JSON(http.StatusOK, r)
	return nil
}

func (s *Server) EditSaved() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid id"))
	}
	query, err := s.client.Saved.
		Query().
		Where(
			saved.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
			saved.IDEQ(id),
		).
		WithFile().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("edit saved query saved error: ", err.Error())
		return err
	}
	edit := s.client.Saved.Update().Where(saved.IDEQ(query.ID))
	caption := s.ctx.PostForm("caption")
	switch {
	case caption != "":
		edit.SetCaption(s.ctx.PostForm("caption"))
		query.Caption = caption
	}
	if err := edit.Exec(s.ctx); err != nil {
		slog.Error("edit saved error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, query)
	return nil
}

func (s *Server) DeleteSaved() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid id"))
	}
	query, err := s.client.Saved.
		Query().
		Where(
			saved.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
			saved.IDEQ(id),
		).
		WithFile().
		WithDir().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("delete saved query saved error: ", err.Error())
		return err
	}

	// TODO - remove all link dirs.
	fmt.Println(query.Edges.Dir)

	if err := s.client.Saved.
		Update().
		Where(
			saved.IDEQ(query.ID),
		).
		RemoveDir(query.Edges.Dir...).
		Exec(s.ctx); err != nil {
		slog.Error("delete saved, remove all dirs error: ", err.Error())
		return err
	}
	if err := s.client.Saved.
		DeleteOne(query).
		Exec(s.ctx); err != nil {
		slog.Error("delete saved error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) LinkDir() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid saved id"))
	}

	di := s.ctx.PostForm("dir_id")
	if strings.TrimSpace(di) == "" {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid saved dir_id"))
	}
	dirID, err := strconv.ParseUint(di, 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid dir id"))
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		slog.Error("link dir tx error: ", err.Error())
		return err
	}
	defer tx.Rollback()

	qs, err := tx.Saved.
		Query().
		Where(
			saved.IDEQ(id),
			saved.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		WithDir().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("link dir query saved error: ", err.Error())
		return err
	}

	qd, err := tx.Dir.Query().Where(
		dir.HasOwnerWith(
			user.IDEQ(
				server.NewContext(s.ctx).GetUserID(),
			),
		),
		dir.IDEQ(dirID),
	).Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("link dir query dir error: ", err.Error())
		return err
	}

	if qs.Edges.Dir != nil {
		for _, r := range qs.Edges.Dir {
			if r.ID == qd.ID {
				return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("dir already linked"))
			}
		}
	}

	if err := tx.Saved.
		Update().
		Where(
			saved.IDEQ(id),
		).
		AddDir(qd).
		Exec(s.ctx); err != nil {
		slog.Error("link dir error: ", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		slog.Error("link dir commit error: ", err.Error())
		return err
	}
	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) UnlinkDir() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid saved id"))
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		slog.Error("unlink dir tx error: ", err.Error())
		return err
	}

	defer tx.Rollback()

	qs, err := tx.Saved.
		Query().
		Where(
			saved.IDEQ(id),
			saved.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		WithDir().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("unlink dir query saved error: ", err.Error())
		return err
	}

	qdOptions := tx.Dir.Query().Where(dir.HasOwnerWith(
		user.IDEQ(
			server.NewContext(s.ctx).GetUserID(),
		),
	))

	dirID, exists := s.ctx.GetPostForm("dir_id")
	if exists {
		parse, err := strconv.ParseUint(dirID, 10, 64)
		if err != nil {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid dir id"))
		}
		qdOptions.Where(dir.IDEQ(parse))
	} else {
		qdOptions.Where(dir.NameEQ(DefaultRootDirName))
	}

	qd, err := qdOptions.Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("unlink dir query dir error: ", err.Error())
		return err
	}

	if qs.Edges.Dir != nil {
		found := func() bool {
			for _, r := range qs.Edges.Dir {
				if r.ID == qd.ID {
					return true
				}
			}
			return false
		}()

		if !found {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("dir not linked"))
		}
	}

	if err := tx.Saved.
		Update().
		Where(
			saved.IDEQ(id),
		).
		RemoveDir(qd).
		Exec(s.ctx); err != nil {
		slog.Error("unlink dir error: ", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		slog.Error("unlink dir commit error: ", err.Error())
		return err
	}
	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})

	return nil
}
