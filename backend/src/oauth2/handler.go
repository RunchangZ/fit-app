package oauth2

import (
    "net/http"

    "golang.org/x/oauth2"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
    authURL := GetGoogleOAuth2Config().AuthCodeURL("state", oauth2.AccessTypeOffline)
    http.Redirect(w, r, authURL, http.StatusFound)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    token, err := GetGoogleOAuth2Config().Exchange(r.Context(), code)
    if err != nil {
        http.Error(w, "Error exchanging code for token", http.StatusInternalServerError)
        return
    }

    // Use the token to fetch user information or perform other actions

    // Example: Print the access token
    w.Write([]byte("Access Token: " + token.AccessToken))
}
