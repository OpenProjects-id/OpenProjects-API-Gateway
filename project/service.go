package project

type Service interface {
	GetProjects(userID int) ([]Project, error)
	GetProjectByID(input GetProjectDetailInput) (Project, error)
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
