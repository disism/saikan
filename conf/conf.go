package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("../../conf")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal errors config file: %w", err))
	}
}

const (
	DatabaseName  = "saikan"
	RedisAddr     = "localhost:6379"
	RedisPassword = ""
)

func SQLiteConnectURL(name string) string {
	return fmt.Sprintf("file:%s.db?&cache=shared&_fk=1", name)
}

func GetRedisAddr() string {
	addr := viper.GetString("redis.addr")
	if strings.TrimSpace(addr) != "" {
		return addr
	}
	return RedisAddr
}

func GetRedisPassword() string {
	passwd := viper.GetString("redis.password")
	if strings.TrimSpace(passwd) != "" {
		return passwd
	}
	return RedisPassword
}

func GetJWTSecret() string {
	return viper.GetString("jwt.secret")
}

func GetServerAddr() string {
	return viper.GetString("server.addr")
}

func GetIPFSAddr() string {
	return viper.GetString("ipfs.addr")
}
