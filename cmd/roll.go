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
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

// rollCmd represents the roll command
var rollCmd = &cobra.Command{
	Use:   "roll",
	Short: "rick roll",
	Long: "serenade yourself witht the golden voice of Rick Astley",
	Run: func(cmd *cobra.Command, args []string) {
		roll()
	},
}

func init() {
	rootCmd.AddCommand(rollCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rollCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rollCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func roll() {
	client := getAuthenticatedClientWithRetry()
	searchResults, err := client.Search("never gonna give you up", spotify.SearchTypeTrack)
	check(err)

	opts := spotify.PlayOptions{
		DeviceID:        nil,
		PlaybackContext: nil,
		URIs:            nil,
		PlaybackOffset:  nil,
		PositionMs:      0,
	}

	uri := searchResults.Tracks.Tracks[0].URI
	opts.URIs = append(opts.URIs, uri)

	_ = client.PlayOpt(&opts)
}
