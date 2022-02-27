package api

import (
	"githubStats/config"
	"log"
	"testing"
)

func TestGetRepos(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test 1"},
	}

	config.ReadConfig()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetRepos()
			if err != nil {
				t.Errorf("err = %v, expected nil", err)
			}
		})
	}
}

func TestRequestGithub(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
	}{
		{"1", "repos"},
	}

	config.ReadConfig()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := requestGithub(tt.endpoint)
			if err != nil {
				log.Println(err)
			}
		})
	}
}
