package music

import (
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/artist"
	"github.com/disism/saikan/internal/errors"
	"github.com/disism/saikan/internal/server"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
)

type Artist struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (s *Server) CreateArtist() error {
	defer s.client.Close()

	name := s.ctx.PostForm("name")

	query, err := func() (*ent.Artist, error) {
		search, err := s.client.Artist.
			Query().
			Where(artist.NameEQ(name)).
			Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				create, err := s.client.Artist.
					Create().
					SetName(name).
					Save(s.ctx)
				if err != nil {
					return nil, fmt.Errorf("create artist create error: %v", err)
				}
				return create, nil
			}
			return nil, fmt.Errorf("create artist query error: %v", err)
		}
		return search, nil
	}()
	if err != nil {
		slog.Error("create artist error: ", err.Error())
		return err
	}

	r := &Artist{
		ID:   strconv.FormatUint(query.ID, 10),
		Name: query.Name,
	}

	s.ctx.JSON(http.StatusOK, r)
	return nil
}

func (s *Server) SearchArtist() error {
	defer s.client.Close()

	name := s.ctx.Query("name")
	if name == "" {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf(errors.NameIsEmpty))
	}

	all, err := s.client.Artist.Query().
		Where(
			artist.NameContains(
				name,
			),
		).All(s.ctx)
	if err != nil {
		slog.Error("query artist query error: ", err.Error())
		return err
	}

	r := make([]*Artist, len(all))
	for i, a := range all {
		r[i] = &Artist{
			ID:   strconv.FormatUint(a.ID, 10),
			Name: a.Name,
		}
	}

	s.ctx.JSON(http.StatusOK, r)
	return nil
}

func (s *Server) QueryArtist(artistID uint64) (*ent.Artist, error) {
	return s.client.Artist.Get(s.ctx, artistID)
}
