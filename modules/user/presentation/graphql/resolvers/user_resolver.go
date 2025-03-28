package resolvers

import (
	"context"

	"github.com/chef-01/live-tracking-server/modules/user/presentation/controller"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type UserResolver struct {
	Controller *controller.UserController
}

func NewUserResolver(ctrl *controller.UserController) *UserResolver {
	return &UserResolver{Controller: ctrl}
}

func (r *UserResolver) Query_users(ctx context.Context) ([]*User, error) {
	users, err := r.Controller.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var result []*User
	for _, u := range users {
		result = append(result, &User{
			ID:    int(u.ID),
			Name:  u.Name,
			Email: u.Email,
		})
	}
	return result, nil
}

func (r *UserResolver) Query_user(ctx context.Context, id int) (*User, error) {
	user, err := r.Controller.GetByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r *UserResolver) Mutation_createUser(ctx context.Context, name, email string) (*User, error) {
	user, err := r.Controller.Create(ctx, name, email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r *UserResolver) Mutation_updateUser(ctx context.Context, id int, name, email string) (*User, error) {
	user, err := r.Controller.Update(ctx, uint(id), name, email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r *UserResolver) Mutation_deleteUser(ctx context.Context, id int) (bool, error) {
	return r.Controller.Delete(ctx, uint(id))
}
