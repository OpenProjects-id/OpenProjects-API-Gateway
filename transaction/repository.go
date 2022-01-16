package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByProjectID(projectID int) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByProjectID(projectID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Preload("User").Where("project_id = ?", projectID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
