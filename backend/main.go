package main

import (
    "fmt"
    "net/http"
    "github.com/RunchangZ/fit-app/oauth2"

)



func main() {
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/login", oauth2.HandleLogin)
    http.HandleFunc("/callback", oauth2.HandleCallback)

    fmt.Println("Server listening on :8080")
    http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<a href='/login'>Login with OAuth2</a>")
}