package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	GithubApiToken string

	override *overrideStruct
)

type overrideStruct struct {
	GithubApiToken *string `json:"GithubApiToken"`
}

func ReadConfig() error {
	if exists := findFile("./config/override.json"); exists {
		log.Println("Override detected - Reading file")

		file, err := ioutil.ReadFile("./config/override.json")
		if err != nil {
			return err
		}

		log.Println(string(file))

		err = json.Unmarshal(file, &override)
		if err != nil {
			return err
		}

		GithubApiToken = *override.GithubApiToken
	} else {
		log.Println("No override detected. Using production config")
		GithubApiToken = os.Getenv("GITHUB_API_TOKEN")
	}
	return nil
}

func findFile(path string) bool {
	abs, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	log.Println(abs)

	file, err := os.Open(abs)
	if err != nil {
		return false
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return true
}
