package main

import (
	"fmt"
	"log"

	"github.com/Suhaan-Bhandary/mcp-task-manager/db"
	"github.com/Suhaan-Bhandary/mcp-task-manager/repo"
	"github.com/Suhaan-Bhandary/mcp-task-manager/task"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("db error %s", err)
	}

	err = db.Migrate(database)
	if err != nil {
		log.Fatalf("error while migrating db: %s", err)
	}

	taskRepo := repo.NewTask(database)
	taskService := task.NewService(taskRepo)

	fmt.Printf("Setup Done %#v\n", taskService)
}
