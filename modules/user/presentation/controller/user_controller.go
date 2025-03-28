package controller

import (
	"context"

	"github.com/chef-01/live-tracking-server/modules/user/domain/entities"
	"github.com/chef-01/live-tracking-server/modules/user/domain/usecase"
)

type UserController struct {
	Usecase *usecase.UserUsecase
}

func NewUserController(uc *usecase.UserUsecase) *UserController {
	return &UserController{Usecase: uc}
}

func (c *UserController) Create(ctx context.Context, name, email string) (*entities.User, error) {
	user := &entities.User{Name: name, Email: email}
	err := c.Usecase.CreateUser(user)
	return user, err
}

func (c *UserController) Update(ctx context.Context, id uint, name, email string) (*entities.User, error) {
	user := &entities.User{ID: id, Name: name, Email: email}
	err := c.Usecase.UpdateUser(user)
	return user, err
}

func (c *UserController) Delete(ctx context.Context, id uint) (bool, error) {
	err := c.Usecase.DeleteUser(id)
	return err == nil, err
}

func (c *UserController) GetByID(ctx context.Context, id uint) (*entities.User, error) {
	return c.Usecase.GetUser(id)
}

func (c *UserController) GetAll(ctx context.Context) ([]*entities.User, error) {
	return c.Usecase.GetAllUsers()
}
