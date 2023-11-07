package server

import (
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/albums"
	"github.com/disism/saikan/ent/artists"
	"github.com/disism/saikan/ent/images"
	"github.com/disism/saikan/ent/musics"
	"github.com/disism/saikan/ent/users"
	"github.com/disism/saikan/internal/helper"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strconv"
)

type Albums interface {
	CreateAlbum() error
	GetAlbums() error
	GetAlbum() error
	EditAlbum() error
	AddMusicsToAlbum() error
	RemoveMusicsFromAlbum() error
	DelAlbum() error
}

const (
	ErrorAlbumAlreadyExists = "album already exists"
	ErrorAlbumNotFound      = "album not found"
)

type Album struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        string    `json:"date"`
	Image       *Image    `json:"image"`
	Musics      []*Music  `json:"musics"`
	Artists     []*Artist `json:"artists"`
}

func AlbumFromEnt(a *ent.Albums) *Album {
	r := &Album{
		ID:          strconv.FormatUint(a.ID, 10),
		Title:       a.Title,
		Description: a.Description,
		Date:        a.Date,
		Image:       FMTImage(a.Edges.Image),
	}
	if a.Edges.Musics != nil {
		for _, music := range a.Edges.Musics {
			r.Musics = append(r.Musics, FMTMusic(music))
		}
	}
	if a.Edges.Artists != nil {
		for _, artist := range a.Edges.Artists {
			r.Artists = append(r.Artists, FMTArtist(artist))
		}
	}
	return r
}

type CreateAlbumForm struct {
	Title       string   `form:"title" json:"title" binding:"required,startsnotwith= ,endsnotwith= "`
	Date        string   `form:"date" json:"date" binding:"required,min=4,numeric"`
	ArtistIDs   []string `form:"artist_ids" json:"artist_ids" binding:"required,gt=0,dive,min=1,numeric,excludes= "`
	CoverID     string   `form:"cover_id" json:"cover_id" binding:"required,numeric"`
	Description string   `form:"description" json:"description" binding:"omitempty"`
}

func (s *Server) CreateAlbum() error {
	defer s.client.Close()

	var forms CreateAlbumForm
	if err := s.ctx.ShouldBind(&forms); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	exist, err := s.client.Albums.
		Query().
		Where(
			albums.TitleEQ(forms.Title),
			albums.HasUsersWith(
				users.IDEQ(
					s.GetUserID(),
				),
			),
		).
		Exist(s.ctx)
	if err != nil {
		return fmt.Errorf("check album existence error: %w", err)
	}
	if exist {
		ErrorConflict(s.ctx, ErrorAlbumAlreadyExists)
		return nil
	}

	a, err := func(ids []string) ([]uint64, error) {
		var r []uint64
		for _, id := range ids {
			uid, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				return nil, err
			}
			r = append(r, uid)
		}
		return r, nil
	}(forms.ArtistIDs)
	if err != nil {
		return fmt.Errorf("create album artists parse ids error: %w", err)
	}

	// Start error group for concurrent queries
	g, ctx := errgroup.WithContext(s.ctx)

	var aa []*ent.Artists
	g.Go(func() error {
		var err error
		aa, err = s.client.Artists.
			Query().
			Where(
				artists.IDIn(a...),
			).
			All(ctx)
		return err
	})

	cid, err := strconv.ParseUint(forms.CoverID, 10, 64)
	if err != nil {
		return fmt.Errorf("create albums parse cover id error: %w", err)
	}

	var cover *ent.Images
	g.Go(func() error {
		var err error
		cover, err = s.client.Images.Get(ctx, cid)
		return err
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("create albums query error: %w", err)
	}

	tx, err := s.client.Debug().Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("create albums tx error: %w", err)
	}
	defer tx.Rollback()

	if err := tx.Albums.
		Create().
		SetTitle(forms.Title).
		SetDescription(forms.Description).
		SetDate(forms.Date).
		AddArtists(aa...).
		SetImage(cover).
		AddUserIDs(s.GetUserID()).
		Exec(s.ctx); err != nil { // Use s.ctx here as well
		return fmt.Errorf("create albums save error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("create albums commit error: %w", err)
	}

	s.Ok()
	return nil
}

func BuildAlbumQuery(client *ent.Client) *ent.AlbumsQuery {
	return client.
		Albums.
		Query().
		WithArtists().
		WithImage(func(query *ent.ImagesQuery) {
			query.WithFile()
		}).
		WithMusics(func(query *ent.MusicsQuery) {
			query.WithFile()
			query.WithArtists()
		})
}

func (s *Server) GetAlbums() error {
	defer s.client.Close()

	all, err := s.client.Albums.
		Query().
		Where(
			albums.HasUsersWith(
				users.IDEQ(
					s.GetUserID(),
				),
			),
		).
		WithArtists().
		WithImage(func(query *ent.ImagesQuery) {
			query.WithFile()
		}).
		All(s.ctx)
	if err != nil {
		return fmt.Errorf("get albums error: %w", err)
	}
	r := make([]*Album, len(all))
	for i, a := range all {
		r[i] = AlbumFromEnt(a)
	}

	s.ctx.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"albums": r,
	})
	return nil
}

type GetAlbumParams struct {
	ID uint64 `uri:"id" binding:"required"`
}

func (s *Server) GetAlbum() error {
	defer s.client.Close()

	var params GetAlbumParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return err
	}

	a, err := BuildAlbumQuery(s.client).
		Where(
			albums.IDEQ(params.ID),
		).
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNotFound(s.ctx, ErrorAlbumNotFound)
			return nil
		}
		return fmt.Errorf("get album error: %w", err)
	}

	s.ctx.JSON(http.StatusOK, AlbumFromEnt(a))
	return nil
}

type AddMusicsToAlbumParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

type AddMusicsToAlbumForm struct {
	MusicIDs []string `form:"music_ids" json:"music_ids" binding:"required,gt=0,dive,min=1,numeric,excludes= "`
}

func (s *Server) AddMusicsToAlbum() error {
	defer s.client.Close()

	var params AddMusicsToAlbumParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	var forms AddMusicsToAlbumForm
	if err := s.ctx.ShouldBind(&forms); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	query, err := BuildAlbumQuery(s.client).
		Where(
			albums.IDEQ(
				params.ID,
			),
		).
		WithUsers().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNotFound(s.ctx, ErrorAlbumNotFound)
			return nil
		}
		return fmt.Errorf("add musics to album query error: %w", err)
	}

	c, err := helper.StrSliceToUint64Slice(forms.MusicIDs)
	if err != nil {
		return fmt.Errorf("add musics to album convert error: %w", err)
	}

	all, err := s.client.Musics.
		Query().
		Where(
			musics.IDIn(c...),
			musics.HasUsersWith(
				users.IDEQ(
					s.GetUserID(),
				),
			),
		).
		All(s.ctx)
	if err != nil {
		return fmt.Errorf("add musics to album query error: %w", err)
	}
	mer := append(query.Edges.Musics, all...)
	unique := make(map[uint64]*ent.Musics)
	for _, r := range mer {
		unique[r.ID] = r
	}

	final := make([]*ent.Musics, 0, len(unique))
	for _, artist := range unique {
		final = append(final, artist)
	}

	edit := s.client.Albums.Update().Where(albums.IDEQ(query.ID))
	if query.Edges.Users != nil {
		if len(query.Edges.Users) == 1 && query.Edges.Users[0].ID == s.GetUserID() {
			if err := edit.ClearMusics().AddMusics(all...).Exec(s.ctx); err != nil {
				return fmt.Errorf("add musics to album error: %w", err)
			}
		} else if len(query.Edges.Users) < 1 {
			if err := edit.AddUserIDs(s.GetUserID()).ClearMusics().AddMusics(final...).Exec(s.ctx); err != nil {
				return fmt.Errorf("add musics to album create error: %w", err)
			}
		} else {
			if err := s.client.Albums.
				Update().
				Where(
					albums.IDEQ(
						params.ID,
					),
				).
				RemoveUserIDs(
					s.GetUserID(),
				).
				Exec(s.ctx); err != nil {
				return fmt.Errorf("edit album remove user error: %w", err)
			}
			if err := s.client.Albums.Create().SetTitle(query.Title).
				SetDescription(query.Description).
				SetDate(query.Date).
				AddArtists(query.Edges.Artists...).
				AddMusics(final...).
				SetImageID(query.Edges.Image.ID).
				Exec(s.ctx); err != nil {
				return fmt.Errorf("add musics to album create error: %w", err)
			}

		}

	}

	Success(s.ctx, gin.H{"code": http.StatusOK})
	return nil
}

type RemoveMusicsFromAlbumParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

type RemoveMusicsFromAlbumForm struct {
	MusicIDs []string `form:"music_ids" json:"music_ids" binding:"required,gt=0,dive,min=1,numeric,excludes= "`
}

func (s *Server) RemoveMusicsFromAlbum() error {
	defer s.client.Close()

	var params RemoveMusicsFromAlbumParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}
	var forms RemoveMusicsFromAlbumForm
	if err := s.ctx.ShouldBind(&forms); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	query, err := BuildAlbumQuery(s.client).
		Where(
			albums.IDEQ(
				params.ID,
			),
		).
		WithUsers().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNotFound(s.ctx, ErrorAlbumNotFound)
			return nil
		}
		return fmt.Errorf("remove musics from album query error: %w", err)
	}

	c, err := helper.StrSliceToUint64Slice(forms.MusicIDs)
	if err != nil {
		return fmt.Errorf("remove musics from album convert error: %w", err)
	}

	all, err := s.client.Musics.
		Query().
		Where(
			musics.IDIn(c...),
		).
		All(s.ctx)
	if err != nil {
		return fmt.Errorf("remove musics from album query error: %w", err)
	}

	edit := s.client.Albums.Update().Where(albums.IDEQ(query.ID))
	if query.Edges.Users != nil {
		if len(query.Edges.Users) == 1 && query.Edges.Users[0].ID == s.GetUserID() {
			if err := edit.RemoveMusics(all...).Exec(s.ctx); err != nil {
				return fmt.Errorf("add musics to album error: %w", err)
			}
		} else if len(query.Edges.Users) < 1 {
			if err := edit.AddUserIDs(s.GetUserID()).RemoveMusics(all...).Exec(s.ctx); err != nil {
				return fmt.Errorf("add musics to album create error: %w", err)
			}
		} else {
			if err := s.client.Albums.
				Update().
				Where(
					albums.IDEQ(
						params.ID,
					),
				).
				RemoveUserIDs(
					s.GetUserID(),
				).
				Exec(s.ctx); err != nil {
				return fmt.Errorf("edit album remove user error: %w", err)
			}
			removes := make(map[uint64]struct{})
			for _, music := range all {
				removes[music.ID] = struct{}{}
			}

			final := make([]*ent.Musics, 0)
			for _, music := range query.Edges.Musics {
				if _, found := removes[music.ID]; !found {
					final = append(final, music)
				}
			}

			if err := s.client.Albums.Create().SetTitle(query.Title).
				SetDescription(query.Description).
				SetDate(query.Date).
				AddArtists(query.Edges.Artists...).
				AddMusics(final...).
				SetImageID(query.Edges.Image.ID).
				Exec(s.ctx); err != nil {
				return fmt.Errorf("add musics to album create error: %w", err)
			}

		}
	}

	Success(s.ctx, gin.H{"code": http.StatusOK})
	return nil
}

type EditAlbumParams struct {
	ID uint64 `uri:"id" binding:"required"`
}

type EditAlbumForm struct {
	Title           string   `form:"title" json:"title" binding:"omitempty,min=1,max=255,startsnotwith= ,endsnotwith= "`
	Date            string   `form:"date" json:"date" binding:"omitempty,min=4,numeric"`
	CoverID         string   `form:"cover_id" json:"cover_id" binding:"omitempty,numeric"`
	AddArtistIDs    []string `form:"add_artist_ids" json:"add_artist_ids" binding:"omitempty,gt=0,dive,min=1,numeric,excludes= "`
	RemoveArtistIDs []string `form:"remove_artist_ids" json:"remove_artist_ids" binding:"omitempty,gt=0,dive,min=1,numeric,excludes= "`
	Description     string   `form:"description" json:"description" binding:"omitempty"`
}

func (s *Server) EditAlbum() error {
	defer s.client.Close()

	var params EditAlbumParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	var forms EditAlbumForm
	if err := s.ctx.ShouldBind(&forms); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	query, err := BuildAlbumQuery(s.client).
		Where(
			albums.IDEQ(
				params.ID,
			),
		).
		WithUsers().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNotFound(s.ctx, ErrorAlbumNotFound)
			return nil
		}
		return fmt.Errorf("edit album query error: %w", err)
	}

	edit := s.client.Albums.Update().Where(albums.IDEQ(query.ID))
	create := s.client.Albums.Create().AddUserIDs(s.GetUserID())
	if forms.Title != "" {
		edit.SetTitle(forms.Title)
		query.Title = forms.Title
	}
	if forms.Description != "" {
		edit.SetDescription(forms.Description)
		query.Description = forms.Description
	}
	if forms.Date != "" {
		edit.SetDate(forms.Date)
		query.Date = forms.Date
	}
	if forms.CoverID != "" {
		cid, _ := strconv.ParseUint(forms.CoverID, 10, 64)
		exist, err := s.client.Images.Query().Where(images.IDEQ(cid)).Exist(s.ctx)
		if err != nil {
			return err
		}
		if !exist {
			ErrorNotFound(s.ctx, ImageNotFound)
			return nil
		}
		edit.SetImageID(cid)
		query.Edges.Image.ID = cid
	}
	if len(forms.AddArtistIDs) != 0 || len(forms.RemoveArtistIDs) != 0 {
		adds, err := helper.StrSliceToUint64SliceMap(forms.AddArtistIDs)
		if err != nil {
			return fmt.Errorf("edit album artists parse ids error: %w", err)
		}
		removes, err := helper.StrSliceToUint64SliceMap(forms.RemoveArtistIDs)
		if err != nil {
			return fmt.Errorf("edit album artists parse ids error: %w", err)
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
				return fmt.Errorf("edit album update error: %w", err)
			}
			Success(s.ctx, gin.H{
				"code": http.StatusOK,
				"id":   strconv.FormatUint(query.ID, 10),
			})
		} else if len(query.Edges.Users) < 1 {
			if err := edit.AddUserIDs(s.GetUserID()).Exec(s.ctx); err != nil {
				return fmt.Errorf("edit album update error: %w", err)
			}
			Success(s.ctx, gin.H{
				"code": http.StatusOK,
				"id":   strconv.FormatUint(query.ID, 10),
			})
		} else {
			if err := s.client.Albums.Update().Where(albums.IDEQ(params.ID)).RemoveUserIDs(s.GetUserID()).Exec(s.ctx); err != nil {
				return fmt.Errorf("edit album remove user error: %w", err)
			}
			c, err := create.
				SetTitle(query.Title).
				SetDescription(query.Description).
				SetDate(query.Date).
				AddArtists(query.Edges.Artists...).
				AddMusics(query.Edges.Musics...).
				SetImageID(query.Edges.Image.ID).
				Save(s.ctx)
			if err != nil {
				return fmt.Errorf("edit album create error: %w", err)
			}
			Success(s.ctx, gin.H{
				"code": http.StatusOK,
				"id":   strconv.FormatUint(c.ID, 10),
			})
		}
	}

	return nil
}

type DelAlbumParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

func (s *Server) DelAlbum() error {
	defer s.client.Close()

	var params DelAlbumParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}
	query, err := s.client.Albums.
		Query().
		Where(
			albums.IDEQ(
				params.ID,
			),
			albums.HasUsersWith(
				users.IDEQ(
					s.GetUserID(),
				),
			),
		).
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("delete album query error: %w", err)
	}
	if err := s.client.Albums.Update().Where(albums.IDEQ(query.ID)).RemoveUserIDs(s.GetUserID()).Exec(s.ctx); err != nil {
		return fmt.Errorf("delete album error: %w", err)
	}

	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}
