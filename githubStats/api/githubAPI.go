package api

import (
	"encoding/json"
	"fmt"
	"github.com/jacobbrewer1/githubStats/config"
	"io/ioutil"
	"net/http"
)

func GetRepos() ([]Repository, error) {
	jsonRaw, err := requestGithub("user/repos")
	if err != nil {
		return nil, err
	}
	return decodeRepos(jsonRaw)
}

func decodeRepos(jsonRaw json.RawMessage) ([]Repository, error) {
	var repositories []Repository
	err := json.Unmarshal(jsonRaw, &repositories)
	return repositories, err
}

func requestGithub(endpoint string) (json.RawMessage, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/%v", endpoint), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %v", config.GithubApiToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
