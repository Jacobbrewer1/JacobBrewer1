package main

import (
	"github.com/jacobbrewer1/githubStats/config"
	"log"
)

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Println(err)
		return
	}
	githubHandler()
}
