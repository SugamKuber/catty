package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"
	"server/internal/types"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *types.Database) *gin.Engine {
	r := gin.Default()

	r.GET("/h", handlers.Health)

	r.POST("/signup", func(c *gin.Context) {
		handlers.Signup(c, db)
	})
	r.POST("/login", func(c *gin.Context) {
		handlers.Login(c, db)
	})

	authRoutes := r.Group("/auth")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/me", func(c *gin.Context) {
			handlers.GetUserData(c, db)
		})
	}

	speechRoutes := r.Group("/speech")
	speechRoutes.Use(middleware.AuthMiddleware())
	{
		speechRoutes.POST("/upload", func(c *gin.Context) {
			handlers.Speech(c)
		})
	}

	reportRoutes := r.Group("/report")
	reportRoutes.Use(middleware.AuthMiddleware())
	{
		reportRoutes.GET("/generate", func(c *gin.Context) {
			handlers.Generate(c, db)
		})
	}

	return r
}
