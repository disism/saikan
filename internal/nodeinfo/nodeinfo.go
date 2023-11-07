package nodeinfo

type OidcProvider struct {
	Name                  string `json:"name"`
	ConfigurationEndpoint string `json:"configuration_endpoint"`
}
