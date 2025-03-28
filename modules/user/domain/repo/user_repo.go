package repo

import "github.com/chef-01/live-tracking-server/modules/user/domain/entities"



type UserRepository interface {
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(id uint) error
	GetByID(id uint) (*entities.User, error)
	GetAll() ([]*entities.User, error)
}
