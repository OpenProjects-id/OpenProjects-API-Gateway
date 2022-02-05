package handler

import (
	"net/http"
	"open_projects/helper"
	"open_projects/transaction"
	"open_projects/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetProjectTransactions(c *gin.Context) {
	var input transaction.GetProjectTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		reponse := helper.APIResponse("Failed to get project's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, reponse)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	transactions, err := h.service.GetTransactionsByProjectID(input)
	if err != nil {
		reponse := helper.APIResponse("Failed to get project's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, reponse)
		return
	}

	response := helper.APIResponse("Project's transactions", http.StatusOK, "success", transaction.FormatProjectTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userID)
	if err != nil {
		reponse := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, reponse)
		return
	}
	response := helper.APIResponse("User's transactions", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
