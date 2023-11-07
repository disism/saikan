package authx

import (
	"fmt"
	"github.com/disism/saikan/conf"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/internal/server"
	"github.com/disism/saikan/jwt"
	"github.com/disism/saikan/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
)

func (s *Server) AuthN() error {
	defer s.client.Close()

	a, ok := server.NewContext(s.ctx).GetAuthorizationHeader()
	if !ok {
		return fmt.Errorf("NoAuthorizationHeader")
	}

	parse, err := jwt.Parser(a)
	if err != nil {
		slog.Error("authn parse id token error: ", err.Error())
		return err
	}
	if err := oidc.
		NewValidation(s.ctx).
		SetIssuer(parse.Issuer).
		SetAudience(parse.Audience).
		SetIDToken(a).
		Build().
		Validation(); err != nil {
		return fmt.Errorf("IDTokenUnauthorized")
	}

	query, err := func() (*ent.User, error) {
		ufmt := FMTUsername(parse.Subject, parse.Issuer)
		query, err := s.GetUserByUsername(ufmt)
		if err != nil {
			if ent.IsNotFound(err) {
				create, err := s.client.User.
					Create().
					SetUsername(ufmt).
					Save(s.ctx)
				if err != nil {
					slog.Error("authn create user error: ", err.Error())
					return nil, err
				}
				return create, nil
			}
			slog.Error("authn query user error: ", err.Error())
			return nil, err
		}

		return query, nil
	}()
	if err != nil {
		return err
	}

	device, err := s.client.Device.Create().
		SetUser(query).
		SetIP(s.ctx.ClientIP()).
		SetDevice(s.ctx.Request.UserAgent()).
		Save(s.ctx)
	if err != nil {
		slog.Error("authn create device error: ", err.Error())
		return err
	}
	// Putting the combined user information into the cache.
	device.Edges.User = query
	if err := SETDeviceCache(s.ctx, device.ID, device); err != nil {
		return err
	}
	k, err := jwt.New(
		conf.GetServerAddr(),
		strconv.FormatUint(query.ID, 10),
		jwt.WithID(strconv.FormatUint(device.ID, 10)),
	).Generate(
		conf.GetJWTSecret(),
	)
	if err != nil {
		slog.Error("authn generate jwt error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"access_token": k,
	})
	return nil
}
