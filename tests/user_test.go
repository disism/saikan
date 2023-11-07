package tests

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/disism/saikan/conf"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/device"
	"github.com/disism/saikan/ent/user"
	"github.com/disism/saikan/internal/enthook"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestUserCreate(t *testing.T) {
	ctx := context.Background()
	client, err := ent.Open(
		dialect.SQLite,
		conf.SQLiteConnectURL(userDBName),
	)
	if err != nil {
		t.Errorf("failed to create ent client: %v", err)
		return
	}
	defer client.Close()
	if err := client.Schema.Create(ctx); err != nil {
		t.Errorf("failed to create schema resources: %v", err)
		return
	}
	enthook.IDHook(client)

	u, err := client.User.Create().
		SetUsername(username).
		Save(ctx)
	if err != nil {
		t.Errorf("failed to create authx: %v", err)
		return
	}
	t.Logf("authx:\n%+v", u)

	d, err := client.Device.
		Create().
		SetIP(ip).
		SetDevice(deviceName).
		SetUser(u).
		Save(ctx)
	if err != nil {
		t.Errorf("failed to create device: %v", err)
		return
	}
	t.Logf("device:\n%+v", d)
}

func TestQueryUsers(t *testing.T) {
	ctx := context.Background()
	client, err := ent.Open(
		dialect.SQLite,
		conf.SQLiteConnectURL(userDBName),
	)
	if err != nil {
		t.Errorf("failed to create ent client: %v", err)
		return
	}
	defer client.Close()

	byUsername, err := client.User.Query().
		Where(user.UsernameEQ(username)).
		Only(ctx)
	if err != nil {
		t.Errorf("failed to query authx by username: %v", err)
		return
	}
	t.Logf("query authx by username:\n%+v", byUsername)

	getUserByID, err := client.User.Get(ctx, userID)
	if err != nil {
		t.Errorf("failed to get authx by id: %v", err)
		return
	}
	t.Logf("get authx by id:\n%+v", getUserByID)
}

func TestQueryDevices(t *testing.T) {
	ctx := context.Background()
	client, err := ent.Open(
		dialect.SQLite,
		conf.SQLiteConnectURL(userDBName),
	)
	if err != nil {
		t.Errorf("failed to create ent client: %v", err)
		return
	}
	defer client.Close()

	d, err := client.Device.Query().
		Where(device.IDEQ(deviceID)).
		Only(ctx)
	if err != nil {
		t.Errorf("failed to query device by id: %v", err)
		return
	}
	t.Logf("query device:\n%+v", d)

	// QUERY USER DEVICES BY USER_ID
	queryDevicesByUserID, err := client.Device.Query().
		Where(
			device.HasUserWith(user.IDEQ(userID)),
		).
		All(ctx)
	if err != nil {
		t.Errorf("failed to query authx devices: %v", err)
		return
	}
	t.Logf("query authx devices:\n%+v", queryDevicesByUserID)
}
