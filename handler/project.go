package handler

import (
	"net/http"
	"open_projects/helper"
	"open_projects/project"
	"strconv"

	"github.com/gin-gonic/gin"
)

type projectHandler struct {
	service project.Service
}

func NewProjectHandler(service project.Service) *projectHandler {
	return &projectHandler{service}
}

func (h *projectHandler) GetProjects(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	projects, err := h.service.GetProjects(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get Projects", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of Projects", http.StatusOK, "success", projects)
	c.JSON(http.StatusOK, response)
}