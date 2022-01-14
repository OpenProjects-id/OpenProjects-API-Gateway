package participation

import "time"

type Participation struct {
	ID        int
	ProjectID int
	UserID    int
	Status    string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
