package authservice

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

type authConfig struct {
	github *oauth2.Config
}

type Auth struct {
	config authConfig
}

func NewWithConfig(gitubConfig *oauth2.Config) *Auth {
	return &Auth{
		config: authConfig{
			github: gitubConfig,
		},
	}
}

func (r *Auth) Login() string {
	verifier := oauth2.GenerateVerifier()
	url := r.config.github.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))

	return url
}

func (r *Auth) Callback(code string) string {
	token, err := r.config.github.Exchange(context.TODO(), code)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(token.AccessToken)

	return token.AccessToken
}
