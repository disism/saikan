package saved

import (
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/dir"
	"github.com/disism/saikan/ent/user"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Dir struct {
	ID         string    `json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Name       string    `json:"name"`
	Subdirs    []*Subdir `json:"subdirs"`
	Saves      []*Saved  `json:"saves"`
}

type Subdir struct {
	ID         string    `json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Name       string    `json:"name"`
}

func (s *Server) MKDir() error {
	defer s.client.Close()

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	id := s.ctx.Query("dir_id")
	name := s.ctx.PostForm("name")

	if strings.TrimSpace(name) == "" {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid dir name"))
	}
	query := tx.Dir.
		Query().
		Where(
			dir.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		)

	r, err := func() (*ent.Dir, error) {
		if strings.TrimSpace(id) != "" {
			i, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				return nil, server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid dir id"))
			}
			q, err := query.Where(dir.IDEQ(i)).WithSubdir().Only(s.ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					return nil, server.NewContext(s.ctx).ErrorNoPermission()
				}
				slog.Error("mk dir query dir error: ", err.Error())
				return nil, err
			}
			return q, nil
		} else {
			n, err := query.Where(dir.NameEQ(DefaultRootDirName)).WithSubdir().Only(s.ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					create, err := tx.Dir.Create().SetOwnerID(server.NewContext(s.ctx).GetUserID()).SetName(DefaultRootDirName).Save(s.ctx)
					if err != nil {
						return nil, err
					}
					return create, nil
				}
				slog.Error("mk dir query root dir error: ", err.Error())
				return nil, err
			}
			return n, nil
		}
	}()
	if err != nil {
		return err
	}
	if r.Edges.Subdir != nil {
		for _, d := range r.Edges.Subdir {
			if d.Name == name {
				return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("dir exists"))
			}
		}
	}

	if strings.EqualFold(DefaultRootDirName, name) {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("root dir creation is not allowed"))
	}

	create, err := tx.Dir.
		Create().
		SetOwnerID(
			server.NewContext(s.ctx).GetUserID(),
		).
		SetName(name).
		AddPdir(r).
		Save(s.ctx)
	if err != nil {
		slog.Error("add dir create dir error: ", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	s.ctx.JSON(http.StatusOK, &Dir{
		ID:         strconv.FormatUint(create.ID, 10),
		CreateTime: create.CreateTime,
		UpdateTime: create.UpdateTime,
		Name:       create.Name,
		Subdirs:    nil,
	})
	return nil
}

func (s *Server) ListDir() error {
	defer s.client.Close()

	id := s.ctx.Query("dir_id")

	options := s.client.Dir.
		Query().
		Where(
			dir.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		)

	query, err := func() (*ent.Dir, error) {
		if strings.TrimSpace(id) != "" {
			i, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				return nil, server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid dir id"))
			}
			q, err := options.Where(dir.IDEQ(i)).
				WithSubdir().
				WithSaves(
					func(q *ent.SavedQuery) {
						q.WithFile()
					},
				).
				WithPdir().
				Only(s.ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					return nil, server.NewContext(s.ctx).ErrorNoPermission()
				}
				slog.Error("mk dir query dir error: ", err.Error())
				return nil, err
			}
			return q, nil
		} else {
			n, err := options.
				Where(dir.NameEQ(DefaultRootDirName)).
				WithSaves(
					func(q *ent.SavedQuery) {
						q.WithFile()
					},
				).
				WithSubdir().
				Only(s.ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					create, err := s.client.Dir.Create().SetOwnerID(server.NewContext(s.ctx).GetUserID()).SetName(DefaultRootDirName).Save(s.ctx)
					if err != nil {
						return nil, err
					}
					return create, nil
				}
				slog.Error("mk dir query root dir error: ", err.Error())
				return nil, err
			}
			return n, nil
		}
	}()
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf("not found"))
		}
		slog.Error("list dir query dir error: ", err.Error())
		return err
	}
	r := &Dir{
		ID:         strconv.FormatUint(query.ID, 10),
		CreateTime: query.CreateTime,
		UpdateTime: query.UpdateTime,
		Name:       query.Name,
	}
	if query.Edges.Subdir != nil {
		subdirs := make([]*Subdir, len(query.Edges.Subdir))
		for i, sub := range query.Edges.Subdir {
			subdirs[i] = &Subdir{
				ID:         strconv.FormatUint(sub.ID, 10),
				CreateTime: sub.CreateTime,
				UpdateTime: sub.UpdateTime,
				Name:       sub.Name,
			}
		}
		r.Subdirs = subdirs
	}
	if query.Edges.Saves != nil {
		saves := make([]*Saved, len(query.Edges.Saves))
		for i, sav := range query.Edges.Saves {
			saves[i] = &Saved{
				ID:         strconv.FormatUint(sav.ID, 10),
				CreateTime: sav.CreateTime,
				UpdateTime: sav.UpdateTime,
				Name:       sav.Name,
				Caption:    sav.Caption,
				File: &File{
					ID:   strconv.FormatUint(sav.Edges.File.ID, 10),
					Hash: sav.Edges.File.Hash,
					Name: sav.Edges.File.Name,
					Size: strconv.FormatUint(sav.Edges.File.Size, 10),
				},
			}
		}
		r.Saves = saves
	}
	s.ctx.JSON(http.StatusOK, r)
	return nil
}

func (s *Server) ListDirs() error {
	defer s.client.Close()
	query, err := s.client.Dir.
		Query().
		Where(
			dir.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).All(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf("not found"))
		}
		slog.Error("list dirs query dir error: ", err.Error())
		return err
	}

	r := make([]*Dir, len(query))
	for i, d := range query {
		r[i] = &Dir{
			ID:         strconv.FormatUint(d.ID, 10),
			CreateTime: d.CreateTime,
			UpdateTime: d.UpdateTime,
			Name:       d.Name,
		}
	}
	s.ctx.JSON(http.StatusOK, r)
	return nil
}

func (s *Server) RenameDir() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid id"))
	}
	name := s.ctx.PostForm("name")

	if strings.EqualFold(name, DefaultRootDirName) {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("root dir renaming is not allowed"))
	}
	query, err := s.client.Dir.
		Query().
		Where(
			dir.IDEQ(id),
			dir.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("rename dir query dir error: ", err.Error())
		return err
	}
	// If the directory name already exists at the same level, it is not allowed to be modified.
	subs, err := query.
		QueryPdir().
		WithSubdir().
		Only(s.ctx)
	if err != nil {
		return err
	}
	if subs != nil {
		for _, d := range subs.Edges.Subdir {
			if d.Name == name {
				return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("dir exists"))
			}
		}
	}

	if err := s.client.Dir.Update().Where(dir.IDEQ(query.ID)).SetName(name).Exec(s.ctx); err != nil {
		slog.Error("rename dir update dir error: ", err.Error())
		return err
	}
	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) MVDir() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid id"))
	}
	// new dir
	di, err := strconv.ParseUint(s.ctx.Param("dir_id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid new dir id"))
	}

	if id == di {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("its own subdirectory"))
	}
	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query, err := tx.Dir.
		Query().
		Where(
			dir.IDEQ(id),
			dir.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		WithPdir().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("mv dir query dir error: ", err.Error())
		return err
	}

	nquery, err := tx.Dir.
		Query().
		Where(
			dir.IDEQ(di),
			dir.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		WithSubdir().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("mv dir query new dir error: ", err.Error())
		return err
	}
	if nquery.Edges.Subdir != nil {
		for _, n := range nquery.Edges.Subdir {
			if n.ID == di {
				return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("dir exists"))
			}
		}
	}
	if query.Edges.Pdir != nil {
		for _, n := range query.Edges.Pdir {
			if n.ID == di {
				return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("dir exists"))
			}
		}
	}

	if err := tx.Dir.
		Update().
		Where(
			dir.IDEQ(query.ID),
		).
		RemovePdir(query.Edges.Pdir...).
		AddPdir(nquery).
		Exec(s.ctx); err != nil {
		slog.Error("mv dir update dir error: ", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		slog.Error("mv dir commit error: ", err.Error())
		return err
	}
	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) RMDir() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("invalid id"))
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		slog.Error("remove dir tx error: ", err.Error())
		return err
	}
	defer tx.Rollback()
	query, err := s.client.Dir.
		Query().
		Where(
			dir.IDEQ(id),
			dir.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		WithSubdir().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("remove dir query dir error: ", err.Error())
		return err
	}
	if query.Name == DefaultRootDirName {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("root dir deletion is not allowed"))
	}

	// TODO REMOVE ALL LINK FILES:Recursively remove all subdirectories with SAVED links

	remove := make([]*ent.Dir, 0)

	var re func(query *ent.Dir) error
	re = func(query *ent.Dir) error {
		for _, sub := range query.Edges.Subdir {
			remove = append(remove, sub)
			f, err := tx.Dir.Query().Where(dir.IDEQ(sub.ID)).WithSubdir().Only(s.ctx)
			if err != nil {
				return err
			}
			if err := re(f); err != nil {
				return err
			}
		}
		return nil
	}
	if err := re(query); err != nil {
		slog.Error("remove dir delete dir error: ", err.Error())
		return err
	}

	remove = append(remove, query)
	for _, sub := range remove {
		if err := tx.Dir.DeleteOne(sub).Exec(s.ctx); err != nil {
			slog.Error("remove dir delete dir error: ", err.Error())
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		slog.Error("remove dir commit error: ", err.Error())
		return err
	}
	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}
