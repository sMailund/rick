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
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "play a song",
	Long: `play a selected song. Functions as "resume" when called without arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		playCommand(cmd)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	playCmd.Flags().IntP("result", "r", 0, "Play an entry from search results")
}

func playCommand(cmd *cobra.Command) {
	client := getAuthenticatedClientWithRetry()
	entryToPlay, err := cmd.Flags().GetInt("result")
	check(err)

	if entryToPlay > 0 {
		err = playEntry(client, entryToPlay)
	} else {
		err = play(client)
	}

	check(err)
}

func play(client spotify.Client) error {
	return client.Play()
}

func playEntry(client spotify.Client, entry int) error {
	opts := spotify.PlayOptions{
		DeviceID:        nil,
		PlaybackContext: nil,
		URIs:            nil,
		PlaybackOffset:  nil,
		PositionMs:      0,
	}

	results, err := getSearchResults()
	if err != nil {
		return err
	}

	if entry >= len(results.Results) {
		return errors.New("invalid results index")
	}

	uri := spotify.URI(results.Results[entry-1].URI)
	opts.URIs = append(opts.URIs, uri)

	return client.PlayOpt(&opts)
}