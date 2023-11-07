package nodeinfo

import (
	"context"
	"github.com/disism/saikan/ent"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ctx    *gin.Context
	client *ent.Client
}

type Client struct {
	ctx    *context.Context
	client *ent.Client
}

func NewServer(ctx *gin.Context, client *ent.Client) *Server {
	return &Server{ctx: ctx, client: client}
}

func NewClient(ctx *context.Context, client *ent.Client) *Client {
	return &Client{ctx: ctx, client: client}
}

func (s *Client) AddOIDC(name, conf string) error {
	defer s.client.Close()

	if err := s.client.Oidc.
		Create().
		SetName(name).
		SetConfigurationEndpoint(conf).
		Exec(*s.ctx); err != nil {
		return err
	}
	return nil
}

func (s *Client) ListOIDC() ([]*ent.Oidc, error) {
	defer s.client.Close()

	all, err := s.client.Oidc.Query().All(*s.ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (s *Client) RemoveOIDC(id uint64) error {
	defer s.client.Close()

	if err := s.client.Oidc.DeleteOneID(id).Exec(*s.ctx); err != nil {
		return err
	}
	return nil
}

func (s *Server) GetAllOIDC() ([]*OidcProvider, error) {
	defer s.client.Close()

	all, err := s.client.Oidc.Query().All(s.ctx)
	if err != nil {
		return nil, err
	}
	var providers []*OidcProvider
	for _, provider := range all {
		providers = append(providers, &OidcProvider{
			Name:                  provider.Name,
			ConfigurationEndpoint: provider.ConfigurationEndpoint,
		})
	}
	return providers, nil
}
