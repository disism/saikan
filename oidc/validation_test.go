package oidc

import (
	"context"
	"github.com/disism/saikan/jwt"
	"testing"
)

const (
	IdToken = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImYyNjJhMzIxNDIxM2QxOTRjOTI5OTFkNjczNWIxNTNiIn0.eyJzdWIiOiJodnR1cmluZ2dhIiwicGljdHVyZSI6Imh0dHBzOi8vYXZhdGFycy5naXRodWJ1c2VyY29udGVudC5jb20vdS8zNTkyMDM4OT92PTQiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJodnR1cmluZ2dhIiwicHJvZmlsZSI6InNhaWthbi5jb20gZGV2ZWxvcGVyIiwiZW1haWwiOiJodnR1cmluZ2dhQGRpc2lzbS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF0X2hhc2giOiJtM0IyM0pZMTNzVk05OTJPR3pBNTZRIiwiYXVkIjoiZGlzaXNtX2NsaWVudCIsImV4cCI6MTY5ODM3OTgzNiwiaWF0IjoxNjk4Mzc2MjM2LCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjMwMzMifQ.IG5J9PD9f2tv2nXgMY8Tx1Mo743Eq4-cbNHIXccf8edSV_7vyUd3EtAIGCor9865T1JhIFnWRafZg-zscvNQxFukmWtTT0sgnuuTf1RjskfxetDOUXOnlJp7CTPAV9uRIrsKY-0HtBTDbiH94kvyMarRFiekbgk_lcF5JcPEfkZeeuB82pNwr9nzo2d0bPF0vxOuuax2UBMeJCb_DG5DYcfaHRf9mmxvaYs2XB_c4nBrzILlQ5eRXnvtb-3sMFJMdMfF1WVQEkUGqCOiNzc92hMENFp5xo8QIo7x3F4T6wa4iqPeXR0BLBGikPMhCXAXDf6ID0DnARTJKESYFse8qYwrYVADm3s56NUWFS7kUu9mecjUhI8tmJWX68JA0Fl-P7zbJqTloKvla9V5t-2aprVcD2BGOvDhufSFnrcfT-ZERhGnkcc-jdPgzn0ruWr_pJxX2fwJRXAqAgdsh4zGTVVXRzDD7NGWWrfb5o2fBa_VY49csBugaMTS6oLX-p7h`
)

func TestValidation(t *testing.T) {
	// Resolving tokens Obtaining the issuer and audience of the claim
	parse, err := jwt.Parser(IdToken)
	if err != nil {
		t.Errorf("failed to parse jwt: %v", err)
		return
	}

	// Validating ID_TOKEN to the OIDC PROVIDER api-server
	ctx := context.Background()

	if err := NewValidation(ctx).
		SetIssuer(parse.Issuer).
		SetIDToken(IdToken).
		SetAudience(parse.Audience).
		Build().
		Validation(); err != nil {
		t.Errorf("failed to validate jwt: %v", err)
		return
	}
}
