package repository

import (
	"template/internal/db"
	"template/internal/models"
)

type Repository interface {
	CreateExample(example *models.Example) (err error)
	ListExamples(examples *[]models.Example) (err error)
	UpdateExample(example *models.Example) (err error)
	DeleteExample(example *models.Example) (err error)
	GetExample(example *models.Example) (err error)
}

type DefaultRepository struct{}

func (d *DefaultRepository) CreateExample(example *models.Example) (err error) {
	err = db.DB.Create(&example).Error
	return
}

func (d *DefaultRepository) UpdateExample(example *models.Example) (err error) {
	err = db.DB.Updates(&example).Error
	return
}

func (d *DefaultRepository) DeleteExample(example *models.Example) (err error) {
	err = db.DB.Delete(&example).Error
	return
}

func (d *DefaultRepository) GetExample(example *models.Example) (err error) {
	err = db.DB.First(&example).Error
	return
}

func (d *DefaultRepository) ListExamples(examples *[]models.Example) (err error) {
	err = db.DB.Find(&examples).Error
	return
}
