package music

import (
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/album"
	"github.com/disism/saikan/ent/image"
	"github.com/disism/saikan/ent/user"
	img "github.com/disism/saikan/internal/image"
	"github.com/disism/saikan/internal/saved"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) QueryUserAlbum(albumID uint64) (*ent.Album, error) {
	return s.client.
		Album.
		Query().
		Where(
			album.IDEQ(albumID),
			album.HasUsersWith(
				user.IDEQ(server.NewContext(s.ctx).GetUserID()),
			),
		).
		WithArtists().
		WithMusics(
			func(query *ent.MusicQuery) {
				query.WithFile()
			},
		).
		WithImage(
			func(query *ent.ImageQuery) {
				query.WithFile()
			},
		).
		Only(s.ctx)
}

type Album struct {
	ID          string     `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description"`
	Year        string     `json:"year"`
	Image       *img.Image `json:"image"`
	Musics      []*Music   `json:"musics"`
	Artists     []*Artist  `json:"artists"`
}

func FMTAlbumObject(a *ent.Album) *Album {
	r := &Album{
		ID:          strconv.FormatUint(a.ID, 10),
		Title:       a.Title,
		Description: a.Description,
		Year:        strconv.FormatInt(int64(a.Year), 10),
	}
	if a.Edges.Image != nil {
		r.Image = &img.Image{
			ID:     strconv.FormatUint(a.Edges.Image.ID, 10),
			Width:  a.Edges.Image.Width,
			Height: a.Edges.Image.Height,
			File: &saved.File{
				ID:   strconv.FormatUint(a.Edges.Image.Edges.File.ID, 10),
				Hash: a.Edges.Image.Edges.File.Hash,
				Name: a.Edges.Image.Edges.File.Name,
				Size: strconv.FormatUint(a.Edges.Image.Edges.File.Size, 10),
			},
		}
	}
	if a.Edges.Musics != nil {
		r.Musics = make([]*Music, len(a.Edges.Musics))
		for i, m := range a.Edges.Musics {
			r.Musics[i] = FMTMusicObject(m)
		}
	}
	if a.Edges.Artists != nil {
		r.Artists = make([]*Artist, len(a.Edges.Artists))
		for i, a := range a.Edges.Artists {
			r.Artists[i] = &Artist{
				ID:   strconv.FormatUint(a.ID, 10),
				Name: a.Name,
			}
		}
	}
	return r
}

func (s *Server) CreateAlbum() error {
	title := s.ctx.PostForm("title")
	aids := s.ctx.PostFormArray("artist_id")
	description := s.ctx.PostForm("description")
	year := s.ctx.PostForm("year")
	cover := s.ctx.PostForm("image_id")

	if strings.TrimSpace(cover) == "" {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("image is required"))
	}

	if strings.TrimSpace(year) == "" {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("year is required"))
	}

	defer s.client.Close()

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		slog.Error("create album tx error: ", err.Error())
		return err
	}
	defer tx.Rollback()

	query, err := func() (*ent.Album, error) {
		only, err := tx.Album.Query().Where(
			album.TitleEQ(title),
			album.HasUsersWith(
				user.IDEQ(server.NewContext(s.ctx).GetUserID()),
			),
		).WithArtists().
			WithImage(
				func(query *ent.ImageQuery) {
					query.WithFile()
				},
			).
			Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				var artists []*ent.Artist
				if err := func() error {
					for _, artist := range aids {
						id, err := strconv.ParseUint(artist, 10, 64)
						if err != nil {
							slog.Error("create album parse artist error: ", err.Error())
							return err
						}
						a, err := tx.Artist.Get(s.ctx, id)
						if err != nil {
							if ent.IsNotFound(err) {
								return fmt.Errorf("artist not found")
							}
							slog.Error("create album get artist error: ", err.Error())
							return err
						}
						artists = append(artists, a)
					}
					return nil
				}(); err != nil {
					return nil, err
				}
				cid, err := strconv.ParseUint(cover, 10, 64)
				if err != nil {
					slog.Error("create album parse cover error: ", err.Error())
					return nil, err
				}
				img, err := tx.Image.Get(s.ctx, cid)
				if err != nil {
					if ent.IsNotFound(err) {
						return nil, fmt.Errorf("cover not found")
					}
					slog.Error("create album get cover error: ", err.Error())
					return nil, err
				}

				y, err := strconv.ParseUint(year, 10, 32)
				if err != nil {
					slog.Error("create album parse year error: ", err.Error())
					return nil, err
				}

				create, err := tx.Album.Create().
					SetTitle(title).
					AddArtists(artists...).
					SetImage(img).
					SetDescription(description).
					SetYear(uint32(y)).
					AddUserIDs(
						server.NewContext(s.ctx).GetUserID(),
					).
					Save(s.ctx)
				if err != nil {
					return nil, err
				}
				if err != nil {
					slog.Error("create album create error: ", err.Error())
					return nil, err
				}
				return create, nil
			}
			slog.Error("create album query error: ", err.Error())
			return nil, err
		}
		return only, nil
	}()
	if err != nil {
		return server.NewContext(s.ctx).ErrorBadRequest(err)
	}

	if err := tx.Commit(); err != nil {
		slog.Error("create album commit error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, FMTAlbumObject(query))
	return nil
}

func (s *Server) GetAlbums() error {
	defer s.client.Close()

	all, err := s.client.
		Album.
		Query().
		Where(
			album.HasUsersWith(
				user.IDEQ(server.NewContext(s.ctx).GetUserID()),
			),
		).
		WithArtists().
		WithImage(
			func(query *ent.ImageQuery) {
				query.WithFile()
			},
		).
		All(s.ctx)
	if err != nil {
		return err
	}

	r := make([]*Album, len(all))
	for i, a := range all {
		r[i] = FMTAlbumObject(a)
	}

	s.ctx.JSON(http.StatusOK, r)
	return nil
}

// GetAlbum retrieves an album from the Music API.
//
// It takes no parameters and returns an error.
func (s *Server) GetAlbum() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}

	query, err := s.QueryUserAlbum(id)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf("album not found"))
		}
		return err
	}

	s.ctx.JSON(http.StatusOK, FMTAlbumObject(query))
	return nil
}

func (s *Server) AddMusicsToAlbum() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidAlbumID"))
	}
	al, err := s.QueryUserAlbum(id)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf("album not found"))
		}
		return err
	}
	var musics []*ent.Music
	for _, s2 := range s.ctx.PostFormArray("music_id") {
		mi, err := strconv.ParseUint(s2, 10, 64)
		if err != nil {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidMusicID"))
		}
		qm, err := s.QueryUserMusic(mi)
		if err != nil {
			if ent.IsNotFound(err) {
				return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf("music not found"))
			}
			return err
		}
		if al.Edges.Musics != nil {
			if err := func() error {
				for _, a := range al.Edges.Musics {
					if a.ID == qm.ID {
						return fmt.Errorf("music [%s]: already exists", qm.Name)
					}
				}
				return nil
			}(); err != nil {
				return server.NewContext(s.ctx).ErrorUnprocessableEntity(err)
			}
		}
		musics = append(musics, qm)
	}
	if err := s.client.Album.Update().Where(album.IDEQ(al.ID)).AddMusics(musics...).Exec(s.ctx); err != nil {
		slog.Error("add musics to album error: ", err.Error())
		return err
	}
	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) RemoveMusicsFromAlbum() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidAlbumID"))
	}
	al, err := s.QueryUserAlbum(id)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf("album not found"))
		}
		return err
	}
	var musics []*ent.Music
	for _, s2 := range s.ctx.PostFormArray("music_id") {
		mi, err := strconv.ParseUint(s2, 10, 64)
		if err != nil {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidMusicID"))
		}
		qm, err := s.QueryUserMusic(mi)
		if err != nil {
			if ent.IsNotFound(err) {
				return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf("music not found"))
			}
			return err
		}

		found := false
		if al.Edges.Musics != nil {
			for _, r := range al.Edges.Musics {
				if r.ID == qm.ID {
					found = true
					break
				}
			}
		}

		if !found {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("music [%s]: no exists", qm.Name))
		}

		musics = append(musics, qm)
	}
	if err := s.client.Album.Update().Where(album.IDEQ(al.ID)).RemoveMusics(musics...).Exec(s.ctx); err != nil {
		slog.Error("add musics to album error: ", err.Error())
		return err
	}
	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) EditAlbum() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidAlbumID"))
	}
	al, err := s.QueryUserAlbum(id)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		return err
	}

	title := s.ctx.PostForm("title")
	desc := s.ctx.PostForm("description")
	year := s.ctx.PostForm("year")
	cover := s.ctx.PostForm("image_id")

	edit := s.client.Album.Update().Where(album.IDEQ(al.ID))
	switch {
	case strings.TrimSpace(title) != "":
		if al.Title == title {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("album title no changes"))
		}
		edit.SetTitle(title)
	case strings.TrimSpace(desc) != "":
		if al.Description == desc {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("album description no changes"))
		}
		edit.SetDescription(desc)
	case strings.TrimSpace(year) != "":
		y, err := strconv.ParseUint(year, 10, 32)
		if err != nil {
			return err
		}
		if al.Year == uint32(y) {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("album year no changes"))
		}
		edit.SetYear(uint32(y))
	case strings.TrimSpace(cover) != "":
		c, err := strconv.ParseUint(s.ctx.PostForm("image_id"), 10, 64)
		if err != nil {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidCoverID"))
		}
		if al.Edges.Image != nil {
			if al.Edges.Image.ID == c {
				return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("album cover no changes"))
			}
		}
		qc, err := s.client.Image.Query().Where(image.IDEQ(c)).Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf("CoverNotFound"))
			}
			return err
		}
		edit.SetImage(qc)

	default:
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("no changes"))
	}

	if err := edit.Exec(s.ctx); err != nil {
		return err
	}
	server.NewContext(s.ctx).JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) AddArtistsToAlbum() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidAlbumID"))
	}
	al, err := s.QueryUserAlbum(id)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		return err
	}
	var artists []*ent.Artist

	for _, s2 := range s.ctx.PostFormArray("artist_id") {
		aid, err := strconv.ParseUint(s2, 10, 64)
		if err != nil {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidArtistID"))
		}
		artist, err := s.QueryArtist(aid)
		if err != nil {
			return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf("ArtistNotFound"))
		}
		if al.Edges.Artists != nil {
			if err := func() error {
				for _, a := range al.Edges.Artists {
					if a.ID == artist.ID {
						return fmt.Errorf("artist [%s]: already exists", artist.Name)
					}
				}
				return nil
			}(); err != nil {
				return server.NewContext(s.ctx).ErrorUnprocessableEntity(err)
			}
		}
		artists = append(artists, artist)
	}

	if err := s.client.Album.Update().Where(album.IDEQ(al.ID)).AddArtists(artists...).Exec(s.ctx); err != nil {
		slog.Error("add artists to album error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) RemoveArtistsFromAlbum() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidAlbumID"))
	}
	al, err := s.QueryUserAlbum(id)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		return err
	}
	var artists []*ent.Artist
	for _, s2 := range s.ctx.PostFormArray("artist_id") {
		aid, err := strconv.ParseUint(s2, 10, 64)
		if err != nil {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidArtistID"))
		}
		artist, err := s.QueryArtist(aid)
		if err != nil {
			return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf("ArtistNotFound"))
		}

		found := false
		if al.Edges.Artists != nil {
			for _, r := range al.Edges.Artists {
				if r.ID == artist.ID {
					found = true
					break
				}
			}
		}

		if !found {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("artist [%s]: not exists", artist.Name))
		}
		artists = append(artists, artist)
	}

	if err := s.client.Album.Update().Where(album.IDEQ(al.ID)).RemoveArtists(artists...).Exec(s.ctx); err != nil {
		slog.Error("delete artist form album error: ", err.Error())
		return err
	}
	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) DeleteAlbum() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}

	query, err := s.QueryUserAlbum(id)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("delete album query error: ", err.Error())
		return err
	}
	if err := s.client.
		Album.
		Update().
		Where(
			album.IDEQ(query.ID),
		).
		RemoveUserIDs(
			server.NewContext(s.ctx).GetUserID(),
		).
		Exec(s.ctx); err != nil {
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}
