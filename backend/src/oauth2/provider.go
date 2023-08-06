package oauth2

import (
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/github"
)

var (
    githubOAuth2Config = oauth2.Config{
        ClientID:     "your-github-client-id",
        ClientSecret: "your-github-client-secret",
        RedirectURL:  "http://localhost:8080/oauth2/callback",
        Scopes:       []string{"read:user"},
        Endpoint:     github.Endpoint,
    }
)

func GetGitHubOAuth2Config() *oauth2.Config {
    return &githubOAuth2Config
}
