package sqlstore

import (
	"github.com/VladKuch/go_practic/internal/app/store"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db               *sqlx.DB
	searchRepository *SearchRepository
	userRepository   *UserRepository
}

func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Search() *SearchRepository {
	if s.searchRepository != nil {
		return s.searchRepository
	}

	s.searchRepository = &SearchRepository{
		store: s,
	}

	return s.searchRepository
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
