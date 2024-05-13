package repository

import "golang_backend_study/types"

type UserRepository struct {
	userMap []*types.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{},
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	return nil
}

func (u *UserRepository) Update(beforeUser, updatedUser *types.User) error {
	return nil
}

func (u *UserRepository) Delete(user *types.User) error {
	return nil
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}
