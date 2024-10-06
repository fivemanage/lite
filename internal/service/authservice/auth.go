package authservice

import (
	"context"
	"fmt"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/auth"
	"github.com/fivemanage/lite/internal/crypt"
	"github.com/fivemanage/lite/internal/database"
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

// RegisterUser will register the user, not shit, but we need to make sure this is
// only for non-admin users. The actual admin user should have the defalt admin password
// unless changed in ENV, and then be prompted to change it. That should probably be required.
func (a *Auth) RegisterUser(ctx context.Context, register *api.RegisterRequest) {
	// Check if user exists
	exists := a.userExists(ctx, register.Email)
	if exists {
		fmt.Println("User already exists")
		return
	}
	// Create user
	err := a.createUser(ctx, register)
	if err != nil {
		fmt.Println(err)
	}

	// Return a JWT or smth
}

// Login uses email and password to authenticate the user
func (a *Auth) LoginUser() {
}

// OAuthLogin uses OAuth2 to authenticate the user
func (a *Auth) OAuthLogin() string {
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

func (a *Auth) userExists(ctx context.Context, email string) bool {
	user := new(database.User)
	err := a.db.NewSelect().Model(user).Where("email = ?", email).Scan(ctx)
	if err != nil {
		fmt.Println(err)
	}

	if user.ID == 0 {
		return false
	}

	fmt.Println("User exists", user)

	return user.ID != 0
}

// CreateUser creates a new user with email and password
func (a *Auth) createUser(ctx context.Context, register *api.RegisterRequest) error {
	var err error
	hash, err := crypt.HashPassword(register.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}

	user := &database.User{
		Email:        register.Email,
		PasswordHash: hash,
	}

	fmt.Println("Creating user", user)

	_, err = a.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
