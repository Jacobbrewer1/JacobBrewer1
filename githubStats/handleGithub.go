package main

import (
	"encoding/json"
	"errors"
	"githubStats/api"
	"githubStats/config"
	"io/ioutil"
	"log"
	"sort"
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
	for i, r := range repos {
		colour := colours[*r.Language]
		if colour == "" {
			colour = "#000000"
		}
		progress += "f'<span style=\"background-color: {color};'" +
			"f'width: {data.get(\"prop\", 0):0.3f}%;\" '" +
			"f'class=\"progress-item\"></span>'\"\""
		langList += "f\"\"\"" +
			"<li style=\"animation-delay: " + string(i*delay) + "ms;\">" +
			"<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"octicon\" style=\"fill:{color};\"" +
			"viewBox=\"0 0 16 16\" version=\"1.1\" width=\"16\" height=\"16\"><path" +
			"fill-rule=\"evenodd\" d=\"M8 4a4 4 0 100 8 4 4 0 000-8z\"></path></svg>" +
			"<span class=\"lang\">{lang}</span>" +
			"<span class=\"percent\">{data.get(\"prop\", 0):0.2f}%</span>" +
			"</li>"
	}

	if strings.Count(file, "{{ progress }}") > 1 {
		return errors.New("more than 1 {{ progress }} in template")
	}
	file = strings.ReplaceAll(file, "{{ progress }}", progress)
	if strings.Count(file, "{{ lang_list }}") > 1 {
		return errors.New("more than 1 {{ lang_list }} in template")
	}
	file = strings.ReplaceAll(file, "{{ lang_list }}", progress)
	return nil
}
