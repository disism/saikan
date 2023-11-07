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
	"net/url"
	"strconv"
)

func (s *Server) GetUserByUsername(username string) (*ent.User, error) {
	return s.client.User.Query().Where(user.UsernameEQ(username)).Only(s.ctx)
}

func FMTUsername(sub, iss string) string {
	parse, err := url.Parse(iss)
	if err != nil {
		return fmt.Sprintf("%s:%s", sub, iss)
	}

	return fmt.Sprintf("%s:%s", sub, parse.Hostname())
}

func (s *Server) CreateUser() error {
	defer s.client.Close()

	username := s.ctx.PostForm("username")
	password := s.ctx.PostForm("password")

	exist, err := s.client.User.Query().Where(user.UsernameEQ(username)).Exist(s.ctx)
	if err != nil {
		return err
	}
	if exist {
		return server.NewContext(s.ctx).ErrorUnprocessableEntity(fmt.Errorf("UsernameExist"))
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	create, err := s.client.User.
		Create().
		SetUsername(username).
		SetPassword(string(hash)).
		Save(s.ctx)
	if err != nil {
		return err
	}

	device, err := s.client.Device.Create().
		SetUser(create).
		SetIP(s.ctx.ClientIP()).
		SetDevice(s.ctx.Request.UserAgent()).
		Save(s.ctx)
	if err != nil {
		slog.Error("authn create device error: ", err.Error())
		return err
	}
	// Putting the combined user information into the cache.
	device.Edges.User = create
	if err := SETDeviceCache(s.ctx, device.ID, device); err != nil {
		return err
	}
	k, err := jwt.New(
		conf.GetServerAddr(),
		strconv.FormatUint(create.ID, 10),
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
