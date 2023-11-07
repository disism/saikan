package api_server

import (
	"github.com/gin-gonic/gin"
)

func Run() error {
	r := gin.Default()

	r.Use(CORS())

	r.GET("/version", GetVersionHandler)
	r.GET("/.well-known/node-info", GetNodeInfoHandler)

	//r.GET("/authn", AuthN)

	r.POST("/users/create", CreateUserHandler)
	r.POST("/login", LoginHandler)

	r.Use(AuthMW())

	devices := r.Group("/_devices/v1")
	{
		devices.GET("", ListDevicesHandler)
		devices.DELETE("/:id", DeleteDeviceHandler)
	}

	ipfs := r.Group("/_ipfs/v1")
	{
		ipfs.POST("/add", IPFSAddFilesHandler)
	}

	images := r.Group("/_images/v1")
	{
		images.POST("", AddImageHandler)
	}
	//
	musics := r.Group("/_streaming/v1/musics")
	{
		musics.POST("", AddMusicsHandler)
		musics.POST("/:id", SaveMusicHandler)
		musics.GET("", ListMusicsHandler)
		musics.GET("/:id", GetMusicHandler)
		musics.PUT("/:id", EditMusicHandler)
		musics.DELETE("/:id", DelMusicHandler)

		artists := musics.Group("/artists")
		{
			artists.POST("", CreateArtistHandler)
			artists.GET("", SearchArtistsHandler)
		}

		playlists := musics.Group("/playlists")
		{
			playlists.POST("", CreatePlaylistHandler)
			playlists.GET("", ListPlaylistsHandler)
			playlists.GET("/:id", GetPlaylistHandler)
			playlists.PUT("/:id", EditPlaylistHandler)
			playlists.PUT("/:id/images/remove", RemovePlaylistImageHandler)
			playlists.PUT("/:id/musics/add", AddMusicsToPlaylistHandler)
			playlists.PUT("/:id/musics/remove", RemoveMusicsFromPlaylistHandler)
			playlists.DELETE("/:id", DelPlaylistHandler)
		}

		albums := musics.Group("/albums")
		{
			albums.POST("", CreateAlbumHandler)
			albums.GET("", GetAlbumsHandler)
			albums.GET("/:id", GetAlbumHandler)
			albums.PUT("/:id/musics/add", AddMusicsToAlbumHandler)
			albums.PUT("/:id/musics/remove", RemoveMusicsFromAlbumHandler)
			albums.PUT("/:id", EditAlbumHandler)
			albums.DELETE("/:id", DelAlbumHandler)
		}
	}

	if err := r.Run(":8032"); err != nil {
		return err
	}
	return nil
}
