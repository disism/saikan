package server

import (
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/artists"
	"github.com/disism/saikan/ent/files"
	"github.com/disism/saikan/ent/musics"
	"github.com/disism/saikan/ent/users"
	"github.com/disism/saikan/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Musics interface {
	AddMusics() error
	SaveMusic() error
	ListMusics() error
	GetMusic() error
	EditMusic() error
	DelMusic() error
}

const (
	MusicNotFound       = "music not found"
	MusicExistInLibrary = "music already exists in your library"
)

type Music struct {
	ID          string    `json:"id,omitempty"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	File        *File     `json:"file"`
	Artists     []*Artist `json:"artists"`
	Albums      []*Album  `json:"albums"`
}

func FMTMusic(m *ent.Musics) *Music {
	r := &Music{
		ID:          strconv.FormatUint(m.ID, 10),
		CreateTime:  m.CreateTime,
		UpdateTime:  m.UpdateTime,
		Name:        m.Name,
		Description: m.Description,
	}
	if m.Edges.File != nil {
		r.File = FMTFile(m.Edges.File)
	}
	if m.Edges.Artists != nil {
		for _, artist := range m.Edges.Artists {
			r.Artists = append(r.Artists, FMTArtist(artist))
		}
	}
	if m.Edges.Albums != nil {
		for _, album := range m.Edges.Albums {
			r.Albums = append(r.Albums, AlbumFromEnt(album))
		}
	}
	return r
}

type ExistsMusic struct {
	Name  string  `json:"name"`
	Music []Music `json:"music"`
}

type AddMusicsForm struct {
	Hash string `form:"hash" json:"hash" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
	Size string `form:"size" json:"size" binding:"required"`
}

func (s *Server) AddMusics() error {
	defer s.client.Close()

	var forms []*AddMusicsForm
	if err := s.ctx.ShouldBind(&forms); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	h := make(map[string]bool)
	hashes := make([]string, 0, len(forms))
	creates := make([]*Music, 0, len(forms))
	exists := make([]*ExistsMusic, 0, len(forms))

	for _, form := range forms {
		if _, e := h[form.Hash]; e {
			ErrorBadRequest(s.ctx, fmt.Sprintf("duplicate hash detected: %s", form.Hash))
			return nil
		}
		h[form.Hash] = true
		hashes = append(hashes, form.Hash)

		tx, err := s.client.Tx(s.ctx)
		if err != nil {
			return fmt.Errorf("add musics tx error: %w", err)
		}
		defer tx.Rollback()

		f, err := tx.Files.
			Query().
			Where(files.HashEQ(form.Hash)).
			Only(s.ctx)

		if ent.IsNotFound(err) {
			size, parseErr := strconv.ParseUint(form.Size, 10, 64)
			if parseErr != nil {
				return fmt.Errorf("add music parse size error: %w", parseErr)
			}
			create, createErr := tx.Files.Create().
				SetHash(form.Hash).
				SetSize(size).
				SetName(form.Name).
				Save(s.ctx)
			if createErr != nil {
				return fmt.Errorf("add musics create file error: %w", createErr)
			}
			name := create.Name[:strings.LastIndex(create.Name, ".")]

			music, musicErr := tx.Musics.
				Create().
				AddUserIDs(
					s.GetUserID(),
				).
				SetName(name).
				SetFileID(create.ID).
				Save(s.ctx)
			if musicErr != nil {
				return musicErr
			}
			creates = append(creates, FMTMusic(music))
		} else if err != nil {
			return fmt.Errorf("add musics query file error: %w", err)
		} else {
			all, err := f.QueryMusics().All(s.ctx)
			if err != nil {
				return fmt.Errorf("add musics query existing musics error: %w", err)
			}
			if len(all) == 0 {
				name := f.Name[:strings.LastIndex(f.Name, ".")]

				music, musicErr := tx.Musics.
					Create().
					AddUserIDs(s.GetUserID()).
					SetName(name).
					SetFileID(f.ID).
					Save(s.ctx)
				if musicErr != nil {
					return musicErr
				}
				creates = append(creates, FMTMusic(music))
			} else {
				var musicList []Music
				for _, m := range all {
					musicList = append(musicList, *FMTMusic(m))
				}
				exists = append(exists, &ExistsMusic{
					Name:  all[0].Name,
					Music: musicList,
				})
			}
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("add musics commit tx error: %w", err)
		}
	}

	s.ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"creates": creates,
		"exists":  exists,
	})
	return nil
}

type SaveMusicParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

func (s *Server) SaveMusic() error {
	defer s.client.Close()

	var params SaveMusicParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	music, err := s.client.Musics.
		Query().
		Where(musics.IDEQ(params.ID)).
		WithUsers().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNotFound(s.ctx, MusicNotFound)
			return nil
		}
		return fmt.Errorf("save music query error: %w", err)
	}

	maps := make(map[uint64]struct{})
	for _, r := range music.Edges.Users {
		maps[r.ID] = struct{}{}
	}

	if _, exists := maps[s.GetUserID()]; exists {
		ErrorConflict(s.ctx, MusicExistInLibrary)
		return nil
	}

	if err := s.client.Users.
		Update().
		Where(
			users.IDEQ(
				s.GetUserID(),
			),
		).
		AddMusics(music).
		Exec(s.ctx); err != nil {
		return fmt.Errorf("save music error: %w", err)
	}

	Success(s.ctx, gin.H{"code": http.StatusOK})
	return nil
}

type GetMusicsQuery struct {
	Name string `form:"name"`
}

func (s *Server) ListMusics() error {
	defer s.client.Close()

	query := BuildMusicQuery(s.client)

	var params GetMusicsQuery
	if err := s.ctx.ShouldBindQuery(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}
	if strings.TrimSpace(params.Name) != "" {
		query.Where(musics.NameContainsFold(params.Name))
	} else {
		query.Where(
			musics.HasUsersWith(
				users.IDEQ(
					s.GetUserID(),
				),
			),
		)
	}

	all, err := query.All(s.ctx)
	if err != nil {
		return fmt.Errorf("list musics error: %w", err)
	}

	r := make([]*Music, len(all))
	for i, a := range all {
		r[i] = FMTMusic(a)
	}

	Success(s.ctx, r)

	return nil
}

func BuildMusicQuery(client *ent.Client) *ent.MusicsQuery {
	return client.
		Musics.
		Query().
		WithFile().
		WithArtists().
		WithPlaylists(
			func(query *ent.PlaylistsQuery) {
				query.WithImage(func(query *ent.ImagesQuery) {
					query.WithFile()
				})
			},
		).
		WithAlbums(
			func(query *ent.AlbumsQuery) {
				query.WithImage(func(query *ent.ImagesQuery) {
					query.WithFile()
				})
				query.WithArtists()
			},
		)
}

type GetMusicParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

func (s *Server) GetMusic() error {
	defer s.client.Close()

	var params GetMusicParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	query, err := BuildMusicQuery(s.client).Where(
		musics.IDEQ(params.ID),
		musics.HasUsersWith(
			users.IDEQ(
				s.GetUserID(),
			),
		),
	).Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("get music error: %w", err)
	}

	Success(s.ctx, FMTMusic(query))
	return nil
}

type EditMusicParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

type EditMusicForm struct {
	Name            string   `form:"name" json:"name"`
	Description     string   `form:"description" json:"description"`
	AddArtistIDs    []string `form:"add_artist_ids" json:"add_artist_ids" binding:"omitempty,gt=0,dive,min=1"`
	RemoveArtistIDs []string `form:"remove_artist_ids" json:"remove_artist_ids" binding:"omitempty,gt=0,dive,min=1"`
}

func (f *EditMusicForm) IsEmpty() bool {
	return strings.TrimSpace(f.Name) == "" &&
		strings.TrimSpace(f.Description) == "" &&
		len(f.AddArtistIDs) == 0 &&
		len(f.RemoveArtistIDs) == 0
}

func (s *Server) EditMusic() error {
	defer s.client.Close()

	var params EditMusicParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	var forms EditMusicForm
	if err := s.ctx.ShouldBind(&forms); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	if forms.IsEmpty() {
		ErrorBadRequest(s.ctx, "edit music form must not be empty")
		return nil
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("edit music new tx error: %w", err)
	}
	defer tx.Rollback()

	query, err := BuildMusicQuery(s.client).
		Where(
			musics.IDEQ(params.ID),
		).
		WithUsers().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			s.ctx.JSON(http.StatusNotFound, gin.H{
				"code":  http.StatusNotFound,
				"error": MusicNotFound,
			})
			return nil
		}
		return fmt.Errorf("edit music query error: %w", err)
	}

	edit := tx.Musics.Update().Where(musics.IDEQ(params.ID))

	create := tx.Musics.Create().AddUserIDs(s.GetUserID())

	if forms.Name != "" && forms.Name != query.Name {
		edit.SetName(forms.Name)
		query.Name = forms.Name
	}

	if forms.Description != "" && forms.Description != query.Description {
		edit.SetDescription(forms.Description)
		query.Description = forms.Description
	}

	if len(forms.AddArtistIDs) != 0 || len(forms.RemoveArtistIDs) != 0 {
		adds, err := helper.StrSliceToUint64SliceMap(forms.AddArtistIDs)
		if err != nil {
			return fmt.Errorf("edit music artists parse ids error: %w", err)
		}
		removes, err := helper.StrSliceToUint64SliceMap(forms.RemoveArtistIDs)
		if err != nil {
			return fmt.Errorf("edit music artists parse ids error: %w", err)
		}

		set := make(map[uint64]struct{})
		for _, artistID := range query.Edges.Artists {
			if _, found := removes[artistID.ID]; !found {
				set[artistID.ID] = struct{}{}
			}
		}
		for r := range adds {
			if _, found := removes[r]; !found {
				set[r] = struct{}{}
			}
		}

		f := make([]uint64, 0, len(set))
		for id := range set {
			f = append(f, id)
		}

		all, err := s.client.Artists.
			Query().
			Where(artists.IDIn(f...)).
			All(s.ctx)
		if err != nil {
			return fmt.Errorf("query existing artists error: %w", err)
		}
		edit.ClearArtists().AddArtists(all...)
		query.Edges.Artists = all
	}

	if query.Edges.Users != nil {
		if len(query.Edges.Users) == 1 && query.Edges.Users[0].ID == s.GetUserID() {
			if err := edit.Exec(s.ctx); err != nil {
				return fmt.Errorf("edit music update error: %w", err)
			}
		} else if len(query.Edges.Users) < 1 {
			if err := edit.AddUserIDs(s.GetUserID()).Exec(s.ctx); err != nil {
				return fmt.Errorf("edit music update error: %w", err)
			}
		} else {
			if err := tx.Musics.Update().Where(musics.IDEQ(params.ID)).RemoveUserIDs(s.GetUserID()).Exec(s.ctx); err != nil {
				return fmt.Errorf("edit music remove user error: %w", err)
			}
			if err := create.
				SetName(query.Name).
				SetDescription(query.Description).
				SetFile(query.Edges.File).
				AddArtists(query.Edges.Artists...).
				Exec(s.ctx); err != nil {
				return fmt.Errorf("edit music create error: %w", err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("edit music commit error: %w", err)
	}

	Success(s.ctx, gin.H{"code": http.StatusOK})
	return nil
}

type DelMusicParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

func (s *Server) DelMusic() error {
	defer s.client.Close()

	var params DelMusicParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	query, err := BuildMusicQuery(s.client).Where(
		musics.IDEQ(params.ID),
		musics.HasUsersWith(
			users.IDEQ(
				s.GetUserID(),
			),
		),
	).Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("del music query error: %w", err)
	}

	if err := s.client.Users.Update().Where(
		users.IDEQ(
			s.GetUserID(),
		),
	).RemoveMusics(query).Exec(s.ctx); err != nil {
		return err
	}

	Success(s.ctx, gin.H{"code": http.StatusOK})
	return nil
}
