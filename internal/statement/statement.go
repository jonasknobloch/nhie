package statement

import (
	"github.com/google/uuid"
	"github.com/nhie-io/api/internal/category"
	"github.com/nhie-io/api/internal/database"
	"github.com/nhie-io/api/internal/translate"
	"golang.org/x/text/language"
	"gorm.io/gorm"
	"time"
)

type (
	Statement struct {
		ID        uuid.UUID         `gorm:"primary_key;type:uuid"`
		Statement string            `gorm:"unique;not null" json:"statement"`
		Category  category.Category `gorm:"type:category;not null" json:"category"`
		CreatedAt time.Time         `json:"-"`
		UpdatedAt time.Time         `json:"-"`
		DeletedAt gorm.DeletedAt    `gorm:"index" json:"-"`
	}
)

func (s *Statement) BeforeCreate(_ *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}

func (s *Statement) Save() error {
	if (s.ID != uuid.UUID{}) {
		translate.ClearCache(s.ID)
	}

	return database.C.Save(&s).Error
}

func (s *Statement) Delete() error {
	translate.ClearCache(s.ID)
	return database.C.Delete(&s).Error
}

func (s *Statement) Translate(tag language.Tag) error {
	t, err := translate.C.Translate(s.ID, s.Statement, tag)
	s.Statement = t
	return err
}
