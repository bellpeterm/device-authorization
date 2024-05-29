package device

import (
	"fmt"

	"golang.org/x/oauth2"
)

type Entra struct{}

func (a Entra) Config(org, clientID string) *Config {
	return &Config{
		OAuth2Config: oauth2.Config{
			ClientID: clientID,
			Endpoint: oauth2.Endpoint{
				AuthURL:  fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/devicecode", org),
				TokenURL: fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", org),
			},
			Scopes: DefaultScopes,
		},
		Audience: fmt.Sprintf("api://%s", clientID),
		Issuer:   fmt.Sprintf("https://sts.windows.net/%s", org),
		KeyURI:   fmt.Sprintf("https://login.microsoftonline.com/%s/.well-known/openid-configuration", org),
	}
}
