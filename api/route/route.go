package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")

	// All Public APIs
	NewAuthRouter(db, publicRouter)
}
