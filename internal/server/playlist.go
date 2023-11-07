package server

import (
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/images"
	"github.com/disism/saikan/ent/musics"
	"github.com/disism/saikan/ent/playlists"
	"github.com/disism/saikan/ent/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Playlists interface {
	CreatePlaylist() error
	ListPlaylists() error
	EditPlaylist() error
	RemoveMusicsFromPlaylist() error
	AddMusicsToPlaylist() error
	DelPlaylist() error
}

const (
	PlaylistExist    = "playlist exist"
	PlaylistNotFound = "playlist not found"
)

type Playlist struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Private     bool     `json:"private"`
	Image       *Image   `json:"image"`
	Musics      []*Music `json:"musics"`
}

func FMTPlaylist(p *ent.Playlists) *Playlist {
	r := &Playlist{
		ID:          strconv.FormatUint(p.ID, 10),
		Name:        p.Name,
		Description: p.Description,
		Private:     p.Private,
	}
	if p.Edges.Image != nil {
		r.Image = FMTImage(p.Edges.Image)
	}
	if p.Edges.Musics != nil {
		r.Musics = make([]*Music, len(p.Edges.Musics))
		for i, music := range p.Edges.Musics {
			r.Musics[i] = FMTMusic(music)
		}
	}
	return r
}

type CreatePlaylistForm struct {
	Name    string `form:"name" json:"name" binding:"required,max=32"`
	Private bool   `form:"private" json:"private"`
}

func (s *Server) CreatePlaylist() error {
	defer s.client.Close()

	var form CreatePlaylistForm
	if err := s.ctx.ShouldBind(&form); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("create playlist new tx error: %w", err)
	}
	defer tx.Rollback()

	exist, err := tx.Playlists.Query().
		Where(
			playlists.NameEQ(form.Name),
			playlists.HasOwnerWith(
				users.IDEQ(s.GetUserID()),
			),
		).
		Exist(s.ctx)
	if err != nil {
		return fmt.Errorf("create playlist exist playlist error: %w", err)
	}
	if exist {
		ErrorConflict(s.ctx, PlaylistExist)
		return nil
	}

	create, err := tx.Playlists.
		Create().
		SetOwnerID(
			s.GetUserID(),
		).
		SetName(form.Name).
		SetPrivate(form.Private).
		Save(s.ctx)
	if err != nil {
		return fmt.Errorf("create playlist error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("create playlist commit error: %w", err)
	}

	s.ctx.JSON(http.StatusCreated, gin.H{
		"code": http.StatusCreated,
		"id":   strconv.FormatUint(create.ID, 10),
	})
	return nil
}

func (s *Server) ListPlaylists() error {
	defer s.client.Close()

	all, err := s.client.Playlists.
		Query().
		Where(
			playlists.HasOwnerWith(
				users.IDEQ(
					s.GetUserID(),
				),
			),
		).
		WithImage(
			func(query *ent.ImagesQuery) {
				query.WithFile()
			},
		).
		All(s.ctx)
	if err != nil {
		return err
	}

	r := make([]*Playlist, len(all))

	for i, p := range all {
		r[i] = FMTPlaylist(p)
	}
	s.ctx.JSON(http.StatusOK, r)
	return nil
}

type GetPlaylistParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

func (s *Server) GetPlaylist() error {
	defer s.client.Close()

	var params GetPlaylistParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	query, err := s.client.Playlists.Query().
		Where(
			playlists.IDEQ(
				params.ID,
			),
		).
		WithImage(
			func(query *ent.ImagesQuery) {
				query.WithFile()
			},
		).
		WithOwner().
		WithMusics(
			func(query *ent.MusicsQuery) {
				query.WithFile()
				query.WithArtists()
			},
		).
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNotFound(s.ctx, PlaylistNotFound)
			return nil
		}
		return fmt.Errorf("get playlist error: %w", err)
	}
	if query.Private {
		if query.Edges.Owner != nil {
			if query.Edges.Owner.ID != s.GetUserID() {
				ErrorNoPermission(s.ctx, NoPermission)
				return nil
			}
		}
	}

	Success(s.ctx, FMTPlaylist(query))
	return nil
}

type EditPlaylistParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}
type EditPlaylistForm struct {
	Name        string `form:"name" json:"name" binding:"max=32"`
	Description string `form:"description" json:"description" binding:"max=500"`
	CoverID     string `form:"cover_id" json:"cover_id"`
	Private     bool   `form:"private" json:"private"`
}

func (s *Server) EditPlaylist() error {
	defer s.client.Close()

	var params EditPlaylistParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}
	var form EditPlaylistForm
	if err := s.ctx.ShouldBind(&form); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	query, err := s.client.Playlists.
		Query().
		Where(
			playlists.IDEQ(params.ID),
			playlists.HasOwnerWith(
				users.IDEQ(
					s.GetUserID(),
				),
			),
		).
		WithImage().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("edit playlist query error: %w", err)
	}

	name := strings.TrimSpace(form.Name)
	desc := strings.TrimSpace(form.Description)
	cid := strings.TrimSpace(form.CoverID)

	edit := s.client.Playlists.Update().Where(playlists.IDEQ(query.ID))

	if name != "" && name != query.Name {
		edit.SetName(name)
	}
	if desc != query.Description {
		edit.SetDescription(desc)
	}

	if cid != "" {
		iid, err := strconv.ParseUint(cid, 10, 64)
		if err != nil {
			ErrorBadRequest(s.ctx, err.Error())
			return nil
		}

		image, err := s.client.Images.
			Query().
			Where(
				images.IDEQ(iid),
			).
			Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				ErrorBadRequest(s.ctx, ImageNotFound)
				return nil
			}
			return fmt.Errorf("edit playlist image query error: %w", err)
		}

		edit.SetImageID(image.ID)
	}
	if form.Private != query.Private {
		edit.SetPrivate(form.Private)
	}
	if err := edit.Exec(s.ctx); err != nil {
		return fmt.Errorf("edit playlist error: %w", err)
	}

	Success(s.ctx, gin.H{"code": http.StatusOK})
	return nil
}

type RemovePlaylistImageParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

func (s *Server) RemovePlaylistImage() error {
	defer s.client.Close()

	var params RemovePlaylistImageParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	query, err := s.client.Playlists.
		Query().
		Where(
			playlists.IDEQ(
				params.ID,
			),
			playlists.HasOwnerWith(
				users.IDEQ(
					s.GetUserID(),
				),
			),
		).
		WithImage().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return err
	}
	if query.Edges.Image == nil {
		ErrorBadRequest(s.ctx, ImageNotFound)
		return nil
	}
	if err := s.client.Playlists.Update().Where(playlists.IDEQ(query.ID)).ClearImage().Exec(s.ctx); err != nil {
		return fmt.Errorf("remove playlist image error: %w", err)
	}
	Success(s.ctx, gin.H{"code": http.StatusOK})
	return nil
}

type AddMusicsToPlaylistParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

type AddMusicsToPlaylistForm struct {
	MusicIDs []string `form:"music_ids" json:"music_ids" binding:"required,gt=0,dive,min=1"`
}

func (s *Server) AddMusicsToPlaylist() error {
	defer s.client.Close()

	var params AddMusicsToPlaylistParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	query, err := s.client.Playlists.
		Query().
		Where(
			playlists.IDEQ(
				params.ID,
			),
			playlists.HasOwnerWith(
				users.IDEQ(
					s.GetUserID(),
				),
			),
		).
		WithMusics().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("add musics to playlist query error: %w", err)
	}

	var forms AddMusicsToPlaylistForm
	if err := s.ctx.ShouldBind(&forms); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	maps := make(map[uint64]bool)
	for _, v := range query.Edges.Musics {
		maps[v.ID] = true
	}

	conv := make([]uint64, 0, len(forms.MusicIDs))
	for _, str := range forms.MusicIDs {
		parse, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return fmt.Errorf("add musics to playlist parse music id error: %w", err)
		}
		conv = append(conv, parse)
	}

	all, err := s.client.Musics.
		Query().
		Where(
			musics.IDIn(conv...),
			musics.HasUsersWith(
				users.IDEQ(
					s.GetUserID(),
				),
			),
		).
		All(s.ctx)
	if err != nil {
		return fmt.Errorf("add musics to playlist query error: %w", err)
	}

	var adds []*ent.Musics
	for _, music := range all {
		if _, ok := maps[music.ID]; !ok {
			adds = append(adds, music)
		}
	}
	if err := s.client.Playlists.Update().Where(playlists.IDEQ(query.ID)).AddMusics(adds...).Exec(s.ctx); err != nil {
		return fmt.Errorf("add musics to playlist error: %w", err)
	}

	Success(s.ctx, gin.H{"code": http.StatusOK})
	return nil
}

type RemoveMusicsFromPlaylistParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

type RemoveMusicsFromPlaylistForm struct {
	MusicIDs []string `form:"music_ids" json:"music_ids" binding:"required,gt=0,dive,min=1"`
}

func (s *Server) RemoveMusicsFromPlaylist() error {
	defer s.client.Close()

	var params RemoveMusicsFromPlaylistParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	query, err := s.client.Playlists.
		Query().
		Where(
			playlists.IDEQ(
				params.ID,
			),
			playlists.HasOwnerWith(
				users.IDEQ(
					s.GetUserID(),
				),
			),
		).
		WithMusics().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("remove musics from playlist query error: %w", err)
	}

	var forms RemoveMusicsFromPlaylistForm
	if err := s.ctx.ShouldBind(&forms); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	maps := make(map[uint64]bool)
	for _, v := range query.Edges.Musics {
		maps[v.ID] = true
	}

	conv := make([]uint64, 0, len(forms.MusicIDs))
	for _, str := range forms.MusicIDs {
		parse, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return fmt.Errorf("add musics to playlist parse music id error: %w", err)
		}
		conv = append(conv, parse)
	}

	var removes []*ent.Musics
	for _, id := range conv {
		if _, ok := maps[id]; ok {
			music, err := s.client.Musics.Get(s.ctx, id)
			if err != nil {
				return fmt.Errorf("get music error: %w", err)
			}
			removes = append(removes, music)
		}
	}
	if err := s.client.Playlists.Update().Where(playlists.IDEQ(query.ID)).RemoveMusics(removes...).Exec(s.ctx); err != nil {
		return fmt.Errorf("remove musics from playlist error: %w", err)
	}

	Success(s.ctx, gin.H{"code": http.StatusOK})
	return nil
}

type DelPlaylistParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

func (s *Server) DelPlaylist() error {
	defer s.client.Close()

	var params DelPlaylistParams
	if err := s.ctx.ShouldBindUri(&params); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}
	query, err := s.client.Playlists.
		Query().
		Where(
			playlists.IDEQ(
				params.ID,
			),
			playlists.HasOwnerWith(
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
		return fmt.Errorf("delete playlist query error: %w", err)
	}

	if err := s.client.Playlists.DeleteOne(query).Exec(s.ctx); err != nil {
		return fmt.Errorf("delete playlist error: %w", err)
	}

	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	return nil
}
