package route

import (
	"go-clean-architecture/module/user/controller"
	"go-clean-architecture/module/user/repository"
	"go-clean-architecture/module/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAuthRouter(db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)

	c := controller.AuthController{
		AuthService: *service.NewAuthService(ur),
	}

	group.POST("/signup", c.Signup)
	group.POST("/login", c.Login)
}
