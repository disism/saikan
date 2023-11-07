package authx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/disism/saikan/cache"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/device"
	"github.com/disism/saikan/internal/database"
	"github.com/redis/go-redis/v9"
	"golang.org/x/exp/slog"
	"strconv"
	"time"
)

const (
	DeviceCacheDB = 1
)

func DeviceCacheKEY(id uint64) string {
	return fmt.Sprintf("%s_%s", "DEVICE", strconv.FormatUint(id, 10))
}

// SETDeviceCache sets the value of a device in the cache.
//
// It takes the following parameters:
//   - ctx: the context.Context object for the request.
//   - id: the ID of the device.
//   - v: the value to be set in the cache.
//
// It returns an error if there was an issue setting the value in the cache.
func SETDeviceCache(ctx context.Context, id uint64, v any) error {
	r := cache.NewRdbClient(DeviceCacheDB)

	marshal, err := json.Marshal(v)
	if err != nil {
		slog.Error("authn marshal device error: ", err.Error())
		return err
	}
	if err := r.Set(ctx, DeviceCacheKEY(id), marshal, time.Duration(14)*24*time.Hour).Err(); err != nil {
		return err
	}
	return nil
}

// GETDeviceCache retrieves data from the device cache.
//
// It takes a context and an ID as parameters.
// It returns the retrieved data and an error.
func GETDeviceCache(ctx context.Context, id uint64) (*ent.Device, error) {
	rdb := cache.NewRdbClient(DeviceCacheDB)
	r, err := rdb.Get(ctx, DeviceCacheKEY(id)).Result()
	if err != nil {
		if err == redis.Nil {
			db, err := database.New(ctx)
			if err != nil {
				return nil, err
			}
			defer db.Close()

			query, err := db.Client.
				Device.
				Query().
				Where(
					device.IDEQ(id),
				).
				WithUser().
				Only(ctx)
			if err != nil {
				slog.Error("exists device query error: ", err.Error())
				return nil, err
			}

			if err := SETDeviceCache(ctx, id, query); err != nil {
				return nil, err
			}
			return query, err
		}

		slog.Error("device cache get device error: ", err.Error())
		return nil, err
	}
	var d *ent.Device
	if err := json.Unmarshal([]byte(r), &d); err != nil {
		return nil, err
	}
	return d, nil
}

// DELDeviceCache deletes a device cache entry from the cache database.
//
// ctx is the context.Context that carries the deadline, cancellation signal, and other values across API boundaries.
// id is the uint64 identifier of the device cache entry to be deleted.
// Returns an error if there was a problem deleting the cache entry.
func DELDeviceCache(ctx context.Context, id uint64) error {
	rdb := cache.NewRdbClient(DeviceCacheDB)
	if err := rdb.Del(ctx, DeviceCacheKEY(id)).Err(); err != nil {
		return err
	}
	return nil
}
