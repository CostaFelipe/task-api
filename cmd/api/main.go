package main

import (
	"fmt"
	"log"

	"github.com/CostaFelipe/task-api/config"
	"github.com/CostaFelipe/task-api/internal/database"
	"github.com/CostaFelipe/task-api/internal/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	//ctx := context.Background()

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

	//user := entity.NewUser("Jonny", "teste@gmail.com", "123456")

	//dueDate := time.Now().Add(48 * time.Hour)

	/*task := entity.Task{
		Title:       "Aprender Go",
		Description: "Estudar a linguagem Go.",
		Priority:    "low",
		DueDate:     &dueDate,
		UserID:      1,
	}*/

	//userDb := repository.NewUserRepositoy(db)
	//taskDb := repository.NewTaskRepository(db)

	/*err = userDb.Create(ctx, user)
	if err != nil {
		fmt.Println("inserir dado ao banco deu errado")
	}*/

	/*err = taskDb.Create(ctx, &task)
	if err != nil {
		fmt.Println("erro ao inserir task no banco deu errado")
	}*/

	//taskID := 1
	//userID := 1

	//err = taskDb.Delete(ctx, 1, 1)

	/*task, err := taskDb.FindByID(ctx, taskID, userID)
	if err != nil {
		fmt.Println("error ao buscar task")
	}*/

	/*fmt.Print("task:", task)

	filter := &dto.TaskFilter{
		Completed: nil,
		Priority:  nil,
		Page:      1,
		Limit:     1,
	}

	tasks, total, err := taskDb.FindAllByUserID(ctx, userID, filter)
	if err != nil {
		log.Fatal(err)
	}
	*/

	/*fmt.Println()

	fmt.Printf("Total de tarefas: %d\n", total)
	for _, task := range *tasks {
		fmt.Printf("Tarefa: %+v\n", task)
	}
	*/
	cmd := middleware.NewAuthMiddleware(*cfg)
	token, err := cmd.GenerateToken(1, "jdoe@gmail.com")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(token)
}
