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
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/user"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spfy",
	Short: "spotify for the terminal",
	Long:  `spfy is a CLI tool for controlling Spotify from the terminal`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.spfy.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getSpfyDir() string {
	curUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	homeDirectory := curUser.HomeDir
	return fmt.Sprintf("%v/.spfy", homeDirectory)
}

func TokenFileLocation() string {
	return fmt.Sprintf("%v/token.json", getSpfyDir())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	spfydir := getSpfyDir() // TODO maybe it's better to use viper for this?

	if _, err := os.Stat(spfydir); os.IsNotExist(err) {
		fmt.Println("Setting up necessary file structure...")
		if err := os.Mkdir(spfydir, 0755); err != nil {
			fmt.Println("could not set up necessary files to run project, panicking")
			panic(err)
		}
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".spfy" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".spfy")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
