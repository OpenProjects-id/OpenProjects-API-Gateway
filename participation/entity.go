package participation

import (
	"open_projects/project"
	"open_projects/user"
	"time"
)

type Participation struct {
	ID        int
	ProjectID int
	UserID    int
	Status    string
	Code      string
	User      user.User
	Project   project.Project
	CreatedAt time.Time
	UpdatedAt time.Time
}
