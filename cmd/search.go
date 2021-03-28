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
	"strings"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for music",
	Long: `Search the spotify database for music. Supply flags to specify search for albums or playlists, defaults to song search if no flags are supplied.`,
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
	searchTerm := strings.Join(args, " ")

	toggleSong, _ := cmd.Flags().GetBool("song")
	toggleAlbum, _ := cmd.Flags().GetBool("album")
	togglePlaylist, _ := cmd.Flags().GetBool("playlist")
	toggleArtist, _ := cmd.Flags().GetBool("artist")

	fmt.Printf("searchterm: '%v'\n", searchTerm)
	fmt.Printf("song %v\n", toggleSong)
	fmt.Printf("album %v\n", toggleAlbum)
	fmt.Printf("playlist %v\n", togglePlaylist)
	fmt.Printf("artist %v\n", toggleArtist)
}