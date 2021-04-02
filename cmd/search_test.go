package cmd

import (
	"github.com/zmb3/spotify"
	"strings"
	"testing"
)

func Test_constructSearchType_default(t *testing.T) {
	searchType := constructSearchType(false, false, false, false)

	if searchType != defaultSearchType {
		t.Errorf("expected %v, got %T", defaultSearchType, searchType)
	}
}

func Test_constructSearchType_album(t *testing.T) {
	searchType := constructSearchType(false, true, false, false)

	if searchType != spotify.SearchTypeAlbum {
		t.Errorf("constructsearchtype is not of expected searchType track, got %v", searchTypeString(searchType))
	}
}

func searchTypeString(st spotify.SearchType) string {
	types := []string{}
	if st&spotify.SearchTypeAlbum != 0 {
		types = append(types, "album")
	}
	if st&spotify.SearchTypeArtist != 0 {
		types = append(types, "artist")
	}
	if st&spotify.SearchTypePlaylist != 0 {
		types = append(types, "playlist")
	}
	if st&spotify.SearchTypeTrack != 0 {
		types = append(types, "track")
	}
	return strings.Join(types, ",")
}

func Test_verifyParams_none(t *testing.T) {
	flags := []bool{false, false, false, false}

	err := verifyParams(flags)

	if err != nil {
		t.Errorf("verifyParams returned unexpected error %v", err)
	}
}

func Test_verifyParams_one(t *testing.T) {
	flags := []bool{false, true, false, false}

	err := verifyParams(flags)

	if err != nil {
		t.Errorf("verifyParams returned unexpected error %v", err)
	}
}

func Test_verifyParams_multiple(t *testing.T) {
	flags := []bool{false, true, true, false}

	err := verifyParams(flags)

	if err == nil {
		t.Errorf("verifyParams did not return expected error")
	}
}

func Test_parseResults(t *testing.T) {

}

