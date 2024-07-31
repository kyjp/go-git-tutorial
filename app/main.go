package main

import (
	// "web/app/models"
	"web/app/repositories"
	"web/app/services"
	"web/app/controllers"
	"web/app/middlewares"
	"web/app/infra"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/gin-contrib/cors"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	r.Use(cors.Default())
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")
	itemRouter.GET("", itemController.FindAll)
	itemRouterWithAuth.GET("/:id", itemController.FindById)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)
	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	return r
}

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	r := setupRouter(db)
	// items := []models.Item{
	// 	{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
	// 	{ID: 2, Name: "商品1", Price: 2000, Description: "説明2", SoldOut: true},
	// 	{ID: 3, Name: "商品1", Price: 3000, Description: "説明3", SoldOut: false},
	// }

	// itemRepository := repositories.NewItemMemoryRepository(items)
	r.Run("0.0.0.0:8080")
}