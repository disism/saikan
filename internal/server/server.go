package server

import (
	"github.com/disism/saikan/ent"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	Unauthorized = "Unauthorized"
)

type Server struct {
	ctx    *gin.Context
	client *ent.Client
}

func NewServer(ctx *gin.Context, client *ent.Client) *Server {
	return &Server{ctx: ctx, client: client}
}

func (s *Server) GetUserID() uint64 {
	const z = 0
	g, exists := s.ctx.Get(ContextUserIDKEY)
	if !exists {
		ErrorUnauthorized(s.ctx, Unauthorized)
		return z
	}

	id, err := strconv.ParseUint(g.(string), 10, 64)
	if err != nil {
		ErrorInternalServerError(s.ctx)
		return z
	}
	return id
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

const (
	InternalServerError = "InternalServerError"
	NoPermission        = "no permission"
)

func Success(c *gin.Context, resp any) {
	c.JSON(http.StatusOK, resp)
}

func (s *Server) Ok() {
	s.ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

func ErrorJSONResponse(c *gin.Context, status int, message string) {
	errResp := ErrorResponse{
		Code:  status,
		Error: message,
	}
	c.JSON(status, errResp)
}

func ErrorInternalServerError(c *gin.Context) {
	ErrorJSONResponse(c, http.StatusInternalServerError, InternalServerError)
	return
}

func ErrorConflict(c *gin.Context, message string) {
	ErrorJSONResponse(c, http.StatusConflict, message)
}

func ErrorNotFound(c *gin.Context, message string) {
	ErrorJSONResponse(c, http.StatusNotFound, message)
}

func ErrorUnauthorized(c *gin.Context, message string) {
	ErrorJSONResponse(c, http.StatusUnauthorized, message)
}

func ErrorBadRequest(c *gin.Context, message string) {
	ErrorJSONResponse(c, http.StatusBadRequest, message)
}

func ErrorNoPermission(c *gin.Context, message string) {
	ErrorJSONResponse(c, http.StatusForbidden, message)
}
