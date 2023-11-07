package api_server

import (
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/music"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func AddMusic(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).AddMusic(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func GetMusics(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("get music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).GetMusics(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func GetMusic(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).GetMusic(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func EditMusic(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("edit music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).EditMusic(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func AddArtistsToMusic(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).AddArtistsToMusic(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func RemoveArtistsFromMusic(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).RemoveArtistsFromMusic(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func AddMusicToLibrary(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music to lib new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}

	if err := music.NewServer(c, client.Client).AddMusicToLibrary(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func AddMusicsToPlaylist(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).AddMusicsToPlaylist(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func RemoveMusicsFromPlaylist(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).RemoveMusicsFromPlaylist(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func DeleteMusic(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("delete music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}

	if err := music.NewServer(c, client.Client).DeleteMusic(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func CreateArtist(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("create music artist new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).CreateArtist(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func SearchArtist(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("search artist new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).SearchArtist(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func CreatePlaylist(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).CreatePlaylist(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func GetPlaylists(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).GetPlaylists(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func GetPlaylist(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).GetPlaylist(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func EditPlaylist(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).EditPlaylist(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func EditPlaylistImage(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).EditPlaylistImage(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func RemovePlaylistImage(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).RemovePlaylistImage(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func DeletePlaylist(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).DeletePlaylist(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func CreateAlbum(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).CreateAlbum(); err != nil {
		slog.Error("create album error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
}

func GetAlbums(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).GetAlbums(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func GetAlbum(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).GetAlbum(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func AddMusicsToAlbum(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).AddMusicsToAlbum(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func RemoveMusicsFromAlbum(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).RemoveMusicsFromAlbum(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func EditAlbum(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).EditAlbum(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func AddArtistsToAlbum(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).AddArtistsToAlbum(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func RemoveArtistsFromAlbum(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).RemoveArtistsFromAlbum(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}

func DeleteAlbum(c *gin.Context) {
	client, err := database.New(c)
	if err != nil {
		slog.Error("add music new database error: ", err.Error())
		server.NewContext(c).ErrorInternalServer()
	}
	if err := music.NewServer(c, client.Client).DeleteAlbum(); err != nil {
		server.NewContext(c).ErrorInternalServer()
	}
}
