package main

import (
	"log"

	"ringhover-go/internal/config"
	"ringhover-go/internal/dao"
	"ringhover-go/internal/db"
	api "ringhover-go/internal/http"
	"ringhover-go/internal/http/handlers"
	"ringhover-go/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	// open DB if DSN, if not crash
	conn, err := db.Open(cfg.MySQLDSN)
	if err != nil {
		log.Fatal(err)
	}

	dao := dao.NewDao(conn)
	service := service.NewModelisationService(dao)
	handler := handlers.NewTaskHandler(service)

	r := api.NewRouter(handler)
	if err := r.Run(":" + cfg.HTTPPort); err != nil {
		log.Fatal(err)
	}
}
