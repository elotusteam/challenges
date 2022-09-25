package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleLoginConfig oauth2.Config
}

var AppConfig Config

var RandomState = "random"

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func LoadConfig() {
	// Oauth configuration for Google
	AppConfig.GoogleLoginConfig = oauth2.Config{
		ClientID:     "1054035610324-3e4jngrebbnavqu54bu7m4kr3p27fkuq.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}
