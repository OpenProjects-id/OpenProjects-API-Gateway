package participation

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByProjectID(projectID int) ([]Participation, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByProjectID(projectID int) ([]Participation, error) {
	var participations []Participation

	err := r.db.Preload("User").Where("project_id = ?", projectID).Order("id desc").Find(&participations).Error
	if err != nil {
		return participations, err
	}

	return participations, nil
}
