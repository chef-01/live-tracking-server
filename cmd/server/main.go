package main

import (
	"log"
	"os"

	"github.com/chef-01/live-tracking-server/config"
	"github.com/chef-01/live-tracking-server/modules/user/data/repo_impl"
	"github.com/chef-01/live-tracking-server/modules/user/domain/usecase"
	"github.com/chef-01/live-tracking-server/modules/user/presentation/controller"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	"github.com/chef-01/live-tracking-server/modules/user/presentation/resolvers"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using default environment variables")
	}

	// Initialize Postgres connection
	config.InitPostgres()

	// Initialize repositories and usecase
	userRepo := repo_impl.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUsecase)

	// Initialize resolvers
	userResolver := resolvers.NewUserResolver(userController)

	// Setup Gin
	router := gin.Default()

	// Set up GraphQL Handler
	router.POST("/query", handler.NewDefaultServer(resolvers.NewExecutableSchema(resolvers.Config{Resolvers: userResolver})))
	router.GET("/playground", playground.Handler("GraphQL Playground", "/query"))

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server started at :%s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
