package handler

import (
	"fmt"
	"net/http"
	"open_projects/helper"
	"open_projects/project"
	"open_projects/user"
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

	err := c.ShouldBindUri(&input)
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

func (h *projectHandler) CreateProject(c *gin.Context) {
	var input project.CreateProjectInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create project", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newProject, err := h.service.CreateProject(input)
	if err != nil {
		response := helper.APIResponse("Failed to create project", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create project", http.StatusOK, "success", project.FormatProject(newProject))
	c.JSON(http.StatusOK, response)
}

func (h *projectHandler) UpdateProject(c *gin.Context) {
	var inputID project.GetProjectDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		reponse := helper.APIResponse("Failed to update project", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, reponse)
		return
	}

	var inputData project.CreateProjectInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update project", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	inputData.User = currentUser

	updatedProject, err := h.service.UpdateProject(inputID, inputData)
	if err != nil {
		reponse := helper.APIResponse("Failed to update project", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, reponse)
		return
	}

	response := helper.APIResponse("Success to update project", http.StatusOK, "success", project.FormatProject(updatedProject))
	c.JSON(http.StatusOK, response)

}

func (h *projectHandler) UploadImage(c *gin.Context) {
	var input project.CreateProjectImageInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload project image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload project image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("project_images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload project image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveProjectImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload project image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("project image successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
