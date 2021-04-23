package store_test

import (
	"github.com/stretchr/testify/assert"
	"src/myFirstGoProject/chapter003/http-rest-api/internal/app/model"
	"src/myFirstGoProject/chapter003/http-rest-api/internal/app/store"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	u, err := s.User().Create(&model.User{
		Email:             "create_test@email.ru",
		EncryptedPassword: "create_testEncPWD",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	email := "findById_test@email.ru"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
	s.User().Create(&model.User{
		Email:             email,
		EncryptedPassword: "findById_testEncPWD",
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
