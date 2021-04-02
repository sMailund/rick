package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func getSpfyDir() string {
	curUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	homeDirectory := curUser.HomeDir
	return fmt.Sprintf("%v/.spfy", homeDirectory)
}

func persistJSON(unmarshalledJSON interface{}, location string) error {
	marshalledJSON, err := json.Marshal(unmarshalledJSON)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(location, marshalledJSON, 0644)
}

func tokenFileLocation() string {
	return fmt.Sprintf("%v/token.json", getSpfyDir())
}

func resultsFileLocation() string {
	return fmt.Sprintf("%v/results.json", getSpfyDir())
}

func persistSearchResults(results SearchResults) error {
	return persistJSON(results, resultsFileLocation())
}

func getSearchResults() (SearchResults, error) {
	resultsFile, err := os.Open(resultsFileLocation())
	searchResults := SearchResults{}
	if err != nil {
		return searchResults, err
	}
	defer resultsFile.Close()

	byteValue, _ := ioutil.ReadAll(resultsFile)
	err = json.Unmarshal(byteValue, &searchResults)

	return searchResults, err
}
