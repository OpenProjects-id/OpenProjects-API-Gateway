package participation

import (
	"errors"
	"open_projects/project"
)

type service struct {
	repository        Repository
	projectRepository project.Repository
}

type Service interface {
	GetParticipationsByProjectID(input GetProjectParticipationsInput) ([]Participation, error)
	GetParticipationsByUserID(userID int) ([]Participation, error)
}

func NewService(repository Repository, projectRepository project.Repository) *service {
	return &service{repository, projectRepository}
}

func (s *service) GetParticipationsByProjectID(input GetProjectParticipationsInput) ([]Participation, error) {

	project, err := s.projectRepository.FindByID(input.ID)
	if err != nil {
		return []Participation{}, err
	}

	if project.UserID != input.User.ID {
		return []Participation{}, errors.New("not an owner of the project")
	}

	participation, err := s.repository.GetByProjectID(input.ID)
	if err != nil {
		return participation, err
	}

	return participation, nil
}

func (s *service) GetParticipationsByUserID(userID int) ([]Participation, error) {
	participations, err := s.repository.GetByUserID(userID)
	if err != nil {
		return participations, err
	}
	return participations, nil
}
