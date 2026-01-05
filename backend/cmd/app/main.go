package main

import (
	_ "Backend/cmd/docs"
	"Backend/internal/middleware"
	db "Backend/internal/models"
	"Backend/internal/routes"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	DB      *sql.DB
	Server  *http.Server
	Logger  *zap.Logger
	Queries *db.Queries
}

// @title Sharobim Deacons API
// @version 1.0
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app := &App{}

	if err := app.setupLogger(); err != nil {
		log.Fatal(err)
	}
	defer app.closeLogger()

	loadConfig(app)

	if err := app.initDatabase(); err != nil {
		app.Logger.Fatal(err.Error())
	}

	router := setupRouter(app)

	handler := &routes.Handler{}

	handler.Setup(router)

	app.startServer(router)

	app.handleGracefulShutdown()
}

func (app *App) setupLogger() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	app.Logger = logger

	return nil

}

func loadConfig(app *App) {
	if err := godotenv.Load(); err != nil {
		app.Logger.Info("No .env file found or error loading .env. Assuming environment variables are set directly.")
	}

	ginMode := os.Getenv("MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	app.Logger.Info("Gin mode set", zap.String("mode", ginMode))
}

func (app *App) initDatabase() error {
	dbConnectionString := os.Getenv("MY_DATABASE_URL")
	if dbConnectionString == "" {
		return fmt.Errorf("MY_DATABASE_URL environment variable is not set")
	}

	var err error
	app.DB, err = sql.Open("postgres", dbConnectionString)
	app.Queries = db.New(app.DB)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = app.DB.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	app.DB.SetMaxOpenConns(25)
	app.DB.SetMaxIdleConns(10)
	app.DB.SetConnMaxLifetime(5 * time.Minute)
	app.DB.SetConnMaxIdleTime(1 * time.Minute)

	app.Logger.Info("Successfully connected to database.")
	return nil
}

func setupRouter(app *App) *gin.Engine {
	router := gin.New()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Use(middleware.ZapLogger(app.Logger))
	router.Use(middleware.InjectQueries(app.Queries))
	router.Use(middleware.InjectDatabase(app.DB))
	router.Use(gin.Recovery())

	if err := router.SetTrustedProxies(nil); err != nil {
		app.Logger.Fatal("Failed to set trusted proxies", zap.Error(err))
	}
	return router
}

func (app *App) startServer(router *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		app.Logger.Info("PORT environment variable not set, defaulting to ", zap.String("port", port))
	}
	app.Server = &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      router.Handler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		if err := app.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.Logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

}

func (app *App) handleGracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	app.Logger.Info("Shutdown signal received. Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if app.Server != nil {
		if err := app.Server.Shutdown(ctx); err != nil {
			app.Logger.Fatal("Server forced to shutdown", zap.Error(err))
		}
	}

	if app.DB != nil {
		app.Logger.Info("Closing database connection during graceful shutdown...")
		if err := app.DB.Close(); err != nil {
			app.Logger.Error("Error closing database connection gracefully: %v", zap.Error(err))
		} else {
			app.Logger.Info("Database connection closed successfully during graceful shutdown.")
		}
	}

	app.Logger.Info("Server exiting gracefully.")
}

func (app *App) closeLogger() {
	if err := app.Logger.Sync(); err != nil {
		log.Fatal(err)
	}
}
