package authx

import (
	"fmt"
	"github.com/disism/saikan/conf"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/user"
	"github.com/disism/saikan/internal/server"
	"github.com/disism/saikan/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
)

func (s *Server) Login() error {
	defer s.client.Close()

	username := s.ctx.PostForm("username")
	password := s.ctx.PostForm("password")

	query, err := s.client.User.Query().Where(user.UsernameEQ(username)).Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorUnauthorizedWithError(fmt.Errorf("account does not exist"))
		}
		slog.Error("login query user error: ", err.Error())
		return err
	}

	// CompareHashAndPassword to compare whether the password is correct or not.
	if err := bcrypt.CompareHashAndPassword([]byte(query.Password), []byte(password)); err != nil {
		return server.NewContext(s.ctx).ErrorUnauthorizedWithError(fmt.Errorf("password verification failed"))
	}

	device, err := s.client.Device.Create().
		SetUser(query).
		SetIP(s.ctx.ClientIP()).
		SetDevice(s.ctx.Request.UserAgent()).
		Save(s.ctx)
	if err != nil {
		slog.Error("login create device error: ", err.Error())
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
		slog.Error("login generate jwt error: ", err.Error())
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"access_token": k,
	})

	return nil
}
