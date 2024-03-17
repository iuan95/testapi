package db

import (
	"context"
	"testapi/controllers"
	"testapi/models"
)

type Querier interface{
	
	CreateUser(ctx context.Context, user controllers.UserParams) error
	DeleteUser(ctx context.Context, id int) error
	GetUserById(ctx context.Context, id int) (models.Item, error)
	GetAllUsers(ctx context.Context) ([]models.Item, error)
}
var _ Querier = (*Queries)(nil)