package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"testEffMobile/config"
	"testEffMobile/internal/handlers"
	"testEffMobile/internal/service"
	database "testEffMobile/packages/database/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cnfg := config.LoadConfig()

	var logLevel slog.Level
	if cnfg.LogLevel == config.DEBUG {
		logLevel = slog.LevelDebug
	} else {
		logLevel = slog.LevelInfo
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	}))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cnfg.DBhost, cnfg.DBuser, cnfg.DBpassword, cnfg.DBname, cnfg.DBport, cnfg.DBsslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get database instance: %v", err))
	}
	defer sqlDB.Close()
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Minute)

	err = sqlDB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	err = db.AutoMigrate(&database.User{})
	if err != nil {
		panic(fmt.Sprintf("Migration failed: %v", err))
	}

	genderApi := cnfg.GenderAPI
	ageApi := cnfg.AgeAPI
	nationalizeApi := cnfg.NationalizeAPI
	enricher := service.NewEnricherService(genderApi, ageApi, nationalizeApi, logger)

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	handlers.SetupRoutes(router, db, enricher, logger)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cnfg.AppPort),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Info("Starting server", "port", cnfg.AppPort)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
