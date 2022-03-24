package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jacobbrewer1/githubStats/api"
	"github.com/jacobbrewer1/githubStats/config"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type BySize []api.Repository

func (r BySize) Len() int {
	return len(r)
}

func (r BySize) Less(i, j int) bool {
	return *r[i].Size > *r[j].Size
}

func (r BySize) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

var colours map[string]string

func githubHandler() ([]api.Repository, error) {
	r, err := api.GetRepos()
	if err != nil {
		return nil, err
	}
	if err := createLanguages(r); err != nil {
		return nil, err
	}
	return nil, nil
}

func createLanguages(repos []api.Repository) error {
	if !config.FindFile("./assets/templates/languages.svg") {
		return errors.New("language file does not exist")
	}
	f, err := ioutil.ReadFile("./assets/templates/languages.svg")
	if err != nil {
		return err
	}
	file := string(f)
	log.Println(file)
	sort.Sort(BySize(repos))

	c, err := ioutil.ReadFile("./config/colours.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(c, &colours); err != nil {
		return err
	}

	var progress string
	var langList string
	delay := 150
	userSize := 0
	for _, r := range repos {
		userSize += *r.Size
	}

	var percent float64
	for i, r := range repos {
		if r.Language == nil {
			continue
		}
		percent = (float64(*r.Size) / float64(userSize)) * 100
		colour := colours[*r.Language]
		if colour == "" {
			colour = "#000000"
		}
		progress += "<span style=\"background-color: " + colour +
			";width: " + fmt.Sprintf("%f", percent) + "%\" " +
			"class=\"progress-item\"></span>"
		langList += "<li style=\"animation-delay: " + strconv.Itoa((i+1)*delay) + "ms;\">" +
			"<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"octicon\" style=\"fill:" + colour +
			";\" viewBox=\"0 0 16 16\" version=\"1.1\" width=\"16\" height=\"16\"><path " +
			"fill-rule=\"evenodd\" d=\"M8 4a4 4 0 100 8 4 4 0 000-8z\"></path></svg>" +
			"<span class=\"lang\">" + *r.Language + "</span>" +
			"<span class=\"percent\">" + fmt.Sprintf("%f", percent) + "%</span>" +
			"</li> \n \n"
		percent = 0
	}

	if strings.Count(file, "{{ progress }}") > 1 {
		return errors.New("more than 1 {{ progress }} in template")
	}
	file = strings.ReplaceAll(file, "{{ progress }}", progress)
	if strings.Count(file, "{{ lang_list }}") > 1 {
		return errors.New("more than 1 {{ lang_list }} in template")
	}
	file = strings.ReplaceAll(file, "{{ lang_list }}", langList)
	return nil
}
