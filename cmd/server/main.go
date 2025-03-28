package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/chef-01/live-tracking-server/config"
	"github.com/chef-01/live-tracking-server/modules/user/data/repo_impl"
	"github.com/chef-01/live-tracking-server/modules/user/domain/usecase"
	"github.com/chef-01/live-tracking-server/modules/user/presentation/controller"
	"github.com/gin-gonic/gin"

	"github.com/chef-01/live-tracking-server/modules/user/presentation/graphql/resolvers"

	"github.com/joho/godotenv"
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

	// Set up GraphQL Handler
	srv := handler.NewDefaultServer(resolvers.NewExecutableSchema(resolvers.Config{Resolvers: userResolver}))

	// Setup Gin
	router := gin.Default()

	// GraphQL query route
	router.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)  // Correct type conversion for Gin
	})

	// GraphQL playground route
	router.GET("/playground", func(c *gin.Context) {
		playground.Handler("GraphQL Playground", "/query").ServeHTTP(c.Writer, c.Request) // Correct type conversion for Gin
	})

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
