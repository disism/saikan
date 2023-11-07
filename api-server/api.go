package api_server

import (
	"github.com/gin-gonic/gin"
)

func Run() error {
	r := gin.Default()

	r.Use(CORS())

	r.GET("/version", GetVersion)
	r.GET("/.well-known/node-info", GetNodeInfo)

	r.POST("/users/create", CreateUser)
	r.POST("/login", Login)
	r.GET("/authn", AuthN)

	r.Use(AuthMW())

	devices := r.Group("/_devices/v1")
	{
		devices.GET("", GetDevices)
		devices.DELETE("/:id", DeleteDevice)
	}

	ipfs := r.Group("/_ipfs/v1")
	{
		ipfs.POST("/add", AddFiles)
	}

	image := r.Group("/_image/v1")
	{
		image.POST("", AddImages)
	}

	saved := r.Group("/_saved/v1")
	{
		saved.POST("", AddSaved)
		saved.GET("", GetSaves)
		saved.GET("/:id", GetSaved)
		saved.PUT("/:id", EditSaved)
		saved.DELETE("/:id", DeleteSaved)
		saved.PUT("/:id/link", LinkDir)
		saved.PUT("/:id/unlink", UnlinkDir)

		dirs := saved.Group("/dirs")
		{
			dirs.POST("", MKDir)
			dirs.GET("", ListDir)
			dirs.GET("/all", ListDirs)
			dirs.PATCH("/:id/name", RenameDir)
			dirs.PUT("/:id/mv/:dir_id", MVDir)
			dirs.DELETE("/:id", RMDir)
		}
	}

	musics := r.Group("/_streaming/v1/musics")
	{
		musics.POST("", AddMusic)
		musics.GET("", GetMusics)
		musics.GET("/:id", GetMusic)
		musics.PUT("/:id", EditMusic)
		musics.POST("/:id/artists", AddArtistsToMusic)
		musics.DELETE("/:id/artists", RemoveArtistsFromMusic)
		musics.POST("/:id/library", AddMusicToLibrary)
		musics.DELETE("/:id", DeleteMusic)

		// "/_streaming/v1/musics/artists"
		artists := musics.Group("/artists")
		{
			artists.POST("", CreateArtist)
			artists.GET("", SearchArtist)
		}

		// "/_streaming/v1/musics/playlists"
		playlists := musics.Group("/playlists")
		{
			playlists.POST("", CreatePlaylist)
			playlists.GET("", GetPlaylists)
			playlists.GET("/:id", GetPlaylist)
			playlists.PUT("/:id", EditPlaylist)
			playlists.PATCH("/:id/images", EditPlaylistImage)
			playlists.DELETE("/:id/images", RemovePlaylistImage)
			playlists.POST("/:id/musics", AddMusicsToPlaylist)
			playlists.PUT("/:id/musics", RemoveMusicsFromPlaylist)
			playlists.DELETE("/:id", DeletePlaylist)
		}

		// "/_streaming/v1/musics/albums"
		albums := musics.Group("/albums")
		{
			albums.POST("", CreateAlbum)
			albums.GET("", GetAlbums)
			albums.GET("/:id", GetAlbum)
			albums.POST("/:id/musics", AddMusicsToAlbum)
			albums.PUT("/:id/musics", RemoveMusicsFromAlbum)
			albums.PUT("/:id", EditAlbum)
			albums.POST("/:id/artists", AddArtistsToAlbum)
			albums.PUT("/:id/artists", RemoveArtistsFromAlbum)
			albums.DELETE("/:id", DeleteAlbum)
		}
	}

	if err := r.Run(":8032"); err != nil {
		return err
	}
	return nil
}
