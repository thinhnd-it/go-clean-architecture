package controller

import (
	"fmt"
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

func (ac *AuthController) Login(c *gin.Context) {
	var req request.LoginRequest
	var res response.LoginResponse

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	res, err = ac.AuthService.Login(c, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ac *AuthController) Profile(c *gin.Context) {
	userID := c.GetString("x-user-id")
	fmt.Println(userID)

	profile, err := ac.AuthService.GetUserByID(c, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
