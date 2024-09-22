package route

import (
	"go-clean-architecture/api/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, gin *gin.Engine) {
	// All Public APIs
	publicRouter := gin.Group("/api/v1")
	NewAuthRouter(db, publicRouter)

	// All Private APIs
	protectedRouter := gin.Group("/api/v1")
	protectedRouter.Use(middleware.JwtAuthMiddleware(os.Getenv("ACCESS_TOKEN_SECRET")))

	NewUserRouter(db, protectedRouter)
}
