package handler

import (
	"myfinnplan/helper"
	"myfinnplan/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bankHandler struct {
	bankService service.BankService
}

func NewBankHandler(bankService service.BankService) *bankHandler{
	return &bankHandler{bankService}
}

func (h *bankHandler) GetBankData(c *gin.Context){
	bank, err := h.bankService.GetBankData()

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Get Bank Data Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get Bank Data Success", http.StatusOK, "Success", bank)

	c.JSON(http.StatusOK, response)

}