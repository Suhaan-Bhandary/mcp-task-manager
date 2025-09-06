package main

import (
	"context"
	"flag"
	"log"

	"github.com/Suhaan-Bhandary/mcp-task-manager/db"
	"github.com/Suhaan-Bhandary/mcp-task-manager/mcp/tools"
	"github.com/Suhaan-Bhandary/mcp-task-manager/repo"
	"github.com/Suhaan-Bhandary/mcp-task-manager/task"
	_ "github.com/mattn/go-sqlite3"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	dbUrl := flag.String("db", "task-manager.db", "url to database")
	flag.Parse()

	database, err := db.InitDB(*dbUrl)
	if err != nil {
		log.Fatalf("db error %s", err)
	}

	err = db.Migrate(database)
	if err != nil {
		log.Fatalf("error while migrating db: %s", err)
	}

	taskRepo := repo.NewTask(database)
	taskService := task.NewService(taskRepo)
	taskToolHandler := tools.NewTaskHandler(taskService)

	server := mcp.NewServer(&mcp.Implementation{Name: "Task Manager", Version: "v1.0.0"}, nil)

	// Tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create",
		Description: "Create a new task",
	}, taskToolHandler.Create)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update",
		Description: "Update an existing task",
	}, taskToolHandler.Update)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list",
		Description: "List all tasks",
	}, taskToolHandler.List)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get",
		Description: "Get a task by ID",
	}, taskToolHandler.Get)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete",
		Description: "Delete a task by ID",
	}, taskToolHandler.Delete)

	log.Println("Starting MCP Server...")
	err = server.Run(context.Background(), &mcp.StdioTransport{})
	if err != nil {
		log.Fatal(err)
	}
}
