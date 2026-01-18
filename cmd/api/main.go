package main

import (
	"fmt"
	"log"

	"github.com/CostaFelipe/task-api/config"
	"github.com/CostaFelipe/task-api/internal/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db.Ping())
}
