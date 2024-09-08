package controller

import (
	"go-clean-architecture/module/user/request"
	"go-clean-architecture/module/user/response"
	"go-clean-architecture/module/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.AuthService
}

func (ac *AuthController) Signup(c *gin.Context) {
	var req request.SignupRequest
	var res response.SignupResponse

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	res, err = ac.AuthService.Signup(c, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
