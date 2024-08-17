package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func NewGithubConfig() *oauth2.Config {
	config := &oauth2.Config{
		// ClientID:     GITHUB_CLIENT_ID,
		// ClientSecret: GITHUB_CLIENT_SECRET,
		Scopes: []string{"user:email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  github.Endpoint.AuthURL,
			TokenURL: github.Endpoint.TokenURL,
		},
	}

	return config
}
