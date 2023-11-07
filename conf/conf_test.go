package conf

import (
	"testing"
)

func TestGetRedisAddr(t *testing.T) {
	t.Log(GetRedisAddr())
}

func TestGetRedisPassword(t *testing.T) {
	t.Log(GetRedisPassword())
}
