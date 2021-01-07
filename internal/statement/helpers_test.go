package statement

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/nhie-io/api/internal/category"
	"github.com/nhie-io/api/internal/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	DeletedAt: gorm.DeletedAt{},
}

var mock sqlmock.Sqlmock

func init() {
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New()

	if err != nil {
		panic(err)
	}

	connection, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.C = connection

	if _, ok := os.LookupEnv("LOG_MODE"); ok {
		database.C.Logger.LogMode(logger.Info)
	}
}

func mockGetRandomByCategory(rows []*sqlmock.Rows) {
	count := `SELECT count(*) FROM "statements"  WHERE "statements"."deleted_at" IS NULL AND (("statements"."category" = $1))`
	query := `SELECT * FROM "statements" WHERE "statements"."deleted_at" IS NULL AND (("statements"."category" = $1)) ORDER BY random(),"statements"."id" ASC LIMIT 1`

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(count)).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(rows)))
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows...)
	mock.ExpectCommit()
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

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("Unexpected error. %+v", err)
	}
}

func TestGetRandomByCategory(t *testing.T) {
	rows := []*sqlmock.Rows{
		sqlmock.NewRows(row).AddRow(
			expected.ID.String(),
			expected.Statement,
			expected.Category.String(),
			expected.CreatedAt,
			expected.UpdatedAt,
			expected.DeletedAt,
		),
	}

	mockGetRandomByCategory(rows)

	statement, p, err := GetRandomByCategory(category.Offensive)

	if err != nil {
		t.Fatalf("Unexpected error. %+v", err)
	}

	if !reflect.DeepEqual(*statement, *expected) {
		t.Fatalf("Unexpected struct contents. %+v", statement)
	}

	if p == 0 {
		t.Fatalf("Pool should be empty.")
	}
}

func TestGetRandomByCategoryReturnsErrorIfStatementNotFound(t *testing.T) {
	mockGetRandomByCategory([]*sqlmock.Rows{sqlmock.NewRows(row)})

	_, p, err := GetRandomByCategory(category.Offensive)

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("Unexpected error. %+v", err)
	}

	if p != 0 {
		t.Fatalf("Expected empty pool, got %d instead.", p)
	}
}
