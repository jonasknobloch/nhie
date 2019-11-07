package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/neverhaveiever-io/api/database"
	"regexp"
	"time"
)

type (
	Statement struct {
		ID        uuid.UUID  `gorm:"primary_key;type:uuid"`
		Statement string     `gorm:"unique;not null" json:"statement"`
		Category  Category   `gorm:"not null" json:"category" sql:"type:category"`
		CreatedAt time.Time  `json:"-"`
		UpdatedAt time.Time  `json:"-"`
		DeletedAt *time.Time `sql:"index" json:"-"`
	}
)

func (*Statement) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", uuid.New())
}

func (s Statement) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(
			&s.Statement,
			validation.Required,
			validation.Match(regexp.MustCompile(`Never\shave\sI\sever\s.+\.$`)),
		),
		validation.Field(&s.Category),
	)
}

func (s *Statement) Save() error {
	return database.Connection.Save(&s).Error
}

func (s *Statement) Delete() error {
	return database.Connection.Delete(&s).Error
}

func GetStatementById(ID uuid.UUID) (Statement, error) {
	var statement Statement

	if err := database.Connection.Where(&Statement{ID: ID}).Find(&statement).Error; err != nil {
		return statement, err
	}

	return statement, nil
}

func GetRandomStatementByCategory(category Category) (Statement, error) {
	var statement Statement

	if err := database.Connection.Where(
		&Statement{
			Category: category,
		}).Order(gorm.Expr("random()")).First(&statement).Error; err != nil {

		return statement, err
	}

	return statement, nil
}
