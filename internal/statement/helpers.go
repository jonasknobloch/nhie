package statement

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/neverhaveiever-io/api/internal/category"
	"github.com/neverhaveiever-io/api/internal/database"
)

func GetByID(ID uuid.UUID) (*Statement, error) {
	var statement Statement

	if err := database.C.Where(&Statement{ID: ID}).Find(&statement).Error; err != nil {
		return nil, err
	}

	return &statement, nil
}

func GetRandomByCategory(category category.Category) (*Statement, int64, error) {
	var statement Statement
	var poolSize int64

	err := database.C.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Statement{}).Where(&Statement{Category: category}).Count(&poolSize).Error; err != nil {
			return err
		}

		if err := tx.Where(&Statement{Category: category}).Order(gorm.Expr("random()")).First(&statement).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return &statement, poolSize, nil
}
