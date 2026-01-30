package main

import (
	"context"
	"fmt"
	"log"

	"github.com/CostaFelipe/task-api/config"
	"github.com/CostaFelipe/task-api/internal/database"
	"github.com/CostaFelipe/task-api/internal/repository"
)

func main() {
	cfg, err := config.LoadConfig()
	ctx := context.Background()

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db.Ping())

	/*user := entity.User{
		Name:     "Jonh Doe",
		Email:    "jdoe@gmail.com",
		Password: "1234567",
	}*/

	//dueDate := time.Now().Add(48 * time.Hour)

	/*task := entity.Task{
		Title:       "Aprender Go",
		Description: "Estudar a linguagem Go.",
		Priority:    "low",
		DueDate:     &dueDate,
		UserID:      1,
	}*/

	//userDb := repository.NewUserRepositoy(db)
	taskDb := repository.NewTaskRepository(db)

	/*err = userDb.Create(ctx, &user)
	if err != nil {
		fmt.Println("inserir dado ao banco deu errado")
	}
	*/

	/*err = taskDb.Create(ctx, &task)
	if err != nil {
		fmt.Println("erro ao inserir task no banco deu errado")
	}*/

	taskID := 1
	userID := 1

	task, err := taskDb.FindByID(ctx, taskID, userID)
	if err != nil {
		fmt.Println("error ao buscar task")
	}

	fmt.Print(task)
}
