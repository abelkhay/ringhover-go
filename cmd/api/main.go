package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Server struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Récupérer le DSN
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN not set in .env")
	}

	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// DB connect
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		logger.Fatal("DB connect error", zap.Error(err))
	}

	defer db.Close()

	var n int
	if err := db.Get(&n, "SELECT COUNT(*) FROM tasks"); err != nil {
		log.Fatal("query error:", err)
	}
	fmt.Println("DB OK — tasks:", n)

	// server := &Server{db: db, logger: logger}

	// fmt.Println(server)

	// // Router
	// r := gin.Default()

	// r.GET("/tasks", s.getTasks)
	// r.GET("/tasks/:id/subtasks", s.getSubtasks)
	// r.POST("/tasks", s.createTask)
	// r.PATCH("/tasks/:id", s.updateTask)
	// r.DELETE("/tasks/:id", s.deleteTask)

	// Start server
	// log.Fatal(r.Run(":8080"))
}
