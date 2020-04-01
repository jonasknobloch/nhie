package statement

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/neverhaveiever-io/api/internal/category"
	"github.com/neverhaveiever-io/api/internal/database"
	"os"
	"reflect"
	"regexp"
	"testing"
	"time"
)

var row = []string{
	"id",
	"statement",
	"category",
	"created_at",
	"updated_at",
	"deleted_at",
}

var expected = &Statement{
	ID:        uuid.New(),
	Statement: "Never have I ever fucked a coconut.",
	Category:  category.Offensive,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	DeletedAt: nil,
}

var mock sqlmock.Sqlmock

func init() {
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New()

	if err != nil {
		panic(err)
	}

	connection, err := gorm.Open("postgres", db)

	if err != nil {
		panic(err)
	}

	database.C = connection

	if _, ok := os.LookupEnv("LOG_MODE"); ok {
		database.C.LogMode(ok)
	}
}

func TestGetByID(t *testing.T) {
	query := `SELECT * FROM "statements" WHERE "statements"."deleted_at" IS NULL AND (("statements"."id" = $1))`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(expected.ID).WillReturnRows(
		sqlmock.NewRows(row).AddRow(
			expected.ID.String(),
			expected.Statement,
			expected.Category.String(),
			expected.CreatedAt,
			expected.UpdatedAt,
			expected.DeletedAt,
		),
	)

	statement, err := GetByID(expected.ID)

	if err != nil {
		t.Fatalf("Unexpected error. %+v", err)
	}

	if !reflect.DeepEqual(*statement, *expected) {
		t.Fatalf("Unexpected struct contents. %+v", statement)
	}
}

func TestGetByIDReturnsErrorIfStatementNotFound(t *testing.T) {
	query := `SELECT * FROM "statements" WHERE "statements"."deleted_at" IS NULL AND (("statements"."id" = $1))`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(sqlmock.NewRows(row))

	_, err := GetByID(uuid.New())

	if !gorm.IsRecordNotFoundError(err) {
		t.Fatalf("Unexpected error. %+v", err)
	}
}

func TestGetRandomByCategory(t *testing.T) {
	query := `SELECT * FROM "statements" WHERE "statements"."deleted_at" IS NULL AND (("statements"."category" = $1)) ORDER BY random(),"statements"."id" ASC LIMIT 1`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(
		sqlmock.NewRows(row).AddRow(
			expected.ID.String(),
			expected.Statement,
			expected.Category.String(),
			expected.CreatedAt,
			expected.UpdatedAt,
			expected.DeletedAt,
		),
	)

	statement, err := GetRandomByCategory(category.Offensive)

	if err != nil {
		t.Fatalf("Unexpected error. %+v", err)
	}

	if !reflect.DeepEqual(*statement, *expected) {
		t.Fatalf("Unexpected struct contents. %+v", statement)
	}
}

func TestGetRandomByCategoryReturnsErrorIfStatementNotFound(t *testing.T) {
	query := `SELECT * FROM "statements" WHERE "statements"."deleted_at" IS NULL AND (("statements"."category" = $1)) ORDER BY random(),"statements"."id" ASC LIMIT 1`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(sqlmock.NewRows(row))

	_, err := GetRandomByCategory(category.Offensive)

	if !gorm.IsRecordNotFoundError(err) {
		t.Fatalf("Unexpected error. %+v", err)
	}
}
