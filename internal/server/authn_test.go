package server

//func TestAuthN(t *testing.T) {
//	client := resty.New()
//
//	resp, err := client.R().
//		EnableTrace().
//		SetHeader("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImYyNjJhMzIxNDIxM2QxOTRjOTI5OTFkNjczNWIxNTNiIn0.eyJzdWIiOiJodnR1cmluZ2dhIiwicGljdHVyZSI6Imh0dHBzOi8vYXZhdGFycy5naXRodWJ1c2VyY29udGVudC5jb20vdS8zNTkyMDM4OT92PTQiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJodnR1cmluZ2dhIiwicHJvZmlsZSI6InNhaWthbi5jb20gZGV2ZWxvcGVyIiwiZW1haWwiOiJodnR1cmluZ2dhQGRpc2lzbS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF0X2hhc2giOiJXdlh4U1Q5ck9uYUN0cUdjUEg3STJnIiwiYXVkIjoiZGlzaXNtX2NsaWVudCIsImV4cCI6MTY5ODYzMzc4NCwiaWF0IjoxNjk4NjMwMTg0LCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjMwMzMifQ.nsNjP-BROwUw_7giNWN-7lGGWWQoUzKO8O2sQEBSJ5hje8JvV2ZNWyUe4-MU8sCONRY1kpT52xEXWZx81GKYFHnEEpDjtCzJY1uptgRpqKFJ4qI-47SgAO7-PoG2uVsUMq7WTMwFYYLkkNWCjBk0rLfvGYZlebyTL-hps7qXjZLIzLTEdXhYb_LUyxEuFFcSFR45-k-locTK_EEdq53K8I4ugdYDkpH103PIvcJSaw5BksZiBcOuaPVsB1ciMQoBxfqaQzNT8h_mlLHnZGPcAwjTKJ4GzG_jRzXJYB-q6LVJ_0dL8X90BwVnsQpmFI0uYpMafWXdZdoIrhbKXAh-aA-e_bBP_NjL3u2N-TD_2HuSftBbPKQbfajSYHSsTiRVGwa_Z8aDsDx7XsyP08d3ldSThRX77l5xbOFd0TIlhvd9p_BpcSfge5tRdj1eTHwIZI_pKOMT0EqP0zFRV4ToIA2R3LSxAB-tXsiG8wpeU0WrtiJ3R6DFGFjAPq2lxKzp").
//		Get("http://localhost:8032/authn")
//	if err != nil {
//		t.Errorf("Error: %s", err)
//	}
//	t.Logf("Response: %s", resp)
//}
