package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var (
    googleOauthConfig = &oauth.Config{
        RedirectURL: "http://localhost:8080/callback",
        ClientId: os.Getenv("GOOGLE_CLIENT_ID"),
        ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
        Scopes: []string{"https://www.googleapis.com/auth/user.info.email"},
        Endpoint: google.Endpoint,
        )
    // TODO: randomize it
    randonState = "random"



func main(){
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/login", handleLogin)
    http.HandleFunc("/callback", handleCallback)    
    http.ListenAndServe(":8080", nil)
}

func handleHome(w.http.ResponseWriter, r *http.Request){
    var html = `<html><body><a href="/login">Google Log In</a><body></html>`
    fmt.Fprint(w, html)
}

func handleLogin(w.http.ResponseWriter, r *http.Request){
    url := googleOauthConfig.AuthCodeUrl(randonState)
    http.RedirectURL(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w.http.ResponseWriter, r *http.Request){
    if r.FormValue("state") != randonState {
        fmt.Println("state is no valid")
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
    if err != nil {
        fmt.Printf("could not get token: %s\n", err.Error())
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }
    
    res, err := http.Get("http://www.googleapis.con/oauth2/v2/userinfo?access_token="+token.AccessToken)
    if err != nil {
        fmt.Printf("could not create get request: %s\n", err.Error())
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("could not parse response: %s\n", err.Error())
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    fmt.Fprint(w, "Response: %s", content)



}



