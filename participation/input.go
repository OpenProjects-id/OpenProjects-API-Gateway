package participation

import "open_projects/user"

type GetProjectParticipationsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
