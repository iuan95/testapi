package db

import (
	"context"

	"github.com/iuan95/testapi/controllers"
	"github.com/iuan95/testapi/models"
)

type Querier interface{
	CreateUser(ctx context.Context, user controllers.UserParams) error
	DeleteUser(ctx context.Context, id int) error
	GetUserById(ctx context.Context, id int) (models.Item, error)
	GetAllUsers(ctx context.Context) ([]models.Item, error)
}
// var _ Querier = (*db.Queries)(nil)