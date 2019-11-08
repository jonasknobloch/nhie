package statement

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/neverhaveiever-io/api/internal/category"
	"github.com/neverhaveiever-io/api/internal/database"
)

func GetByID(ID uuid.UUID) (Statement, error) {
	var statement Statement

	if err := database.C.Where(&Statement{ID: ID}).Find(&statement).Error; err != nil {
		return statement, err
	}

	return statement, nil
}

func GetRandomByCategory(category category.Category) (Statement, error) {
	var statement Statement

	if err := database.C.Where(
		&Statement{
			Category: category,
		}).Order(gorm.Expr("random()")).First(&statement).Error; err != nil {

		return statement, err
	}

	return statement, nil
}
