package participation

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByProjectID(projectID int) ([]Participation, error)
	GetByUserID(userID int) ([]Participation, error)
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

func (r *repository) GetByUserID(userID int) ([]Participation, error) {
	var participations []Participation

	err := r.db.Preload("Project.ProjectImages", "project_images.is_primary = 1").Where("user_id = ?", userID).Order("id desc").Find(&participations).Error
	if err != nil {
		return participations, err
	}

	return participations, nil
}
