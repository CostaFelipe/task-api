package main

import (
	"fmt"
	"log"

	"github.com/CostaFelipe/task-api/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.DBDriver)
}
