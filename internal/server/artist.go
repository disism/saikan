package server

import (
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/artists"
	"net/http"
	"strconv"
	"strings"
)

type Artists interface {
	SearchArtists() error
	CreateArtist() error
}

const (
	NameIsEmpty = "name is empty"
)

type Artist struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func FMTArtist(a *ent.Artists) *Artist {
	return &Artist{
		ID:   strconv.FormatUint(a.ID, 10),
		Name: a.Name,
	}
}

type CreateArtistForm struct {
	Name string `form:"name" json:"name" binding:"required,startsnotwith= ,endsnotwith= "`
}

func (s *Server) CreateArtist() error {
	defer s.client.Close()

	var forms CreateArtistForm
	if err := s.ctx.ShouldBind(&forms); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}
	create, err := func() (*ent.Artists, error) {
		search, err := s.client.Artists.
			Query().
			Where(
				artists.NameEQ(forms.Name),
			).
			Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				create, err := s.client.Artists.
					Create().
					SetName(forms.Name).
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
		return err
	}

	s.ctx.JSON(http.StatusCreated, FMTArtist(create))
	return nil
}

type SearchArtistsQuery struct {
	Name string `form:"name"`
}

func (s *Server) SearchArtists() error {
	defer s.client.Close()

	var query SearchArtistsQuery
	if err := s.ctx.ShouldBindQuery(&query); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	if strings.TrimSpace(query.Name) == "" {
		ErrorBadRequest(s.ctx, NameIsEmpty)
		return nil
	}

	all, err := s.client.Artists.
		Query().
		Where(
			artists.NameContains(
				query.Name,
			),
		).
		All(s.ctx)
	if err != nil {
		return fmt.Errorf("search artists error: %v", err)
	}

	r := make([]*Artist, len(all))
	for i, a := range all {
		r[i] = FMTArtist(a)
	}

	Success(s.ctx, r)
	return nil
}
