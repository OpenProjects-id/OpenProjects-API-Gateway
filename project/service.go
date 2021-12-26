package project

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetProjects(userID int) ([]Project, error)
	GetProjectByID(input GetProjectDetailInput) (Project, error)
	CreateProject(input CreateProjectInput) (Project, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProjects(userID int) ([]Project, error) {
	if userID != 0 {
		projects, err := s.repository.FindByUserID(userID)
		if err != nil {
			return projects, err
		}

		return projects, nil
	}

	projects, err := s.repository.FindAll()
	if err != nil {
		return projects, err
	}

	return projects, nil
}

func (s *service) GetProjectByID(input GetProjectDetailInput) (Project, error) {
	project, err := s.repository.FindByID(input.ID)

	if err != nil {
		return project, err
	}

	return project, nil
}

func (s *service) CreateProject(input CreateProjectInput) (Project, error) {
	project := Project{}
	project.Name = input.Name
	project.ShortDescription = input.ShortDescription
	project.Description = input.Description
	project.TechStacks = input.TechStacks
	project.TotalBudget = input.TotalBudget
	project.Perks = input.Perks
	project.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	project.Slug = slug.Make(slugCandidate)

	newProject, err := s.repository.Save(project)
	if err != nil {
		return newProject, err
	}

	return newProject, nil
}
