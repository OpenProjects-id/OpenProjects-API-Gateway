package project

import (
	"open_projects/user"
	"time"
)

type Project struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	TechStacks       string
	ParticipantCount int
	TotalBudget      int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ProjectImages    []ProjectImage
	User             user.User
}

type ProjectImage struct {
	ID        int
	ProjectID int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}
