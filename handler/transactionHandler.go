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

type transactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input input.TransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, ok := c.Get("currentUser")
	if !ok {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transaction, err := h.transactionService.CreateTransaction(input, userLogin.(entity.User).UserName)

	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Create Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Create Transaction Success", http.StatusOK, "Success", transaction)

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) EditTransaction(c *gin.Context) {
	var input input.TransactionEditInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, ok := c.Get("currentUser")
	if !ok {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transaction, err := h.transactionService.EditTransaction(input, userLogin.(entity.User).UserName)

	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Edit Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Edit Transaction Success", http.StatusOK, "Success", transaction)

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetAllTransactions(c *gin.Context) {
	transactions, err := h.transactionService.GetAllTransaction()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get All Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get All Transaction Success", http.StatusOK, "Success", transactions)

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetTransactionById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transaction, err := h.transactionService.GetTransactionById(idInt)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get Transaction Success", http.StatusOK, "Success", transaction)

	c.JSON(http.StatusOK, response)
}
func (h *transactionHandler) GetTransactionByUserId(c *gin.Context) {
	userId := c.Param("userId")
	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transaction, err := h.transactionService.GetTransactionByUserId(userIdInt)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get Transaction Success", http.StatusOK, "Success", transaction)

	c.JSON(http.StatusOK, response)
}
func (h *transactionHandler) GetTransactionByBankAccountId(c *gin.Context) {
	bankAccountId := c.Param("bankAccountId")
	bankAccountIdInt, err := strconv.Atoi(bankAccountId)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transaction, err := h.transactionService.GetTransactionByBankAccountId(bankAccountIdInt)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get Transaction Success", http.StatusOK, "Success", transaction)

	c.JSON(http.StatusOK, response)
}
func (h *transactionHandler) GetTransactionByCategoryCode(c *gin.Context) {
	categoryCode := c.Param("categoryCode")

	transaction, err := h.transactionService.GetTransactionByCategoryCode(categoryCode)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get Transaction Success", http.StatusOK, "Success", transaction)

	c.JSON(http.StatusOK, response)
}
func (h *transactionHandler) GetTransactionByAmount(c *gin.Context) {
	amount := c.Param("amount")
	amountFloat, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transaction, err := h.transactionService.GetTransactionByAmount(amountFloat)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get Transaction Success", http.StatusOK, "Success", transaction)

	c.JSON(http.StatusOK, response)
}
func (h *transactionHandler) GetTransactionByNotes(c *gin.Context) {
	notes := c.Param("notes")

	transaction, err := h.transactionService.GetTransactionByNotes(notes)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get Transaction Success", http.StatusOK, "Success", transaction)

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	idint, err := strconv.Atoi(id)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Delete Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, ok := c.Get("currentUser")
	if !ok {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	status, err := h.transactionService.DeleteTransaction(idint, userLogin.(entity.User).UserName)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Delete Transaction Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Delete Transaction Success", http.StatusOK, "Success", status)

	c.JSON(http.StatusOK, response)
}
