package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	config     *Config
	db         *sql.DB
	userRep    *UserRep
	articleRep *ArticleRep
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	log.Println("DB connection created successefully")
	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) User() *UserRep {
	if s.userRep != nil {
		return s.userRep
	}
	s.userRep = &UserRep{
		storage: s,
	}
	return nil
}

func (s *Storage) Article() *ArticleRep {
	if s.articleRep != nil {
		return s.articleRep
	}
	s.articleRep = &ArticleRep{
		storage: s,
	}

	return nil
}
