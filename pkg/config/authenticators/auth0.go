// platform/authenticator/auth.go

package authenticator

import (
	"context"
	"errors"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"os"
	"sync"
)

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

var onceAuth sync.Once
var Auth *Authenticator
var err error

// New instantiates the *Authenticator.
func InitAuth0() (*Authenticator, error) {
	onceAuth.Do(func() {
		provider, err := oidc.NewProvider(
			context.Background(),
			"https://"+os.Getenv("AUTH0_DOMAIN")+"/",
		)
		if err != nil {
			return
		}

		conf := oauth2.Config{
			ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
			ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("AUTH0_CALLBACK_URL"),
			Endpoint:     provider.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
		}
		Auth = &Authenticator{
			Provider: provider,
			Config:   conf,
		}
	})
	return Auth, err
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
