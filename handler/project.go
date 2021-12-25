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

	response := helper.APIResponse("List of Projects", http.StatusOK, "success", project.FormatProjects(projects))
	c.JSON(http.StatusOK, response)
}

func (h *projectHandler) GetProject(c *gin.Context) {
	var input project.GetProjectDetailInput

	err := c.ShouldBindUri((&input))
	if err != nil {
		reponse := helper.APIResponse("Failed to get project's detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, reponse)
		return
	}

	projectDetail, err := h.service.GetProjectByID(input)
	if err != nil {
		reponse := helper.APIResponse("Failed to get project's detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, reponse)
		return
	}

	response := helper.APIResponse("Project Detail", http.StatusOK, "success", project.FormatProjectDetail(projectDetail))
	c.JSON(http.StatusOK, response)
}
