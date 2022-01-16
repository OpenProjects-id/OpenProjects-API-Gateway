package handler

import (
	"net/http"
	"open_projects/helper"
	"open_projects/participation"
	"open_projects/user"

	"github.com/gin-gonic/gin"
)

type participationHandler struct {
	service participation.Service
}

func NewParticipationHandler(service participation.Service) *participationHandler {
	return &participationHandler{service}
}

func (h *participationHandler) GetProjectParticipations(c *gin.Context) {
	var input participation.GetProjectParticipationsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		reponse := helper.APIResponse("Failed to get project's participations", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, reponse)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	participations, err := h.service.GetParticipationsByProjectID(input)
	if err != nil {
		reponse := helper.APIResponse("Failed to get project's participations", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, reponse)
		return
	}

	response := helper.APIResponse("Project's participation", http.StatusOK, "success", participation.FormatProjectParticipations(participations))
	c.JSON(http.StatusOK, response)
}
