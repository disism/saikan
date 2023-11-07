package server

import (
	"github.com/disism/godis/jwt"
	"github.com/disism/saikan/internal/conf"
	"golang.org/x/exp/slog"
	"strconv"
	"strings"
)

const (
	ContextUserIDKEY   = "UserID"
	ContextDeviceIDKEY = "DeviceID"
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

var (
	ErrNoAuthorizationHeader      = &ValidationError{Message: "no authorization header"}
	ErrTokenAuthenticationFailure = &ValidationError{Message: "token authentication failure"}
	ErrTokenInvalid               = &ValidationError{Message: "invalid authorization"}
	ErrUnauthorized               = &ValidationError{Message: "unauthorized"}
)

func (s *Server) ValidateJWT() error {
	a := s.ctx.GetHeader("Authorization")
	if a == "" {
		return ErrNoAuthorizationHeader
	}

	bearer := strings.Split(a, "Bearer ")
	if len(bearer) != 2 {
		return ErrTokenInvalid
	}

	jwtToken := jwt.New(
		jwt.SigningMethodHS256,
		jwt.WithIssuer(
			conf.GetServiceEndpoint(),
		),
	)
	validate, err := jwtToken.Validate(bearer[1], conf.GetJWTSecret())
	if err != nil {
		return ErrTokenAuthenticationFailure
	}

	id, err := strconv.ParseUint(validate.ID, 10, 64)
	if err != nil {
		slog.Error("validate jwt parse id error: ", err)
		return ErrUnauthorized
	}

	device, err := s.client.Devices.Get(s.ctx, id)
	if err != nil {
		slog.Error("validate jwt get device error: ", err)
		return ErrUnauthorized
	}

	s.ctx.Set(ContextUserIDKEY, validate.Subject)
	s.ctx.Set(ContextDeviceIDKEY, device.ID)
	return nil
}

//
//import (
//	"fmt"
//	"github.com/disism/saikan/conf"
//	"github.com/disism/saikan/ent"
//	"github.com/disism/saikan/jwt"
//	"github.com/disism/saikan/oidc"
//	"github.com/gin-gonic/gin"
//	"golang.org/x/exp/slog"
//	"net/http"
//	"strconv"
//)
//
//func (s *Server) AuthN() error {
//	defer s.client.Close()
//
//	a, ok := NewContext(s.ctx).GetAuthorizationHeader()
//	if !ok {
//		return fmt.Errorf("NoAuthorizationHeader")
//	}
//
//	parse, err := jwt.Parser(a)
//	if err != nil {
//		slog.Error("authn parse id token error: ", err.Error())
//		return err
//	}
//	if err := oidc.
//		NewValidation(s.ctx).
//		SetIssuer(parse.Issuer).
//		SetAudience(parse.Audience).
//		SetIDToken(a).
//		Build().
//		Validation(); err != nil {
//		return fmt.Errorf("IDTokenUnauthorized")
//	}
//
//	query, err := func() (*ent.Users, error) {
//		ufmt := FMTUsername(parse.Subject, parse.Issuer)
//		query, err := s.GetUserByUsername(ufmt)
//		if err != nil {
//			if ent.IsNotFound(err) {
//				create, err := s.client.Users.
//					Create().
//					SetUsername(ufmt).
//					Save(s.ctx)
//				if err != nil {
//					slog.Error("authn create user error: ", err.Error())
//					return nil, err
//				}
//				return create, nil
//			}
//			slog.Error("authn query user error: ", err.Error())
//			return nil, err
//		}
//
//		return query, nil
//	}()
//	if err != nil {
//		return err
//	}
//
//	device, err := s.client.Devices.Create().
//		SetUser(query).
//		SetIP(s.ctx.ClientIP()).
//		SetDevice(s.ctx.Request.UserAgent()).
//		Save(s.ctx)
//	if err != nil {
//		slog.Error("authn create device error: ", err.Error())
//		return err
//	}
//	// Putting the combined user information into the cache.
//	device.Edges.User = query
//	if err := SETDeviceCache(s.ctx, device.ID, device); err != nil {
//		return err
//	}
//	k, err := jwt.New(
//		conf.GetServerAddr(),
//		strconv.FormatUint(query.ID, 10),
//		jwt.WithID(strconv.FormatUint(device.ID, 10)),
//	).Generate(
//		conf.GetJWTSecret(),
//	)
//	if err != nil {
//		slog.Error("authn generate jwt error: ", err.Error())
//		return err
//	}
//
//	s.ctx.JSON(http.StatusOK, gin.H{
//		"code":         http.StatusOK,
//		"access_token": k,
//	})
//	return nil
//}
