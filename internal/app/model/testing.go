package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		ID:        5,
		Email:     "test1@mail.ru",
		FirstName: "Test",
		LastName:  "User",
		Password:  "123456",
	}
}
