/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
	"strings"
)

const defaultSearchType spotify.SearchType = spotify.SearchTypeTrack

// SearchResults represents the a collection of search results obtained from a single search through the  search functionality.
// These results are altered and simplified from the response obtained from the Spotify API,
// in order to be easier to work with in terms of displaying and using search results.
type SearchResults struct {
	Tracks    []SearchResultEntry `json:"tracks"`
	Albums    []SearchResultEntry `json:"albums"`
	Playlists []SearchResultEntry `json:"playlists"`
	Artists   []SearchResultEntry `json:"artists"`
}

// SearchResultEntry represents a single row in the search results.
type SearchResultEntry struct {
	// DisplayNamePart1 represents the main textual identifier for the results,
	// such as song title.
	// When printing, this is the part printed before the dash
	DisplayNamePart1 string `json:"display_name_part_1"`

	// DisplayNamePart2 represents the auxiliary textual information for the results,
	// such as artists for a song.
	// When printing, this is the part printed after the dash
	DisplayNamePart2 []string `json:"display_name_part_2"`

	// URI is the spotify-specific resource identifier.
	// These are used in communication with the API, such as specifying the song to play.
	URI string `json:"uri"`
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for music",
	Long:  `Search the spotify database for music. Supply flags to specify search for albums or playlists, defaults to song search if no flags are supplied.`,
	Run: func(cmd *cobra.Command, args []string) {
		search(cmd, args)
	},
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchCmd.Flags().BoolP("song", "s", false, "Toggle song search")
	searchCmd.Flags().BoolP("album", "a", false, "Toggle album search")
	searchCmd.Flags().BoolP("playlist", "p", false, "Toggle playlist search")
	searchCmd.Flags().BoolP("artist", "r", false, "Toggle artist search")
}

func search(cmd *cobra.Command, args []string) {
	client := getAuthenticatedClientWithRetry()

	searchTerm := strings.Join(args, " ")

	toggleSong, _ := cmd.Flags().GetBool("song")
	toggleAlbum, _ := cmd.Flags().GetBool("album")
	togglePlaylist, _ := cmd.Flags().GetBool("playlist")
	toggleArtist, _ := cmd.Flags().GetBool("artist")

	searchType := constructSearchType(toggleSong, toggleAlbum, togglePlaylist, toggleArtist)

	searchResults, err := client.Search(searchTerm, searchType)
	check(err)

	printable := parseResults(*searchResults)

	// TODO persist and display results

	fmt.Printf("%v\n", printable)
}

func constructSearchType(toggleSong bool, toggleAlbum bool, togglePlaylist bool, toggleArtist bool) spotify.SearchType {
	var searchSong spotify.SearchType = 0
	var searchAlbum spotify.SearchType = 0
	var searchPlaylist spotify.SearchType = 0
	var searchArtist spotify.SearchType = 0

	if toggleSong {
		searchSong = spotify.SearchTypeTrack
	}

	if toggleAlbum {
		searchAlbum = spotify.SearchTypeAlbum
	}

	if togglePlaylist {
		searchPlaylist = spotify.SearchTypePlaylist
	}

	if toggleArtist {
		searchArtist = spotify.SearchTypeArtist
	}

	searchType := searchSong | searchAlbum | searchPlaylist | searchArtist

	if searchType == 0 {
		searchType = defaultSearchType
	}

	return searchType
}

func extractArtistNames(artists []spotify.SimpleArtist) []string {
	artistNames := []string{}
	for _, artist := range artists {
		artistNames = append(artistNames, artist.Name)
	}

	return artistNames
}

func parseResults(result spotify.SearchResult) SearchResults {
	parsedResults := SearchResults{}

	var tracks []SearchResultEntry
	for _, track := range result.Tracks.Tracks {
		tracks = append(tracks, SearchResultEntry{
			DisplayNamePart1: track.Name,
			DisplayNamePart2: extractArtistNames(track.Artists),
			URI:              string(track.URI),
		})
	}

	parsedResults.Tracks = tracks

	return parsedResults
}
