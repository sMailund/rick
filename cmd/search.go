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
	"errors"
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
	Results    []SearchResultEntry `json:"results"`
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

	toggles := []bool{toggleSong, toggleAlbum, togglePlaylist, toggleArtist}
	err := verifyParams(toggles)
	check(err)

	searchType := constructSearchType(toggleSong, toggleAlbum, togglePlaylist, toggleArtist)

	searchResults, err := client.Search(searchTerm, searchType)
	check(err)

	parsedResults := parseResults(*searchResults, searchType)
	printResults(parsedResults)

	err = persistSearchResults(parsedResults)
	check(err)
}

func verifyParams(toggles []bool) error {
	alreadyFound := false

	for _, toggle := range toggles {
		if toggle {
			if alreadyFound {
				return errors.New("rick does not support multiple search types")
			} else {
				alreadyFound = true
			}
		}
	}

	return nil
}

func constructSearchType(toggleSong bool, toggleAlbum bool, togglePlaylist bool, toggleArtist bool) spotify.SearchType {
	if toggleSong {
		return spotify.SearchTypeTrack
	}

	if toggleAlbum {
		return spotify.SearchTypeAlbum
	}

	if togglePlaylist {
		return spotify.SearchTypePlaylist
	}

	if toggleArtist {
		return spotify.SearchTypeArtist
	}

	return defaultSearchType
}

func parseResults(results spotify.SearchResult, searchType spotify.SearchType) SearchResults {
	parsedResults := SearchResults{}

	switch searchType {
	case spotify.SearchTypeTrack:
		parsedResults.Results = parseTracks(results.Tracks)
	case spotify.SearchTypeAlbum:
		parsedResults.Results = parseAlbums(results.Albums)
	case spotify.SearchTypePlaylist:
		parsedResults.Results = parsePlaylists(results.Playlists)
	case spotify.SearchTypeArtist:
		parsedResults.Results = parseArtists(*results.Artists)
	}

	return parsedResults
}

func parseTracks(tracks *spotify.FullTrackPage) []SearchResultEntry {
	var parsedTracks []SearchResultEntry
	for _, track := range tracks.Tracks {
		parsedTracks = append(parsedTracks, SearchResultEntry{
			DisplayNamePart1: track.Name,
			DisplayNamePart2: extractArtistNames(track.Artists),
			URI:              string(track.URI),
		})
	}
	return parsedTracks
}

func parseAlbums(albums *spotify.SimpleAlbumPage) []SearchResultEntry {
	var parsedAlbums []SearchResultEntry
	for _, album := range albums.Albums {
		parsedAlbums = append(parsedAlbums, SearchResultEntry{
			DisplayNamePart1: album.Name,
			DisplayNamePart2: extractArtistNames(album.Artists),
			URI:              string(album.URI),
		})
	}
	return parsedAlbums
}

func parsePlaylists(playlists *spotify.SimplePlaylistPage) []SearchResultEntry {
	var parsedPlaylists []SearchResultEntry
	for _, playlist := range playlists.Playlists {
		parsedPlaylists = append(parsedPlaylists, SearchResultEntry{
			DisplayNamePart1: playlist.Name,
			DisplayNamePart2: extractPlaylistTracks(playlist.ID),
			URI:              string(playlist.URI),
		})
	}
	return parsedPlaylists
}

func parseArtists(artists spotify.FullArtistPage) []SearchResultEntry {
	var parsedArtists []SearchResultEntry
	for _, artist := range artists.Artists {
		parsedArtists = append(parsedArtists, SearchResultEntry{
			DisplayNamePart1: artist.Name,
			DisplayNamePart2: extractArtistsTracks(artist.ID),
			URI:              string(artist.URI),
		})
	}
	return parsedArtists

}

func extractArtistNames(artists []spotify.SimpleArtist) []string {
	artistNames := []string{}
	for _, artist := range artists {
		artistNames = append(artistNames, artist.Name)
	}

	return artistNames
}

func extractArtistsTracks(artistId spotify.ID) []string {
	client := getAuthenticatedClientWithRetry()
	// use US top tracks for laziness, not too important to choose the correct location
	artistsTopTracks, err := client.GetArtistsTopTracks(artistId, "US")

	if err != nil {
		// trouble fetching title of songs in playlist, better to display without playlist contents instead of crashing
		return []string{""}
	}

	artistNames := []string{}
	for _, track := range artistsTopTracks {
		artistNames = append(artistNames, track.Name)
	}

	return artistNames
}

func extractPlaylistTracks(playlistId spotify.ID) []string {
	client := getAuthenticatedClientWithRetry()
	playlistTracks, err := client.GetPlaylistTracks(playlistId)

	if err != nil {
		// trouble fetching title of songs in playlist, better to display without playlist contents instead of crashing
		return []string{""}
	}

	artistNames := []string{}
	for _, track := range playlistTracks.Tracks {
		artistNames = append(artistNames, track.Track.Name)
	}

	return artistNames
}

func printResults(results SearchResults) {
	for i, track := range results.Results {
		fmt.Printf("%v: %v - %v\n", i + 1, track.DisplayNamePart1, formatNamePart2(track.DisplayNamePart2))
	}
}

func formatNamePart2(names []string) string {
	return strings.Join(names, ", ")
}
