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

func Test_constructSearchType_multiple(t *testing.T) {
	searchType := constructSearchType(false, false, true, true)

	if searchType != spotify.SearchTypePlaylist|spotify.SearchTypeArtist {
		t.Errorf("constructsearchtype is not of expected searchType artist,playlist, got %v", searchTypeString(searchType))
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
