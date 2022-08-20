package teststore_test

import (
	"testing"

	"github.com/VladKuch/go_practic/internal/app/model"
	"github.com/VladKuch/go_practic/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)
	s.User().Create(u)

	u, err := s.User().Find(u.ID)

	assert.NoError(t, err)

	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)
	s.User().Create(u)

	u, err := s.User().FindByEmail("user@example.org")

	assert.NoError(t, err)

	assert.NotNil(t, u)
}
