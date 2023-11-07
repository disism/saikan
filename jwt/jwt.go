package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// https://datatracker.ietf.org/doc/html/rfc7519
// JWT https://datatracker.ietf.org/doc/html/rfc7519#section-4

var now = time.Now

type JWT struct {
	*jwt.RegisteredClaims
}

type Option func(j *JWT)

func WithAudience(audience []string) Option {
	return func(j *JWT) {
		j.Audience = audience
	}
}

func WithExpiresAt(expiresAt time.Time) Option {
	return func(j *JWT) {
		j.ExpiresAt = &jwt.NumericDate{
			Time: expiresAt,
		}
	}
}

func WithNotBefore(notBefore time.Time) Option {
	return func(j *JWT) {
		j.NotBefore = &jwt.NumericDate{
			Time: notBefore,
		}
	}
}

func WithIssuedAt(issuedAt time.Time) Option {
	return func(j *JWT) {
		j.IssuedAt = &jwt.NumericDate{
			Time: issuedAt,
		}
	}
}

func WithID(id string) Option {
	return func(j *JWT) {
		j.ID = id
	}
}

// New creates a new JWT instance with the provided options.
//
// It accepts variadic JwtOption arguments to configure the JWT instance.
// The function returns a pointer to the created JWT instance.
func New(iss, sub string, opts ...Option) *JWT {
	s := &JWT{
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:   iss,
			Subject:  sub,
			IssuedAt: &jwt.NumericDate{Time: now()},
		},
	}
	for _, srv := range opts {
		srv(s)
	}
	return s
}

// Generate generates a JWT token with the given secret.
//
// Parameters:
// - secret: The secret used to sign the token.
//
// Returns:
// - string: The generated JWT token.
// - errors: An errors if the token generation fails.
func (j *JWT) Generate(secret string) (string, error) {
	s, err := jwt.
		NewWithClaims(
			jwt.SigningMethodHS256,
			*j.RegisteredClaims,
		).
		SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

// Parser parses a JWT token and returns the registered claims.
//
// The function takes a `bearer` string as a parameter, which represents
// the JWT token. It returns a pointer to `jwt.RegisteredClaims` and an
// errors. The `jwt.RegisteredClaims` struct represents the registered
// claims of a JWT token.
func Parser(bearer string) (*jwt.RegisteredClaims, error) {
	t, _, err := new(jwt.Parser).ParseUnverified(bearer, &jwt.RegisteredClaims{})
	if err != nil {
		return nil, fmt.Errorf("parse token errors: %v", err)
	}
	claims, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("failed to get JWT claims")
	}

	return claims, nil
}

func Validate(token, secret string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %v", err)
	}
	return claims, nil
}
