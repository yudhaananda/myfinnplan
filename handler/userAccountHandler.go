package handler

import (
	"myfinnplan/entity"
	"myfinnplan/helper"
	"myfinnplan/input"
	"myfinnplan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userAccountHandler struct {
	userAccountService service.UserAccountService
}

func NewUserAccountHandler(userAccountService service.UserAccountService) *userAccountHandler {
	return &userAccountHandler{userAccountService}
}

func (h *userAccountHandler) CreateUserAccount(c *gin.Context) {
	var input input.UserAccountInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, ok := c.Get("currentUser")
	if !ok {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAccount, err := h.userAccountService.CreateUserAccount(input, userLogin.(entity.User))

	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Create UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Create UserAccount Success", http.StatusOK, "Success", userAccount)

	c.JSON(http.StatusOK, response)
}

func (h *userAccountHandler) EditUserAccount(c *gin.Context) {
	var input input.UserAccountEditInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, ok := c.Get("currentUser")
	if !ok {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAccount, err := h.userAccountService.EditUserAccount(input, userLogin.(entity.User))

	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Edit UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Edit UserAccount Success", http.StatusOK, "Success", userAccount)

	c.JSON(http.StatusOK, response)
}

func (h *userAccountHandler) GetAllUserAccounts(c *gin.Context) {
	userAccounts, err := h.userAccountService.GetAllUserAccount()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get All UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get All UserAccount Success", http.StatusOK, "Success", userAccounts)

	c.JSON(http.StatusOK, response)
}

func (h *userAccountHandler) GetUserAccountById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAccount, err := h.userAccountService.GetUserAccountById(idInt)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get UserAccount Success", http.StatusOK, "Success", userAccount)

	c.JSON(http.StatusOK, response)
}
func (h *userAccountHandler) GetUserAccountByAccountCode(c *gin.Context) {
	accountCode := c.Param("accountCode")

	userAccount, err := h.userAccountService.GetUserAccountByAccountCode(accountCode)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get UserAccount Success", http.StatusOK, "Success", userAccount)

	c.JSON(http.StatusOK, response)
}
func (h *userAccountHandler) GetUserAccountByAccountName(c *gin.Context) {
	accountName := c.Param("accountName")

	userAccount, err := h.userAccountService.GetUserAccountByAccountName(accountName)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get UserAccount Success", http.StatusOK, "Success", userAccount)

	c.JSON(http.StatusOK, response)
}

func (h *userAccountHandler) GetUserAccountByUserId(c *gin.Context) {
	userID := c.Param("createdby")

	userIdInt, err := strconv.Atoi(userID)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAccount, err := h.userAccountService.GetUserAccountByUserId(userIdInt)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get UserAccount Success", http.StatusOK, "Success", userAccount)

	c.JSON(http.StatusOK, response)
}

func (h *userAccountHandler) DeleteUserAccount(c *gin.Context) {
	id := c.Param("id")

	idint, err := strconv.Atoi(id)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Delete UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, ok := c.Get("currentUser")
	if !ok {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	status, err := h.userAccountService.DeleteUserAccount(idint, userLogin.(entity.User).UserName)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Delete UserAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Delete UserAccount Success", http.StatusOK, "Success", status)

	c.JSON(http.StatusOK, response)
}
