package ipfs

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client interface {
	Post(path string) (*resty.Response, error)
}

type Web3Storage struct {
	BaseURL string
}

func (w *Web3Storage) Post(path string) (*resty.Response, error) {
	panic("implement me")
}

type Localhost struct {
	BaseURL string
}

func (l *Localhost) Post(path string) (*resty.Response, error) {
	client := resty.New()

	r, err := client.R().
		SetFile("path", path).
		Post(fmt.Sprintf("%s/api/v0/add?cid-version=1", l.BaseURL))

	if err != nil {
		return nil, err
	}

	return r, nil
}
