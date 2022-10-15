package statement

import (
	"github.com/google/uuid"
	"github.com/nhie-io/api/internal/category"
	"github.com/nhie-io/api/internal/database"
	"golang.org/x/text/language"
)

type Statement struct {
	ID        uuid.UUID         `json:"ID"`
	Statement string            `json:"statement"`
	Category  category.Category `json:"category"`
}

func (s *Statement) Translate(target language.Tag) error {
	var translation string

	if err := database.C.Raw(`SELECT translation FROM translations WHERE statement_id = ? AND language = ?;`, s.ID.String(), target.String()).Scan(&translation).Error; err != nil {
		return err
	}

	s.Statement = translation
	return nil
}
