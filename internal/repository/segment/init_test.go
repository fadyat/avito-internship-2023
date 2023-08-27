package segment

import (
	"avito-internship-2023/internal/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"log"
	"os"
	"testing"
	"time"
)

type SegmentRepoSuite struct {
	suite.Suite
	r  repository.ISegment
	db *sqlx.DB
}

func newTestDatabase() *sqlx.DB {
	dsn, driver := os.Getenv("TEST_DATABASE_DSN"), os.Getenv("TEST_DATABASE_DRIVER")
	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	return db
}

func (s *SegmentRepoSuite) SetupSuite() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	s.db = newTestDatabase()
	s.r = NewRepository(repository.NewTransactor(s.db))
}

func (s *SegmentRepoSuite) clean() {
	_, err := s.db.Exec("TRUNCATE TABLE segment CASCADE")
	if err != nil {
		s.Require().NoError(err)
	}
}

func (s *SegmentRepoSuite) TearDownTest() {
	s.clean()
}

func (s *SegmentRepoSuite) TearDownSuite() {
	err := s.db.Close()
	if err != nil {
		log.Fatal("failed to close database connection: ", err)
	}
}

func TestSegmentRepository(t *testing.T) {
	suite.Run(t, new(SegmentRepoSuite))
}

func equalDates(a, b *time.Time) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil && b != nil || a != nil && b == nil {
		return false
	}

	ay, am, ad := a.Date()
	by, bm, bd := b.Date()

	ah, amin, _ := a.Clock()
	bh, bmin, _ := b.Clock()

	return ay == by && am == bm && ad == bd && ah == bh && amin == bmin
}
