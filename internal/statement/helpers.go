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

		if err := database.C.Raw(`SELECT id, statement, category FROM game WHERE category = ? OFFSET ? LIMIT 1;`, category, rand.Intn(pool)).Scan(&statement).Error; err != nil {
			return err
		}

		return nil
	})

	return &statement, err
}

func GetNextByPreviousIDAndCategory(ID uuid.UUID, category category.Category) (*Statement, error) {
	var pos int
	var nextID string

	if err := database.C.Transaction(func(tx *gorm.DB) error {
		if err := tx.Raw(`SELECT position FROM game WHERE id = ?;`, ID).Scan(&pos).Error; err != nil {
			return err
		}

		if err := tx.Raw(`SELECT id
						FROM (SELECT * FROM game WHERE position > ? UNION ALL SELECT * FROM game WHERE position < ?) AS game
						WHERE category = ?
						LIMIT 1;`, pos, pos, category).Scan(&nextID).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return GetByID(uuid.MustParse(nextID))
}
