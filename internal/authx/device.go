package authx

import (
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/device"
	"github.com/disism/saikan/ent/user"
	"github.com/disism/saikan/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
	"time"
)

type Device struct {
	ID         string    `json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	IP         string    `json:"ip"`
	Device     string    `json:"device"`
}

func (s *Server) GetDevices() error {
	defer s.client.Close()

	all, err := s.client.Device.
		Query().
		Where(
			device.
				HasUserWith(
					user.IDEQ(server.NewContext(s.ctx).GetUserID()),
				),
		).All(s.ctx)
	if err != nil {
		return err
	}
	r := make([]Device, len(all))
	for i, v := range all {
		r[i] = Device{
			ID:         strconv.FormatUint(v.ID, 10),
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			IP:         v.IP,
			Device:     v.Device,
		}
	}
	s.ctx.JSON(http.StatusOK, r)
	return nil
}

func (s *Server) DeleteDevice() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		slog.Error("delete device parse device id error: ", err.Error())
		return err
	}

	exist, err := s.client.Device.
		Query().
		Where(
			device.IDEQ(id),
			device.HasUserWith(
				user.IDEQ(
					server.NewContext(s.ctx).GetUserID(),
				),
			),
		).Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return server.NewContext(s.ctx).ErrorNoPermission()
		}
		slog.Error("delete device query error: ", err.Error())
		return err
	}

	if err := DELDeviceCache(s.ctx, id); err != nil {
		return err
	}

	if err := s.client.Device.DeleteOne(exist).Exec(s.ctx); err != nil {
		return err
	}

	s.ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
	return nil
}
