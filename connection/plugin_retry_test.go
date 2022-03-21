package connection

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"testing"
)

type MainTestSuite struct {
	suite.Suite
	DB    *gorm.DB
	mock  sqlmock.Sqlmock
	RawDB *sql.DB
}

type DbRetryWithoutTransactionTestSuite struct {
	MainTestSuite
}

type DbRetryWithTransactionTestSuite struct {
	MainTestSuite
}

const (
	defaultRetry = 3
)

func TestDbWithRetryTestSuite(s *testing.T) {
	_ = os.Setenv("GORM_RETRY_ATTEMPT", "3")
	_ = os.Setenv("GORM_RETRY_DELAY", "1s")
	suite.Run(s, new(DbRetryWithoutTransactionTestSuite))
	//suite.Run(s, new(DbRetryWithTransactionTestSuite))
}

func (s *MainTestSuite) SetupSuite() {
	var (
		rawDB *sql.DB
		err   error
	)
	rawDB, s.mock, err = sqlmock.New()
	if err != nil {
		s.T().Errorf("Failed to open mock sql db, got error: %v", err)
	}
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 rawDB,
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	err = db.Use(RegisterRetry())
	if err != nil {
		s.T().Errorf("Cannot register plugin: %v", err)
	}
	s.DB = db
	s.RawDB, err = db.DB()
	if err != nil {
		s.T().Error(err)
	}
	err = s.RawDB.Ping()
	if err != nil {
		s.T().Error(err)
	}
}

func (s *DbRetryWithoutTransactionTestSuite) TestDbRetryWithoutTransaction() {
	// Expect to retry 3 times: 2 failed and last 1 success
	for i := 0; i <= defaultRetry; i++ {
		if i != defaultRetry {
			s.mock.ExpectExec("UPDATE products").
				WillReturnError(fmt.Errorf("connection reset by peer [%d]", i))
		} else {
			s.mock.ExpectExec("UPDATE products").
				WillReturnResult(sqlmock.NewResult(1, 1))
		}
	}
	if err := s.DB.Exec("UPDATE products WHERE views = views + 1").Error; err != nil {
		s.T().Errorf("Expect no error return but got error: %s", err)
	}
	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (s *DbRetryWithTransactionTestSuite) TestDbRetryWithTransaction() {
	// Expect to retry 3 times: 2 failed and last 1 success
	for i := 0; i <= defaultRetry; i++ {
		if i != defaultRetry {
			s.mock.ExpectExec("UPDATE products").
				WillReturnError(fmt.Errorf("connection reset by peer [%d]", i))
		} else {
			s.mock.ExpectExec("UPDATE products").
				WillReturnResult(sqlmock.NewResult(1, 1))
		}
	}
	s.mock.ExpectExec("UPDATE product_viewers").WillReturnResult(sqlmock.NewResult(1, 1))
	if err := recordStats(s.DB); err != nil {
		s.T().Errorf("error was not expected while updating stats: %s", err)
	}
	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func recordStats(db *gorm.DB) (err error) {
	tx := db.Begin()
	err = tx.Error
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit().Error
		default:
			tx.Rollback()
		}
	}()

	if err = tx.Exec("UPDATE products SET views = views + 1").Error; err != nil {
		return
	}
	if err = tx.Exec("UPDATE product_viewers SET views = views + 1").Error; err != nil {
		return
	}
	return
}
