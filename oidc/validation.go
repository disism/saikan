package oidc

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
)

// https://openid.net/specs/openid-connect-core-1_0.html#IDTokenValidation

type OIDC struct {
	ctx      context.Context
	IDToken  string
	Issuer   string
	Audience []string
}

type Builder struct {
	oidc *OIDC
}

func NewValidation(ctx context.Context) *Builder {
	return &Builder{
		oidc: &OIDC{
			ctx: ctx,
		},
	}
}

func (b *Builder) SetIDToken(token string) *Builder {
	b.oidc.IDToken = token
	return b
}

func (b *Builder) SetIssuer(issuer string) *Builder {
	b.oidc.Issuer = issuer
	return b
}

func (b *Builder) SetAudience(audience []string) *Builder {
	b.oidc.Audience = audience
	return b
}

func (b *Builder) Build() *OIDC {
	return b.oidc
}

// Validation performs validation of OIDC.
//
// It returns the audience of the token if the verification is successful,
// otherwise it returns an error.
func (o *OIDC) Validation() error {
	provider, err := oidc.NewProvider(o.ctx, o.Issuer)
	if err != nil {
		return err
	}

	for _, audience := range o.Audience {
		v := provider.Verifier(&oidc.Config{
			ClientID: audience,
		})
		_, err := v.Verify(o.ctx, o.IDToken)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("token verification failed for all audiences")
}
