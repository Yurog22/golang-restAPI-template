package repository

import (
	"github.com/google/uuid"
	"template/internal/models"
	"time"
)

type MockRepository struct{}

func (m *MockRepository) CreateExample(example *models.Example) (err error) {
	*example = models.Example{
		ID:          uuid.UUID{},
		StringValue: "Example",
		IntValue:    2,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return nil
}

func (m *MockRepository) UpdateExample(example *models.Example) (err error) {
	*example = models.Example{
		ID:          uuid.New(),
		StringValue: "Example",
		IntValue:    2,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return nil
}

func (m *MockRepository) DeleteExample(*models.Example) (err error) {
	return nil
}

func (m *MockRepository) GetExample(example *models.Example) (err error) {
	*example = models.Example{
		ID:          uuid.UUID{},
		StringValue: "Example",
		IntValue:    2,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return nil
}

func (m *MockRepository) ListExamples(examples *[]models.Example) (err error) {
	*examples = []models.Example{
		{
			ID:          uuid.New(),
			StringValue: "Example 1",
			IntValue:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			StringValue: "Example 2",
			IntValue:    2,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	return nil
}
