package server

import (
	"testing"
)

func TestFMTUsername(t *testing.T) {
	sub := "hvturingga"
	iss := "http://localhost:3033"
	t.Log(FMTUsername(sub, iss))
}
