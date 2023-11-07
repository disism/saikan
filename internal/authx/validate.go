package authx

import (
	"fmt"
	"github.com/disism/saikan/conf"
	"github.com/disism/saikan/internal/errors"
	"github.com/disism/saikan/internal/server"
	"github.com/disism/saikan/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"strconv"
)

const (
	NoAuthorizationHeader      = "NoAuthorizationHeader"
	TokenAuthenticationFailure = "TokenAuthenticationFailure"
	TokenInvalid               = "TokenInvalid"
)

func ValidateJWT(c *gin.Context) error {
	token, ok := server.NewContext(c).GetAuthorizationHeader()
	if !ok {
		return fmt.Errorf(NoAuthorizationHeader)
	}

	validate, err := jwt.Validate(token, conf.GetJWTSecret())
	if err != nil {
		return fmt.Errorf(TokenAuthenticationFailure)
	}

	// Get the device id and check if the device exists.
	id, err := strconv.ParseUint(validate.ID, 10, 64)
	if err != nil {
		slog.Error("validate jwt parse id error: ", err.Error())
		return fmt.Errorf(errors.InternalServerError)
	}

	cache, err := GETDeviceCache(c, id)
	if err != nil {
		return err
	}

	c.Set(server.ContextUserIDKEY, validate.Subject)
	c.Set(server.ContextDeviceIDKEY, cache.ID)
	return nil
}
