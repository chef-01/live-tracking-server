package usecase

import (
	"github.com/chef-01/live-tracking-server/modules/user/domain/entities"
	"github.com/chef-01/live-tracking-server/modules/user/domain/repo"
)

type UserUsecase struct {
	Repo repo.UserRepository
}

func NewUserUsecase(r repo.UserRepository) *UserUsecase {
	return &UserUsecase{Repo: r}
}

func (u *UserUsecase) CreateUser(user *entities.User) error {
	return u.Repo.Create(user)
}

func (u *UserUsecase) UpdateUser(user *entities.User) error {
	return u.Repo.Update(user)
}

func (u *UserUsecase) DeleteUser(id uint) error {
	return u.Repo.Delete(id)
}

func (u *UserUsecase) GetUser(id uint) (*entities.User, error) {
	return u.Repo.GetByID(id)
}

func (u *UserUsecase) GetAllUsers() ([]*entities.User, error) {
	return u.Repo.GetAll()
}
