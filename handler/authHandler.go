package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"myfinnplan/formatter"
	"myfinnplan/helper"
	"myfinnplan/input"
	"myfinnplan/service"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService service.AuthService
	jwtService  service.JwtService
}

func NewAuthHandler(authService service.AuthService, jwtService service.JwtService) *authHandler {
	return &authHandler{authService, jwtService}
}

func (h *authHandler) VerifiedUser(c *gin.Context) {
	tokenString := c.Param("token")

	failedTemplate, err := os.ReadFile("failed.html")
	if err != nil {
		c.Data(http.StatusOK, "text/html", failedTemplate)
		return
	}

	token, err := h.jwtService.ValidateToken(tokenString)

	if err != nil {
		c.Data(http.StatusOK, "text/html", failedTemplate)
		return
	}

	claim, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		c.Data(http.StatusOK, "text/html", failedTemplate)
		return
	}
	userId := int(claim["user_id"].(float64))

	dateTime, err := time.Parse(time.RFC3339Nano, claim["time"].(string))

	if err != nil {
		c.Data(http.StatusOK, "text/html", failedTemplate)
		return
	}

	if dateTime.Before(time.Now()) {
		c.Data(http.StatusOK, "text/html", failedTemplate)
		return
	}

	user, err := h.authService.VerifiedUser(userId)

	if err != nil {
		fmt.Println(err)
		c.Data(http.StatusOK, "text/html", failedTemplate)
		return
	}

	a := helper.Replace{
		Name: user.UserName,
	}

	t, err := template.ParseFiles("success.html")
	if err != nil {
		fmt.Println(err)
		c.Data(http.StatusOK, "text/html", failedTemplate)
		return
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, a); err != nil {
		fmt.Println(err)
		c.Data(http.StatusOK, "text/html", failedTemplate)
		return
	}
	successTemplate := buf.String()

	c.Data(http.StatusOK, "text/html", []byte(successTemplate))
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

	err = h.authService.SendEmail(newProfile, token)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Send Email Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
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
