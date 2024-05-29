package device

// var DefaultScopes = []string{"openid", "profile", "offline_access"}
var DefaultScopes = []string{"openid"}

type Provider interface {
	Config(org, clientID string) *Config
}

func NewProvider(name string) Provider {
	switch name {
	case "okta":
		return Okta{}
	case "auth0":
		return Auth0{}
	case "entra":
		return Entra{}
	default:
		panic("provider not found")
	}
}
