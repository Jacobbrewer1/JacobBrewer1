package main

import (
	"githubStats/config"
	"log"
)

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Println(err)
		return
	}
	githubHandler()
}
