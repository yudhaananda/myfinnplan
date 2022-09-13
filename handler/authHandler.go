package handler

import (
	"myfinnplan/formatter"
	"myfinnplan/helper"
	"myfinnplan/input"
	"myfinnplan/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService service.AuthService
	jwtService  service.JwtService
}

func NewAuthHandler(authService service.AuthService, jwtService service.JwtService) *authHandler {
	return &authHandler{authService, jwtService}
}

func (h *authHandler) RegisterUser(c *gin.Context) {
	var input input.UserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProfile, err := h.authService.RegisterUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Register Failed", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.jwtService.GenerateToken(newProfile.Id, newProfile.UserName)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := formatter.FormatUser(newProfile, token)

	response := helper.APIResponse("account has been registered successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *authHandler) Login(c *gin.Context) {
	var input input.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.authService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.jwtService.GenerateToken(loggedinUser.Id, loggedinUser.UserName)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := formatter.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Logedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
