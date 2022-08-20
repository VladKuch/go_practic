package sqlstore_test

import (
	"testing"

	"github.com/VladKuch/go_practic/internal/app/model"
	"github.com/VladKuch/go_practic/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)

	defer teardown("users")

	u := model.TestUser(t)
	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)

	defer teardown("users")
	u := model.TestUser(t)
	s.User().Create(u)

	u, err := s.User().Find(u.ID)

	assert.NoError(t, err)

	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)

	defer teardown("users")

	s.User().Create(model.TestUser(t))

	u, err := s.User().FindByEmail("user@example.org")

	assert.NoError(t, err)

	assert.NotNil(t, u)
}
