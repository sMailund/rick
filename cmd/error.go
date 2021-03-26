package cmd

import (
	"github.com/zmb3/spotify"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func shouldAttemptReauth(err error) bool {
	switch e := err.(type) {

	case spotify.Error:
		switch e.Status {
		case 401 | 403:
			return true
		default:
			return false
		}

	default:
		return false
	}
}
