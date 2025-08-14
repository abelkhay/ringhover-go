package main

import (
	"github.com/gin-gonic/gin"
    "go.uber.org/zap"

	"ringhover-go/internal/logging"
	"ringhover-go/internal/config"
	"ringhover-go/internal/dao"
	"ringhover-go/internal/db"
	api "ringhover-go/internal/http"
	"ringhover-go/internal/http/handlers"
	"ringhover-go/internal/service"

	"github.com/joho/godotenv"
)

func main() {

	defer func() { _ = logging.L().Sync() }()
	r := gin.New()
	r.Use(logging.ZapMiddleware())

	if err := godotenv.Load(".env"); err != nil {
		logging.L().Fatal("failed to load .env", zap.Error(err))
	}
	cfg := config.Load()

	// open DB if DSN, if not crash
	conn, err := db.Open(cfg.MySQLDSN)
	if err != nil {
		logging.L().Fatal("failed to open DB", zap.Error(err))
	}


	daoClient := dao.NewDao(conn)
	serviceClient := service.NewModelisationService(daoClient)
	handlerClient := handlers.NewTaskHandler(serviceClient)

	r = api.NewRouter(handlerClient)
	if err := r.Run(":" + cfg.HTTPPort); err != nil {
		logging.L().Fatal("server failed", zap.Error(err))
	}
}
