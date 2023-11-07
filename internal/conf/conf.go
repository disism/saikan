package conf

import (
	"os"
)

const (
	ServiceEndpoint        = "SERVICE_ENDPOINT"
	JWTSecret              = "JWT_SECRET"
	IPFSAPIEndpoint        = "IPFS_API_ENDPOINT"
	DefaultIPFSAPIEndpoint = "http://127.0.0.1:5001"
)

func GetServiceEndpoint() string {
	return os.Getenv(ServiceEndpoint)
}

func GetJWTSecret() string {
	return os.Getenv(JWTSecret)
}

func GetIPFSAPIEndpoint() string {
	addr := os.Getenv(IPFSAPIEndpoint)
	if addr != "" {
		return addr
	}
	return DefaultIPFSAPIEndpoint
}
