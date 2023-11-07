package music

import (
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/image"
	"github.com/disism/saikan/ent/playlist"
	"github.com/disism/saikan/ent/user"
	"github.com/disism/saikan/internal/errors"
	"github.com/disism/saikan/internal/saved"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
	"strings"

	img "github.com/disism/saikan/internal/image"
)

type Playlist struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Private     bool       `json:"private"`
	Image       *img.Image `json:"image"`
	Musics      []*Music   `json:"musics"`
}

func (s *Server) CreatePlaylist() error {
	defer s.client.Close()

	desc := s.ctx.PostForm("description")
	name := s.ctx.PostForm("name")
	cover := s.ctx.PostForm("image_id")
	private := s.ctx.PostForm("private")

	// Start a new transaction
	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		slog.Error("create playlist tx error: ", err.Error())
		return err
	}
	defer tx.Rollback()

	// Check if the cover image exists
	if strings.TrimSpace(cover) != "" {
		id, err := strconv.ParseUint(cover, 10, 64)
		if err != nil {
			slog.Error("create playlist parse cover id error: ", err.Error())
			return err
		}
		exists, err := tx.Image.Query().
			Where(image.IDEQ(id)).
			Exist(s.ctx)
		if err != nil {
			slog.Error("create playlist query cover error: ", err.Error())
			return err
		}
		if !exists {
			if ent.IsNotFound(err) {
				return server.NewContext(s.ctx).ErrorNotFound(err)
			}
		}
	}

	// Check if a playlist with the same name already exists for the current user
	create, err := func() (*ent.Playlist, error) {
		query, err := tx.Playlist.Query().Where(
			playlist.NameEQ(name),
			playlist.HasOwnerWith(
				user.IDEQ(server.NewContext(s.ctx).GetUserID()),
			),
		).
			WithImage(func(query *ent.ImageQuery) {
				query.WithFile()
			}).
			Only(s.ctx)
		if err != nil {
			if ent.IsValidationError(err) {
				return nil, server.NewContext(s.ctx).ErrorUnprocessableEntity(err)
			}
			if ent.IsNotFound(err) {
				// Create a new playlist
				p := tx.Playlist.Create()
				p.SetName(name).
					SetDescription(desc).
					SetOwnerID(server.NewContext(s.ctx).GetUserID())
				if strings.TrimSpace(private) != "" {
					b, err := strconv.ParseBool(private)
					if err != nil {
						return nil, server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidPrivateValue: ExpectationBoolean"))
					}
					p.SetPrivate(b)
				}
				if strings.TrimSpace(cover) != "" {
					id, err := strconv.ParseUint(cover, 10, 64)
					if err != nil {
						slog.Error("create playlist parse cover id error: ", err.Error())
						return nil, err
					}
					p.SetImageID(id)
				}
				// Save the new playlist
				create, err := p.Save(s.ctx)
				if err != nil {
					if ent.IsValidationError(err) {
						return nil, server.NewContext(s.ctx).ErrorUnprocessableEntity(err)
					}
					slog.Error("create playlist create error: ", err.Error())
					return nil, err
				}
				return create, nil
			}
			slog.Error("create playlist query error: ", err.Error())
			return nil, err
		}
		return query, nil
	}()
	if err != nil {
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		slog.Error("create playlist commit error: ", err.Error())
		return err
	}

	// Return the created playlist
	if create != nil {
		s.ctx.JSON(http.StatusCreated, create)
	}
	return nil
}

// GetPlaylists retrieves all playlists owned by the user.
//
// It returns an error if there is an issue creating a new database or querying the database.
// The function also returns an error if there is an issue retrieving the playlists.
// If successful, it returns a JSON response with all the playlists.
func (s *Server) GetPlaylists() error {
	defer s.client.Close()

	all, err := s.client.Playlist.
		Query().
		Where(
			playlist.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		WithImage(
			func(query *ent.ImageQuery) {
				query.WithFile()
			},
		).
		All(s.ctx)
	if err != nil {
		return err
	}

	r := make([]*Playlist, len(all))

	for i, p := range all {
		r[i] = FMTPlaylistObject(p)
	}
	s.ctx.JSON(http.StatusOK, r)
	return nil
}

func FMTPlaylistObject(p *ent.Playlist) *Playlist {
	r := &Playlist{
		ID:          strconv.FormatUint(p.ID, 10),
		Name:        p.Name,
		Description: p.Description,
		Private:     p.Private,
	}
	if p.Edges.Image != nil {
		r.Image = &img.Image{
			ID:     strconv.FormatUint(p.Edges.Image.ID, 10),
			Width:  p.Edges.Image.Width,
			Height: p.Edges.Image.Height,
			File: &saved.File{
				ID:   strconv.FormatUint(p.Edges.Image.Edges.File.ID, 10),
				Hash: p.Edges.Image.Edges.File.Hash,
				Name: p.Edges.Image.Edges.File.Name,
				Size: strconv.FormatUint(p.Edges.Image.Edges.File.Size, 10),
			},
		}
	}
	if p.Edges.Musics != nil {
		r.Musics = make([]*Music, len(p.Edges.Musics))
		for i, m := range p.Edges.Musics {
			r.Musics[i] = FMTMusicObject(m)
		}
	}
	return r
}

// GetPlaylist retrieves a playlist by its ID from the Music API.
//
// It takes no parameters.
// It returns an error if the playlist ID is invalid or if there is an error querying the database.
func (s *Server) GetPlaylist() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}
	query, err := s.client.Playlist.
		Query().
		Where(
			playlist.IDEQ(
				id,
			),
		).
		WithImage(
			func(query *ent.ImageQuery) {
				query.WithFile()
			},
		).
		WithOwner().
		WithMusics(
			func(query *ent.MusicQuery) {
				query.WithFile()
				query.WithArtists()
			},
		).
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf(errors.NotFound))
		}
		return err
	}
	if query.Private {
		if query.Edges.Owner != nil {
			if query.Edges.Owner.ID != server.NewContext(s.ctx).GetUserID() {
				return server.NewContext(s.ctx).ErrorNoPermission()
			}
		}
		slog.Error("playlist exists with no owner:. ", query)
	}

	s.ctx.JSON(http.StatusOK, FMTPlaylistObject(query))
	return nil
}

func (s *Server) EditPlaylist() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}
	query, err := s.client.Playlist.Query().
		Where(
			playlist.IDEQ(id),
			playlist.HasOwnerWith(
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
		return err
	}
	name := s.ctx.PostForm("name")
	desc := s.ctx.PostForm("description")

	edit := s.client.Playlist.Update().Where(playlist.IDEQ(id))
	switch {
	case strings.TrimSpace(name) != "":
		if name == query.Name {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("NameUnchanged"))
		}
		edit.SetName(name)
	case strings.TrimSpace(desc) != "":
		if desc == query.Description {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("DescriptionUnchanged"))
		}
		edit.SetDescription(desc)
	default:
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("NoChanges"))
	}

	if err := edit.Exec(s.ctx); err != nil {
		slog.Error("edit playlist error: ", err.Error())
		return err
	}
	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) EditPlaylistImage() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}
	query, err := s.client.Playlist.
		Query().
		Where(
			playlist.IDEQ(
				id,
			),
			playlist.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		WithImage().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		return err
	}
	imgID, err := strconv.ParseUint(s.ctx.PostForm("image_id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}
	if query.Edges.Image != nil {
		if query.Edges.Image.ID == imgID {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("CoverUnchanged"))
		}
	}

	qc, err := s.client.Image.Query().
		Where(image.IDEQ(imgID)).
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("NewCoverNotFound"))
		}
		return err
	}

	if err := s.client.Playlist.
		Update().
		Where(
			playlist.IDEQ(id),
		).
		SetImage(qc).
		Exec(s.ctx); err != nil {
		slog.Error("edit playlist error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) RemovePlaylistImage() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}

	query, err := s.client.Playlist.
		Query().
		Where(
			playlist.IDEQ(
				id,
			),
			playlist.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
	}

	if err := s.client.Playlist.Update().Where(playlist.IDEQ(query.ID)).ClearImage().Exec(s.ctx); err != nil {
		slog.Error("clear playlist cover error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) AddMusicsToPlaylist() error {
	defer s.client.Close()
	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidAlbumID"))
	}
	query, err := s.client.Playlist.
		Query().
		Where(
			playlist.IDEQ(id),
			playlist.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		WithMusics().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf("playlist not found"))
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
		if query.Edges.Musics != nil {
			if err := func() error {
				for _, a := range query.Edges.Musics {
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
	if err := s.client.Playlist.Update().Where(playlist.IDEQ(query.ID)).AddMusics(musics...).Exec(s.ctx); err != nil {
		slog.Error("add musics to playlist error: ", err.Error())
		return err
	}
	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) RemoveMusicsFromPlaylist() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidAlbumID"))
	}
	query, err := s.client.Playlist.
		Query().
		Where(
			playlist.IDEQ(id),
			playlist.HasOwnerWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		WithMusics().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNotFound(fmt.Errorf("playlist not found"))
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
		if query.Edges.Musics != nil {
			for _, r := range query.Edges.Musics {
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
	if err := s.client.Playlist.Update().Where(playlist.IDEQ(query.ID)).RemoveMusics(musics...).Exec(s.ctx); err != nil {
		slog.Error("add musics to playlist error: ", err.Error())
		return err
	}
	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})

	return nil
}
func (s *Server) DeletePlaylist() error {

	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}
	query, err := s.client.Playlist.
		Query().
		Where(
			playlist.IDEQ(
				id,
			),
			playlist.HasOwnerWith(
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
		return err
	}
	if err := s.client.Playlist.DeleteOne(query).Exec(s.ctx); err != nil {
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}
