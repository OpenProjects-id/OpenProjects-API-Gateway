package main

import (
	"log"
	"net/http"
	"open_projects/auth"
	"open_projects/handler"
	"open_projects/helper"
	"open_projects/participation"
	"open_projects/project"
	"open_projects/transaction"
	"open_projects/user"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	projectRepository := project.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)
	participationRepository := participation.NewRepository(db)

	userService := user.NewService(userRepository)
	projectService := project.NewService(projectRepository)
	authService := auth.NewService()
	transactionService := transaction.NewService(transactionRepository, projectRepository)
	participationService := participation.NewService(participationRepository, projectRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	projectHandler := handler.NewProjectHandler(projectService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	participationHandler := handler.NewParticipationHandler(participationService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/projects", projectHandler.GetProjects)
	api.GET("/projects/:id", projectHandler.GetProject)
	api.POST("/projects", authMiddleware(authService, userService), projectHandler.CreateProject)
	api.PUT("/projects/:id", authMiddleware(authService, userService), projectHandler.UpdateProject)
	api.POST("/project-images", authMiddleware(authService, userService), projectHandler.UploadImage)

	api.GET("/projects/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetProjectTransactions)
	api.GET("/projects/:id/participations", authMiddleware(authService, userService), participationHandler.GetProjectParticipations)

	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.GET("/participations", authMiddleware(authService, userService), participationHandler.GetUserParticipations)
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
