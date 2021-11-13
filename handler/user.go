package handler

import (
	"net/http"
	"open_projects/helper"
	"open_projects/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	// token, err := h.jwtSerice.GenerateToken()

	formatter := user.FormatUser(newUser, "jwtTokenDahIntinya")

	response := helper.APIresponse("Account has been registerd", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
