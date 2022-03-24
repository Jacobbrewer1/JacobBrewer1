package main

import (
	"github.com/jacobbrewer1/githubStats/config"
	"log"
	"testing"
)

func TestGithubHandler(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test 1"},
	}

	config.ReadConfig()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := githubHandler()
			if err != nil {
				t.Errorf("err = %v, expected nil", err)
			}
			log.Println(r)
		})
	}
}
