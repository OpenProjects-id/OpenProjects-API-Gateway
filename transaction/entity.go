package transaction

import (
	"open_projects/project"
	"open_projects/user"
	"time"
)

type Transaction struct {
	ID        int
	ProjectID int
	UserID    int
	Amount    int
	Status    string
	Code      string
	User      user.User
	Project   project.Project
	CreatedAt time.Time
	UpdatedAt time.Time
}
