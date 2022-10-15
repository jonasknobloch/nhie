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

	if err := database.C.Raw(`SELECT id, statement, category FROM game WHERE id = ?;`, ID).Scan(&statement).Error; err != nil {
		return nil, err
	}

	return &statement, nil
}

func GetRandomByCategory(category category.Category) (*Statement, error) {
	var pool int
	var statement Statement

	err := database.C.Transaction(func(tx *gorm.DB) error {
		if err := database.C.Raw(`SELECT COUNT(*) FROM game WHERE category = ?;`, category).Scan(&pool).Error; err != nil {
			return err
		}

		if err := database.C.Raw(`SELECT id, statement, category FROM game OFFSET ? LIMIT 1;`, rand.Intn(pool+1)-1).Scan(&statement).Error; err != nil {
			return err
		}

		return nil
	})

	return &statement, err
}
