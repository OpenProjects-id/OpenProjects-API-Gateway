package transaction

import "open_projects/user"

type GetProjectTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
