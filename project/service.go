package project

type Service interface {
	FindProjects(userID int) ([]Project, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindProjects(userID int) ([]Project, error) {
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