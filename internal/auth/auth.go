package auth

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// TODO: Move this to the authservice package
func NewGithubConfig() *oauth2.Config {
	githubClientID := os.Getenv("GITHUB_CLIENT_ID")
	githubClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	config := &oauth2.Config{
		ClientID:     githubClientID,
		ClientSecret: githubClientSecret,
		Scopes:       []string{"user:email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  github.Endpoint.AuthURL,
			TokenURL: github.Endpoint.TokenURL,
		},
	}

	return config
}
