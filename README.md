# Task Manager with MCP Server

A **Task Manager** implemented in Go with a clean **Service → Repository → Database** architecture, exposed as tools via an **MCP server**.
This allows LLMs (e.g., Claude) to call into your task manager using the **Model Context Protocol**.

## Features

* CRUD operations for tasks (`Create`, `Update`, `List`, `Get`, `Delete`)
* Task domain with statuses (`todo`, `in-progress`, `done`)
* Repository layer for persistence (SQLite database by default)
* Service layer for business logic
* MCP server exposing task operations as tools

## Setup

### 1. Clone the repo

```bash
git clone https://github.com/Suhaan-Bhandary/mcp-task-manager.git
cd mcp-task-manager
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Configure database

By default, the project uses **SQLite** with `task-manager.db`.

If you want to switch to MySQL, edit `db/db.go`:

```go
dsn := "username:password@tcp(127.0.0.1:3306)/taskdb"
```

---

## Available Tools

The following MCP tools are exposed:

| Tool Name     | Description             | Input Schema                          |
| ------------- | ----------------------- | ------------------------------------- |
| `create_task` | Create a new task       | `{title, description, status}`        |
| `update_task` | Update an existing task | `{id, title?, description?, status?}` |
| `list_tasks`  | List all tasks          | `-`                                   |
| `get_task`    | Get task by ID          | `{id}`                                |
| `delete_task` | Delete a task by ID     | `{id}`                                |

---

## Example Task Object

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Write README",
  "description": "Document Task Manager with MCP",
  "status": "in-progress",
  "created_at": 1693740000,
  "updated_at": 1693743600
}
```


