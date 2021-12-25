package project

type GetProjectDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
