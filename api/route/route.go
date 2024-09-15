package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/api/v1")

	// All Public APIs
	NewAuthRouter(db, publicRouter)
}
