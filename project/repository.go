package project

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Project, error)
	FindByUserID(userID int) ([]Project, error)
	FindByID(ID int) (Project, error)
	Save(project Project) (Project, error)
	Update(project Project) (Project, error)
	CreateImage(projectImage ProjectImage) (ProjectImage, error)
	MarkAllImagesAsNonoPrimary(projectID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Project, error) {
	var projects []Project

	err := r.db.Preload("ProjectImages", "project_images.is_primary = 1").Find(&projects).Error
	if err != nil {
		return projects, err
	}

	return projects, nil
}

func (r *repository) FindByUserID(userID int) ([]Project, error) {
	var projects []Project

	err := r.db.Where("user_id", userID).Preload("ProjectImages", "project_images.is_primary = 1").Find(&projects).Error
	if err != nil {
		return projects, err
	}

	return projects, nil

}

func (r *repository) FindByID(ID int) (Project, error) {
	var project Project

	err := r.db.Preload("User").Preload("ProjectImages").Where("id = ?", ID).Find(&project).Error

	if err != nil {
		return project, err
	}

	return project, nil

}

func (r *repository) Save(project Project) (Project, error) {
	err := r.db.Create(&project).Error
	if err != nil {
		return project, err
	}

	return project, nil
}

func (r *repository) Update(project Project) (Project, error) {
	err := r.db.Save(&project).Error

	if err != nil {
		return project, err
	}

	return project, nil
}

func (r *repository) CreateImage(projectImage ProjectImage) (ProjectImage, error) {
	err := r.db.Create(&projectImage).Error
	if err != nil {
		return projectImage, err
	}

	return projectImage, nil
}

func (r *repository) MarkAllImagesAsNonoPrimary(projectID int) (bool, error) {
	err := r.db.Model(&ProjectImage{}).Where("project_id = ?", projectID).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
