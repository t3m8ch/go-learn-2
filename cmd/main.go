package main

import (
	"context"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/t3m8ch/go-learn-2/internal/db"
	"github.com/t3m8ch/go-learn-2/internal/products"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()

	err := godotenv.Load()
	if err != nil {
		logger.Warn(".env file not found")
	}

	conn, err := db.InitDb(os.Getenv("DB_URL"), logger)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	r := gin.New()
	database := db.New(conn)

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	products.SetupRoutes(r, database, logger)

	r.Run()
}
