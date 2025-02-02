package main

import (
	"log"

	"github.com/ZnarKhalil/expense-app/config"
	"github.com/ZnarKhalil/expense-app/handler"
	"github.com/ZnarKhalil/expense-app/middleware"
	"github.com/ZnarKhalil/expense-app/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error reading .env file")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	models.AutoMigrate(db)

	r := gin.Default()

	authHandler := handler.NewAuthHandler(db)
	categoryHandler := handler.NewCategoryHandler(db)
	expenseHandler := handler.NewExpenseHandler(db)

	// Public routes
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/refresh", authHandler.RefreshToken)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware(db))
	protected.POST("/logout", authHandler.Logout)

	protected.GET("/categories", categoryHandler.GetCategories)
	protected.POST("/categories", categoryHandler.CreateCategory)
	protected.PUT("/categories/:id", categoryHandler.UpdateCategory)
	protected.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	protected.GET("/expenses", expenseHandler.GetExpenses)
	protected.POST("/expenses", expenseHandler.CreateExpense)
	protected.PUT("/expenses/:id", expenseHandler.UpdateExpense)
	protected.DELETE("/expenses/:id", expenseHandler.DeleteExpense)

	r.Run(":8080")
}
