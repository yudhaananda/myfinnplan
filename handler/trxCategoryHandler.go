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

type trxCategoryHandler struct {
	trxCategoryService service.TrxCategoryService
}

func NewTrxCategoryHandler(trxCategoryService service.TrxCategoryService) *trxCategoryHandler {
	return &trxCategoryHandler{trxCategoryService}
}

func (h *trxCategoryHandler) CreateTrxCategory(c *gin.Context) {
	var input input.TrxCategoryInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, ok := c.Get("currentUser")
	if !ok {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	trxCategory, err := h.trxCategoryService.CreateTrxCategory(input, userLogin.(entity.User).UserName)

	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Create TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Create TrxCategory Success", http.StatusOK, "Success", trxCategory)

	c.JSON(http.StatusOK, response)
}

func (h *trxCategoryHandler) EditTrxCategory(c *gin.Context) {
	var input input.TrxCategoryEditInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, ok := c.Get("currentUser")
	if !ok {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	trxCategory, err := h.trxCategoryService.EditTrxCategory(input, userLogin.(entity.User).UserName)

	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Edit TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Edit TrxCategory Success", http.StatusOK, "Success", trxCategory)

	c.JSON(http.StatusOK, response)
}

func (h *trxCategoryHandler) GetAllTrxCategorys(c *gin.Context) {
	trxCategorys, err := h.trxCategoryService.GetAllTrxCategory()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get All TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get All TrxCategory Success", http.StatusOK, "Success", trxCategorys)

	c.JSON(http.StatusOK, response)
}

func (h *trxCategoryHandler) GetTrxCategoryById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	trxCategory, err := h.trxCategoryService.GetTrxCategoryById(idInt)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get TrxCategory Success", http.StatusOK, "Success", trxCategory)

	c.JSON(http.StatusOK, response)
}
func (h *trxCategoryHandler) GetTrxCategoryByCategoryCode(c *gin.Context) {
	categoryCode := c.Param("categoryCode")

	trxCategory, err := h.trxCategoryService.GetTrxCategoryByCategoryCode(categoryCode)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get TrxCategory Success", http.StatusOK, "Success", trxCategory)

	c.JSON(http.StatusOK, response)
}
func (h *trxCategoryHandler) GetTrxCategoryByCategoryName(c *gin.Context) {
	categoryName := c.Param("categoryName")

	trxCategory, err := h.trxCategoryService.GetTrxCategoryByCategoryName(categoryName)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Get TrxCategory Success", http.StatusOK, "Success", trxCategory)

	c.JSON(http.StatusOK, response)
}

func (h *trxCategoryHandler) DeleteTrxCategory(c *gin.Context) {
	id := c.Param("id")

	idint, err := strconv.Atoi(id)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Delete TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, ok := c.Get("currentUser")
	if !ok {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Edit TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	status, err := h.trxCategoryService.DeleteTrxCategory(idint, userLogin.(entity.User).UserName)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Delete TrxCategory Failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Delete TrxCategory Success", http.StatusOK, "Success", status)

	c.JSON(http.StatusOK, response)
}
