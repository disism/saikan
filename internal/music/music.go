package music

import (
	"encoding/json"
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/file"
	"github.com/disism/saikan/ent/music"
	"github.com/disism/saikan/ent/user"
	saved "github.com/disism/saikan/internal/saved"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	ctx    *gin.Context
	client *ent.Client
}

func NewServer(ctx *gin.Context, client *ent.Client) *Server {
	return &Server{ctx: ctx, client: client}
}

type AddMusicFileData struct {
	Hash string `json:"hash"`
	Name string `json:"name"`
	Size string `json:"size"`
}

type ExistsMusic struct {
	Name  string  `json:"name"`
	Music []Music `json:"music"`
}

type Music struct {
	ID          string      `json:"id,omitempty"`
	CreateTime  time.Time   `json:"create_time"`
	UpdateTime  time.Time   `json:"update_time"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	File        *saved.File `json:"file"`
	Artists     []*Artist   `json:"artists"`
	Albums      []*Album    `json:"albums"`
}

// AddMusic adds music to the collection.
//
// It loops over music files, checks if they are already in use by the user via the CID.
// If the file does not exist, it creates it.
// For each music file, it creates a new music entry with the extracted filename and duration.
// It also queries the existence of the file and retrieves all related information, such as artists and albums.
// Finally, it commits the transaction and returns the created music entries and existing musis.
//
// Returns an error if any operation fails.
func (s *Server) AddMusic() error {
	defer s.client.Close()

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		slog.Error("authx new database tx error: ", err.Error())
		return nil
	}
	defer tx.Rollback()

	var creates []*Music
	var exists []*ExistsMusic
	var ef []*ent.File

	// Loop over music files, check if they are already in use by the user via the CID.
	// Check if the file exists, if not, create it.
	for _, m := range s.ctx.PostFormArray("music") {
		var d AddMusicFileData
		if err := json.Unmarshal([]byte(m), &d); err != nil {
			return err
		}
		if err := func() error {
			qf, err := tx.
				File.
				Query().
				Where(
					file.HashEQ(
						d.Hash,
					),
				).
				Only(s.ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					size, err := strconv.ParseUint(d.Size, 10, 64)
					if err != nil {
						slog.Error("add music parse size error: ", err.Error())
						return err
					}
					cf, err := tx.File.
						Create().
						SetHash(d.Hash).
						SetName(d.Name).
						SetSize(size).
						Save(s.ctx)
					if err != nil {
						slog.Error("add music create music error: ", err.Error())
						return err
					}

					// Extract filenames (remove file extensions)
					// By finding the last dot (.) index position of the last dot (.).
					// Intercept a portion of the string create.Name from the beginning to the position before the last dot to get the filename.
					fn := cf.Name[:strings.LastIndex(cf.Name, ".")]
					create, err := tx.Music.
						Create().
						AddUserIDs(
							server.NewContext(s.ctx).GetUserID(),
						).
						SetName(fn).
						SetFileID(cf.ID).
						Save(s.ctx)
					if err != nil {
						return err
					}
					creates = append(creates, &Music{
						ID:          strconv.FormatUint(create.ID, 10),
						CreateTime:  create.CreateTime,
						UpdateTime:  create.UpdateTime,
						Name:        create.Name,
						Description: create.Description,
						File: &saved.File{
							ID:   strconv.FormatUint(cf.ID, 10),
							Hash: cf.Hash,
							Name: cf.Name,
							Size: strconv.FormatUint(cf.Size, 10),
						},
					})
					return nil
				}
				slog.Error("add music create music error: ", err.Error())
				return err
			}

			ef = append(ef, qf)
			return nil
		}(); err != nil {
			return err
		}
	}
	// already queried the existence of the file, it is queried through the file.
	for _, exist := range ef {
		all, err := tx.Music.
			Query().
			Where(
				music.HasFileWith(
					file.IDEQ(exist.ID),
				),
			).
			WithFile().
			WithArtists().
			WithAlbums(
				func(query *ent.AlbumQuery) {
					query.WithImage(
						func(query *ent.ImageQuery) {
							query.WithFile()
						},
					)
					query.WithArtists()
				},
			).
			All(s.ctx)
		if err != nil {
			return err
		}
		var x []Music
		for _, m2 := range all {
			x = append(x, *FMTMusicObject(m2))
		}
		exists = append(exists, &ExistsMusic{
			Name:  exist.Name,
			Music: x,
		})
	}

	if err := tx.Commit(); err != nil {
		slog.Error("add music commit error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"creates": creates,
		"exists":  exists,
	})
	return nil
}

func FMTMusicObject(m *ent.Music) *Music {
	r := &Music{
		ID:          strconv.FormatUint(m.ID, 10),
		CreateTime:  m.CreateTime,
		UpdateTime:  m.UpdateTime,
		Name:        m.Name,
		Description: m.Description,
		File: &saved.File{
			ID:   strconv.FormatUint(m.Edges.File.ID, 10),
			Hash: m.Edges.File.Hash,
			Name: m.Edges.File.Name,
			Size: strconv.FormatUint(m.Edges.File.Size, 10),
		},
	}
	if m.Edges.Artists != nil {
		for _, artist := range m.Edges.Artists {
			r.Artists = append(r.Artists, &Artist{
				ID:   strconv.FormatUint(artist.ID, 10),
				Name: artist.Name,
			})
		}
	}
	if m.Edges.Albums != nil {

	}
	return r
}

// GetMusics retrieves all the musics associated with the user in the database.
//
// It returns an error if there was a problem retrieving the musics.
func (s *Server) GetMusics() error {
	defer s.client.Close()
	query := s.client.
		Music.
		Query().
		Where(
			music.HasUserWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		WithFile().
		WithArtists().
		WithAlbums(
			func(query *ent.AlbumQuery) {
				query.WithImage(func(query *ent.ImageQuery) {
					query.WithFile()
				})
				query.WithArtists()
			},
		)
	name := s.ctx.Query("name")
	if strings.TrimSpace(name) != "" {
		query.Where(music.NameContainsFold(name))
	}

	all, err := query.All(s.ctx)
	if err != nil {
		return err
	}

	r := make([]*Music, len(all))
	for i, a := range all {
		r[i] = FMTMusicObject(a)
	}

	s.ctx.JSON(http.StatusOK, r)
	return nil
}

func (s *Server) QueryUserMusic(musicID uint64) (*ent.Music, error) {
	return s.client.Music.
		Query().
		Where(
			music.IDEQ(musicID),
			music.HasUserWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).
		WithFile().
		WithArtists().
		WithAlbums(
			func(query *ent.AlbumQuery) {
				query.WithMusics(func(query *ent.MusicQuery) {
					query.WithFile()
				})
				query.WithArtists()
			},
		).
		Only(s.ctx)
}

func (s *Server) GetMusic() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}

	query, err := s.QueryUserMusic(id)
	if err != nil {
		return err
	}
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("get music query error: ", err.Error())
		return err
	}
	s.ctx.JSON(http.StatusOK, FMTMusicObject(query))
	return nil
}

func (s *Server) EditMusic() error {
	defer s.client.Close()

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		slog.Error("edit music new database tx error: ", err.Error())
		return err
	}
	defer tx.Rollback()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}

	query, err := s.QueryUserMusic(id)
	if err != nil {
		slog.Error("edit music query error: ", err.Error())
		return err
	}

	update := tx.Music.
		Update().
		Where(
			music.IDEQ(query.ID),
		)

	name := s.ctx.PostForm("name")
	desc := s.ctx.PostForm("description")

	switch {
	case strings.TrimSpace(name) != "":
		if name == query.Name {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("no change in name"))
		}
		update.SetName(name)
	case strings.TrimSpace(desc) != "":
		if desc == query.Description {
			return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("no change in description"))
		}
		update.SetDescription(desc)
	default:
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("no changes"))
	}
	fmt.Println(name, desc)

	if err := update.Exec(s.ctx); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		slog.Error("edit music commit error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
	return nil
}

func (s *Server) AddArtistsToMusic() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}

	query, err := s.QueryUserMusic(id)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("add music artist query error: ", err.Error())
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
		if query.Edges.Artists != nil {
			if err := func() error {
				for _, a := range query.Edges.Artists {
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

	if err := s.client.Music.Update().Where(music.IDEQ(query.ID)).AddArtists(artists...).Exec(s.ctx); err != nil {
		slog.Error("add artists to music error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})

	return nil
}

func (s *Server) RemoveArtistsFromMusic() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}

	query, err := s.QueryUserMusic(id)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("remove artists music query error: ", err.Error())
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
		if query.Edges.Artists != nil {
			for _, r := range query.Edges.Artists {
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

	if err := s.client.Music.Update().Where(music.IDEQ(query.ID)).RemoveArtists(artists...).Exec(s.ctx); err != nil {
		slog.Error("remove artists from music error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}

func (s *Server) AddMusicToLibrary() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}

	query, err := s.client.Music.
		Query().
		Where(
			music.IDEQ(id),
		).
		WithUser().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("add music to lib query error: ", err.Error())
		return err
	}
	if query.Edges.User != nil {
		for _, u := range query.Edges.User {
			if u.ID == server.NewContext(s.ctx).GetUserID() {
				return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("music already in library"))
			}
		}
	}

	if err := s.client.User.Update().
		Where(
			user.IDEQ(
				server.NewContext(s.ctx).GetUserID(),
			),
		).
		AddMusics(query).
		Exec(s.ctx); err != nil {
		slog.Error("add music to lib error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
	return nil
}

func (s *Server) DeleteMusic() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("InvalidID"))
	}
	query, err := s.QueryUserMusic(id)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("delete music query error: ", err.Error())
		return err
	}

	if err := s.client.
		User.
		Update().
		Where(
			user.IDEQ(
				server.NewContext(s.ctx).GetUserID(),
			),
		).
		RemoveMusics(query).
		Exec(s.ctx); err != nil {
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
	return nil
}
