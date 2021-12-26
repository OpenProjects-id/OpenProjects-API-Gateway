package project

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetProjects(userID int) ([]Project, error)
	GetProjectByID(input GetProjectDetailInput) (Project, error)
	CreateProject(input CreateProjectInput) (Project, error)
	UpdateProject(inputID GetProjectDetailInput, inputData CreateProjectInput) (Project, error)
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

func (s *service) UpdateProject(inputID GetProjectDetailInput, inputData CreateProjectInput) (Project, error) {
	project, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return project, err
	}

	if project.UserID != inputData.User.ID {
		return project, errors.New("You are not the owner of this project.")
	}

	project.Name = inputData.Name
	project.ShortDescription = inputData.ShortDescription
	project.Description = inputData.Description
	project.Perks = inputData.Perks
	project.TotalBudget = inputData.TotalBudget
	project.TechStacks = inputData.TechStacks

	updatedProject, err := s.repository.Update(project)
	if err != nil {
		return updatedProject, err
	}

	return updatedProject, nil
}
