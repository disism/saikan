package authx

import (
	"github.com/disism/saikan/ent"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ctx    *gin.Context
	client *ent.Client
}

func NewServer(ctx *gin.Context, client *ent.Client) *Server {
	return &Server{ctx: ctx, client: client}
}
