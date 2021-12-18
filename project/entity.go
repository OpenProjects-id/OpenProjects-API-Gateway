package project

import "time"

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
	ProjectImages []ProjectImage
}

type ProjectImage struct {
	ID        int
	ProjectID int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}
