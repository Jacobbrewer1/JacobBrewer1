package main

import "githubStats/api"

type repoStat struct {
	language string
	percent  int
}

func githubHandler() ([]api.Repository, error) {
	r, err := api.GetRepos()
	if err != nil {
		return nil, err
	}
	createLanguages(r)
	return nil, nil
}

func createLanguages(repos []api.Repository) {

}
