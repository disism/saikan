package server

import (
	"github.com/disism/saikan/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	ContextUserIDKEY   = "USER_ID"
	ContextDeviceIDKEY = "DEVICE_ID"
)

type Context struct {
	*gin.Context
}

func NewContext(ctx *gin.Context) *Context {
	return &Context{
		Context: ctx,
	}
}

func (c *Context) GetAuthorizationHeader() (string, bool) {
	a := c.GetHeader("Authorization")
	if a == "" {
		return "", false
	}

	bearer := strings.Split(a, "Bearer ")
	if len(bearer) != 2 {
		return "", false
	}

	return bearer[1], true
}

func (c *Context) GetUserID() uint64 {
	g, exists := c.Get(ContextUserIDKEY)
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": errors.Unauthorized,
		})
	}

	id, err := strconv.ParseUint(g.(string), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errors.InternalServerError,
		})
	}
	return id
}

func (c *Context) ErrorInternalServer() {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":  http.StatusInternalServerError,
		"error": errors.InternalServerError,
	})
	return
}

func (c *Context) ErrorNoPermission() error {
	c.JSON(http.StatusForbidden, gin.H{
		"code":  http.StatusForbidden,
		"error": errors.NoPermission,
	})
	return nil
}

func (c *Context) ErrorUnauthorized() {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":  http.StatusUnauthorized,
		"error": errors.Unauthorized,
	})
	return
}

func (c *Context) ErrorUnauthorizedWithError(err error) error {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":  http.StatusUnauthorized,
		"error": err.Error(),
	})
	return nil
}

func (c *Context) ErrorUnprocessableEntity(err error) error {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"code":  http.StatusUnprocessableEntity,
		"error": err.Error(),
	})
	return nil
}

func (c *Context) ErrorNotFound(err error) error {
	c.JSON(http.StatusNotFound, gin.H{
		"code":  http.StatusNotFound,
		"error": err.Error(),
	})
	return nil
}

func (c *Context) ErrorConflict(err error) error {
	c.JSON(http.StatusConflict, gin.H{
		"code":  http.StatusConflict,
		"error": err.Error(),
	})
	return nil
}

func (c *Context) ErrorBadRequest(err error) error {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":  http.StatusBadRequest,
		"error": err.Error(),
	})
	return nil
}
