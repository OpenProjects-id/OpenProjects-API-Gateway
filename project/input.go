package project

import "open_projects/user"

type GetProjectDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateProjectInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	TechStacks       string `json:"tech_stacks" binding:"required"`
	TotalBudget      int    `json:"total_budget" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}
