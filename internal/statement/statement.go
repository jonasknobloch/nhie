package statement

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/neverhaveiever-io/api/internal/cache"
	"github.com/neverhaveiever-io/api/internal/category"
	"github.com/neverhaveiever-io/api/internal/database"
	"github.com/neverhaveiever-io/api/internal/translate"
	"golang.org/x/text/language"
	"time"
)

type (
	Statement struct {
		ID           uuid.UUID         `gorm:"primary_key;type:uuid"`
		Statement    string            `gorm:"unique;not null" json:"statement"`
		Category     category.Category `gorm:"type:category;not null" json:"category"`
		Translations translations      `gorm:"-" json:"translations,omitempty"`
		CreatedAt    time.Time         `json:"-"`
		UpdatedAt    time.Time         `json:"-"`
		DeletedAt    *time.Time        `sql:"index" json:"-"`
	}
)

type translations map[string]string

func (*Statement) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", uuid.New())
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

func (s *Statement) FetchTranslations(tags ...language.Tag) error {
	s.Translations = make(translations)

	var t string
	var err error

	// non cache errors
	var errs []error

	for _, tag := range tags {

		t, err = translate.C.Translate(s.ID, s.Statement, tag)

		if err != nil {
			var e *cache.Error

			// skip tag if not a cache error
			if !errors.As(err, &e) {
				errs = append(errs, err)
				continue
			}
		}

		s.Translations[tag.String()] = t
	}

	if l := len(errs); l > 0 {
		return errs[l-1]
	}

	return err
}
