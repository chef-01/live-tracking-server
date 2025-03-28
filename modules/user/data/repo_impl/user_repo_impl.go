package repo_impl

import (
	"github.com/chef-01/live-tracking-server/config"
	"github.com/chef-01/live-tracking-server/modules/user/data/models"
	"github.com/chef-01/live-tracking-server/modules/user/domain/entities"
)

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Create(user *entities.User) error {
	dbUser := models.User{Name: user.Name, Email: user.Email}
	result := config.DB.Create(&dbUser)
	user.ID = dbUser.ID
	return result.Error
}

func (r *UserRepositoryImpl) Update(user *entities.User) error {
	return config.DB.Model(&models.User{}).Where("id = ?", user.ID).
		Updates(models.User{Name: user.Name, Email: user.Email}).Error
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}

func (r *UserRepositoryImpl) GetByID(id uint) (*entities.User, error) {
	var dbUser models.User
	result := config.DB.First(&dbUser, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities.User{ID: dbUser.ID, Name: dbUser.Name, Email: dbUser.Email}, nil
}

func (r *UserRepositoryImpl) GetAll() ([]*entities.User, error) {
	var dbUsers []models.User
	result := config.DB.Find(&dbUsers)
	if result.Error != nil {
		return nil, result.Error
	}
	var users []*entities.User
	for _, u := range dbUsers {
		users = append(users, &entities.User{ID: u.ID, Name: u.Name, Email: u.Email})
	}
	return users, nil
}
