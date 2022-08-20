package sqlstore

import "github.com/VladKuch/go_practic/internal/app/model"

type SearchRepository struct {
	store *Store
}

func (r *SearchRepository) Create(s *model.Search) (*model.Search, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO search (search_type_id) VALUES($1) RETURNING search_id",
		s.SearchType.ID,
	).Scan(&s.ID); err != nil {
		return nil, err
	}

	return s, nil
}
