package statement

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/neverhaveiever-io/api/internal/category"
	"github.com/neverhaveiever-io/api/internal/database"
	"time"
)

type (
	Statement struct {
		ID        uuid.UUID         `gorm:"primary_key;type:uuid"`
		Statement string            `gorm:"unique;not null" json:"statement"`
		Category  category.Category `gorm:"not null" json:"category" sql:"type:category"`
		CreatedAt time.Time         `json:"-"`
		UpdatedAt time.Time         `json:"-"`
		DeletedAt *time.Time        `sql:"index" json:"-"`
	}
)

func (*Statement) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", uuid.New())
}

func (s *Statement) Save() error {
	return database.C.Save(&s).Error
}

func (s *Statement) Delete() error {
	return database.C.Delete(&s).Error
}
