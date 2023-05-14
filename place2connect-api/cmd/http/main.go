package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fxfrancky/place2connect-api/config"
	_ "github.com/fxfrancky/place2connect-api/docs"
	"github.com/fxfrancky/place2connect-api/internal/handlers"
	"github.com/fxfrancky/place2connect-api/internal/initializers"
	postRepo "github.com/fxfrancky/place2connect-api/internal/repository/post"
	userRepo "github.com/fxfrancky/place2connect-api/internal/repository/user"
	"github.com/fxfrancky/place2connect-api/pkg/shutdown"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

// @title PLACE2CONNECT API
// @version 1.0
// @description PLACE2CONNECT Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name PLACE2CONNECT API Support
// @contact.email contact@owonafx.com
// @license.name pace2connect 2.0
// @schemes http https
// @produce application/json
// @consumes application/json
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {

	// Configure Swagger API
	swagg := swagger.Config{
		Title:        "Swagger API",
		DeepLinking:  false,
		DocExpansion: "none",
	}
	//
	// setup exit code for graceful shutdown
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	env, err := config.LoadConfig(config.ConfigDefaultName)
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}
	// Init All Databases
	postgresDB := initializers.LoadDatabases(env)
	// load config

	// Init All Handlers
	h := initApp(env, postgresDB)

	// run the server
	cleanup, err := run(env, swagg, h)
	// run the cleanup after the server is terminated
	defer cleanup()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}
	// ensure the server is shutdown gracefully & app runs
	shutdown.Gracefully()

}

func run(env *config.Config, swagg swagger.Config, h *handlers.Handler) (func(), error) {

	// h := initApp(env)
	app := h.NewRoutes(env, swagg)

	// start the server
	go func() {
		// create the fiber app
		log.Fatal(app.Listen(":8080"))
	}()

	// return a function to close the server and database
	return func() {
		app.Shutdown()
	}, nil
}

func initApp(env *config.Config, postgresDB *gorm.DB) *handlers.Handler {
	userRepo := userRepo.NewUserRepositoryImpl(postgresDB)
	postRepo := postRepo.NewPostRepositoryImpl(postgresDB)
	h := handlers.NewHandler(userRepo, postRepo)
	return h
}
