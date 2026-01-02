package main

import (
    "context"
    "log"
    "simx-go-todo/internal/config"
    "simx-go-todo/internal/share"
    "simx-go-todo/internal/todo"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

func main() {
    // Load .env for local development
    _ = godotenv.Load()

    // Initialize global DB connection using config/db.go logic
    if err := config.InitDB(); err != nil {
        log.Fatalf("failed to connect to db: %v", err)
    }

    r := gin.Default()

    // Register global middleware
    r.Use(share.GlobalMiddleware())
    
    // Initialize global DB connection
    if err := config.InitDB(); err != nil {
        log.Fatalf("failed to connect to db: %v", err)
    }

    // Set up database-backed repository and usecase
    repo, err := todo.NewTodoRepository()
    if err != nil {
        log.Fatalf("failed to connect to db: %v", err)
    }

    // AutoMigrate: create todos table if not exists
    if err := autoMigrateTodos(repo); err != nil {
        log.Fatalf("failed to migrate db: %v", err)
    }

    usecase := todo.NewTodoUsecase(repo)

    // Set up todo routes with dependency injection
    todo.RegisterRoutes(r, usecase)

    // Start server
    r.Run(":8080")
}

// autoMigrateTodos creates the todos table if it does not exist
func autoMigrateTodos(repoInterface interface{}) error {
    // Use the GetPool() method from todoRepo
    type pooler interface{ GetPool() *pgxpool.Pool }
    if r, ok := repoInterface.(pooler); ok {
        pool := r.GetPool()
        _, err := pool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS todos (
            id VARCHAR(64) PRIMARY KEY,
            title TEXT NOT NULL,
            done BOOLEAN NOT NULL DEFAULT FALSE
        )`)
        return err
    }
    return nil
}