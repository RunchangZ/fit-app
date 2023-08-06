package oauth2

import (
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var (
    googleOAuth2Config = oauth2.Config{
        ClientID:     "your-google-client-id",
        ClientSecret: "your-google-client-secret",
        RedirectURL:  "http://localhost:8080/oauth2/callback",
        Scopes:       []string{"openid", "profile", "email"},
        Endpoint:     google.Endpoint,
    }
)

func GetGoogleOAuth2Config() *oauth2.Config {
    return &googleOAuth2Config
}
