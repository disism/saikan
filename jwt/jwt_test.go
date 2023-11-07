package jwt

import (
	"testing"
)

const (
	TOKEN  = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJkaXNpc20uY29tIiwic3ViIjoiaHZ0dXJpbmdnYSIsImF1ZCI6WyJodHRwczovL2Rpc2lzbS5jb20iXSwiaWF0IjoxNjk4Mzk3OTU1fQ.aNhpBO2Zo2wuMBgEjirh8EUqp3OA8VlNn3yDQf11sHw`
	SECRET = "12345678901234"
)

func TestGenerateToken(t *testing.T) {
	// Test case: Generate token with a valid secret
	const (
		iss = "disism.com"
		sub = "hvturingga"
	)
	// Create a JWT instance
	j := New(iss, sub, WithAudience([]string{"https://disism.com"}))
	token, err := j.Generate(SECRET)
	if err != nil {
		t.Errorf("error generating token: %v", err)
	}
	t.Logf("generated token: %s", token)
}

func TestJWTParse(t *testing.T) {
	claims, err := Parser(TOKEN)
	if err != nil {
		t.Errorf("error parsing token: %v", err)
	}
	t.Logf("%+v", claims)
}

func TestValidate(t *testing.T) {
	validate, err := Validate(TOKEN, SECRET)
	if err != nil {
		t.Error("failed to validate token")
		return
	}
	t.Logf("%+v", validate)
}
