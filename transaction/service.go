package transaction

import (
	"errors"
	"open_projects/project"
)

type service struct {
	repository        Repository
	projectRepository project.Repository
}

type Service interface {
	GetTransactionsByProjectID(input GetProjectTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
}

func NewService(repository Repository, projectRepository project.Repository) *service {
	return &service{repository, projectRepository}
}

func (s *service) GetTransactionsByProjectID(input GetProjectTransactionsInput) ([]Transaction, error) {

	project, err := s.projectRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if project.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the project")
	}

	transactions, err := s.repository.GetByProjectID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
