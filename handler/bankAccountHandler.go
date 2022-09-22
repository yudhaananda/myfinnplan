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

type bankAccountHandler struct {
	bankAccountService service.BankAccountService
}

func NewBankAccountHandler(bankAccountService service.BankAccountService) *bankAccountHandler {
	return &bankAccountHandler{bankAccountService}
}

func (h *bankAccountHandler) CreateBankAccount(c *gin.Context) {
	var input input.BankAccountInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, ok := c.Get("currentUser")
	if !ok {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	bankAccount, err := h.bankAccountService.CreateBankAccount(input, userLogin.(entity.User).UserName)

	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Create BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Create BankAccount Success", http.StatusOK, "Success", bankAccount)

	c.JSON(http.StatusOK, response)
}

func (h *bankAccountHandler) EditBankAccount(c *gin.Context) {
	var input input.BankAccountEditInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, ok := c.Get("currentUser")
	if !ok {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	bankAccount, err := h.bankAccountService.EditBankAccount(input, userLogin.(entity.User).UserName)

	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Edit BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Edit BankAccount Success", http.StatusOK, "Success", bankAccount)

	c.JSON(http.StatusOK, response)
}

func (h *bankAccountHandler) GetAllBankAccounts(c *gin.Context) {
	bankAccounts, err := h.bankAccountService.GetAllBankAccount()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get All BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get All BankAccount Success", http.StatusOK, "Success", bankAccounts)

	c.JSON(http.StatusOK, response)
}

func (h *bankAccountHandler) GetBankAccountById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	bankAccount, err := h.bankAccountService.GetBankAccountById(idInt)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get BankAccount Success", http.StatusOK, "Success", bankAccount)

	c.JSON(http.StatusOK, response)
}
func (h *bankAccountHandler) GetBankAccountByAccountCode(c *gin.Context) {
	accountCode := c.Param("accountCode")

	bankAccount, err := h.bankAccountService.GetBankAccountByAccountCode(accountCode)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get BankAccount Success", http.StatusOK, "Success", bankAccount)

	c.JSON(http.StatusOK, response)
}
func (h *bankAccountHandler) GetBankAccountByAccountIdOwner(c *gin.Context) {
	accountIdOwner := c.Param("id")

	idInt, err := strconv.Atoi(accountIdOwner)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	bankAccount, err := h.bankAccountService.GetBankAccountByAccountIdOwner(idInt)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get BankAccount Success", http.StatusOK, "Success", bankAccount)

	c.JSON(http.StatusOK, response)
}
func (h *bankAccountHandler) GetBankAccountByBankCode(c *gin.Context) {
	bankCode := c.Param("bankCode")

	bankAccount, err := h.bankAccountService.GetBankAccountByBankCode(bankCode)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get BankAccount Success", http.StatusOK, "Success", bankAccount)

	c.JSON(http.StatusOK, response)
}
func (h *bankAccountHandler) GetBankAccountByAmount(c *gin.Context) {
	amount := c.Param("amount")
	amountFloat, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	bankAccount, err := h.bankAccountService.GetBankAccountByAmount(amountFloat)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get BankAccount Success", http.StatusOK, "Success", bankAccount)

	c.JSON(http.StatusOK, response)
}
func (h *bankAccountHandler) GetBankAccountByNotes(c *gin.Context) {
	notes := c.Param("notes")

	bankAccount, err := h.bankAccountService.GetBankAccountByNotes(notes)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get BankAccount Success", http.StatusOK, "Success", bankAccount)

	c.JSON(http.StatusOK, response)
}

func (h *bankAccountHandler) DeleteBankAccount(c *gin.Context) {
	id := c.Param("id")

	idint, err := strconv.Atoi(id)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Delete BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, ok := c.Get("currentUser")
	if !ok {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	status, err := h.bankAccountService.DeleteBankAccount(idint, userLogin.(entity.User).UserName)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Delete BankAccount Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Delete BankAccount Success", http.StatusOK, "Success", status)

	c.JSON(http.StatusOK, response)
}
