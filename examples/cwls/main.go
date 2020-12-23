package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/karashiiro/godestone"
)

func main() {
	s, err := godestone.NewScraper()
	if err != nil {
		log.Fatalln(err)
	}

	c, err := s.FetchCWLS(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	cJSON, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(cJSON))
}
