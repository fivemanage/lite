package authservice

import (
	"context"
	"fmt"

	"github.com/fivemanage/lite/internal/auth"
	"github.com/uptrace/bun"
	"golang.org/x/oauth2"
)

type authConfig struct {
	github *oauth2.Config
}

type Auth struct {
	config authConfig
	db     *bun.DB
}

func New(db *bun.DB) *Auth {
	githubConfig := auth.NewGithubConfig()

	return &Auth{
		config: authConfig{
			github: githubConfig,
		},
		db: db,
	}
}

func (a *Auth) Login() string {
	verifier := oauth2.GenerateVerifier()
	url := a.config.github.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))

	return url
}

func (a *Auth) Callback(code string) *oauth2.Token {
	token, err := a.config.github.Exchange(context.TODO(), code)
	if err != nil {
		fmt.Println(err)
	}

	return token
}
