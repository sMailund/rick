package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.
const redirectURI = "http://localhost:8080/callback"

var (
	auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadCurrentlyPlaying, spotify.ScopeUserReadPlaybackState, spotify.ScopeUserModifyPlaybackState)
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

var client *spotify.Client // TODO, hide from outer scope
var playerState *spotify.PlayerState

func Authenticate() {
	http.HandleFunc("/callback", completeAuth)
	go http.ListenAndServe(":8080", nil)

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client = <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)

	playerState, err = client.PlayerState()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found your %s (%s)\n", playerState.Device.Type, playerState.Device.Name)
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	err = persistToken(*tok)
	check(err)

	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}

func persistToken(token oauth2.Token) error {
	return persistJSON(token, tokenFileLocation())
}

func getAuthenticatedClient() spotify.Client {
	tokenfile, err := os.Open(tokenFileLocation())
	if os.IsNotExist(err) {
		fmt.Println("Could not find token file, sending to authentication...")
		Authenticate()
		return getAuthenticatedClient()
	} else if err != nil {
		panic(err)
	}
	defer tokenfile.Close()

	byteValue, _ := ioutil.ReadAll(tokenfile)
	var token oauth2.Token
	err = json.Unmarshal(byteValue, &token)
	check(err)

	return auth.NewClient(&token)
}

func getAuthenticatedClientWithRetry() spotify.Client {
	client := getAuthenticatedClient()

	if _, err := client.PlayerState(); err != nil {
		if shouldAttemptReauth(err) {
			fmt.Printf("spotify authentication failed with status %v, attempting reauth\n", err.(spotify.Error).Status)
			Authenticate()
			client = getAuthenticatedClient()
			if _, err := client.PlayerState(); err != nil {
				log.Fatal(err)
			} else {
				return client
			}
		} else {
			log.Fatal(err)
		}
	}
	return client
}
