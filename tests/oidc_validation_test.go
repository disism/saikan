package tests

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/disism/saikan/jwt"
	"testing"
)

const (
	ID_TOKEN = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImYyNjJhMzIxNDIxM2QxOTRjOTI5OTFkNjczNWIxNTNiIn0.eyJzdWIiOiJodnR1cmluZ2dhIiwicGljdHVyZSI6Imh0dHBzOi8vYXZhdGFycy5naXRodWJ1c2VyY29udGVudC5jb20vdS8zNTkyMDM4OT92PTQiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJodnR1cmluZ2dhIiwicHJvZmlsZSI6InNhaWthbi5jb20gZGV2ZWxvcGVyIiwiZW1haWwiOiJodnR1cmluZ2dhQGRpc2lzbS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF0X2hhc2giOiJlenlLb2pEOUJ1d1ZINnJjWFptR21RIiwiYXVkIjoiZGlzaXNtX2NsaWVudCIsImV4cCI6MTY5ODI4ODg5OSwiaWF0IjoxNjk4Mjg1Mjk5LCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjMwMzMifQ.CIN2Lq0dx2vP1M85S7T8KQ-oRZylpmDZt-5ULUjyqf0TnmHU2Yh1TZiq_w5yfdpI6vdd_6XYQ-Q5Kb662fTBEl6K9Jrk9LWgoS7ng59-Bkp1fmAlVEqV63CspbzVzyY6rY2ca5-uKKObUqr21xGUCGHAuoLFkgJ181mcYDtilUXE5goWpudZPrTvNpLdtBftgai3T9UryK_NczweuYHZaZFZOaRxCHoQIs8COAc5klu-REwP8-NTpVbWKQgUiUz-aMMzx8-6FjNNPSlryj0SFAHgUU9Pi3inZA1_8oeewenv_kXkc9P6v3KY5nt-HCeCQjU650yzdoC4c_8QaldDptPzmA-LDykaGVkxKkXfvgjYwkAby0nscPgekPc6AM49PgLXGhGJPAQ7MWmILzMEN-J9IXyLxql95HCSfQzR_fPQMjAd6htobU9MLjWuTmXJHVoK881FMYjGz8xAIl2Nu63vxvgOOLDD3YXGLdne4uesOtj0PkBltxuYxM6-zWR_`
)

type OIDC struct {
	ctx     context.Context
	IDToken string
	Issuer  string
}

func TestValidationJWK(t *testing.T) {

	// Resolving tokens Obtaining the issuer and audience of the claim
	parse, err := jwt.Parser(ID_TOKEN)
	if err != nil {
		t.Errorf("failed to parse jwt: %v", err)
		return
	}

	// Validating ID_TOKEN to the OIDC PROVIDER api-server
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, parse.Issuer)
	if err != nil {
		t.Errorf("oidc provider may not be supported: %v", err)
		return
	}
	var verifier = provider.Verifier(&oidc.Config{
		ClientID: parse.Audience[0],
	})
	verify, err := verifier.Verify(ctx, ID_TOKEN)
	if err != nil {
		// oidc provider may not be supported
		t.Errorf("oidc provider may not be supported: %v", err)
		return
	}

	t.Logf("verify: %v", verify)
	t.Logf("iss: %s", verify.Issuer)
	t.Logf("aud: %s", verify.Audience)
	t.Logf("sub: %s", verify.Subject)
}
