package statement

import (
	"github.com/google/uuid"
	"github.com/nhie-io/api/internal/category"
	"github.com/nhie-io/api/internal/database"
	"gorm.io/gorm"
	"math/rand"
)

func GetByID(ID uuid.UUID) (*Statement, error) {
	var statement Statement

	if err := database.C.Where(&Statement{ID: ID}).Take(&statement).Error; err != nil {
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

		if err := tx.Where(&Statement{Category: category}).Offset(rand.Intn(int(poolSize))).Take(&statement).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return &statement, poolSize, nil
}
