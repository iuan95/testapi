package controllers

import (
	"context"
	"fmt"
	"testapi/models"

	"github.com/iuan95/testapi/db"
)

type UserParams struct{
	Name string `json:"name"`
	Description string `json:"description"`
}

func (q *db.Queries) CreateUser(ctx context.Context, user UserParams) ( error){
	fmt.Println("Asd")
	err:=q.DB.Exec(ctx,"INSRT INTO items (name, description) VALUES($1, $2)", user.Name, user.Description)
	return  err
}


func (q *db.Queries) DeleteUser(ctx context.Context, id int) error{
	err:= q.DB.Exec(ctx,"DELETE FROM items WHERE id=$1", id)
	return err

}

func (q *db.Queries) GetUserById(ctx context.Context, id int) (models.Item, error){
	var item models.Item
	err:= q.DB.QueryRow(ctx,"SELECT * FROM items WHERE id=$1", id).Scan(&item.Id, &item.Name, &item.Description)
	return err
}

func (q *db.Queries) GetAllUsers(ctx context.Context) ([]models.Item, error){
	var items []models.Item
	rows, _:= q.DB.Query(ctx,"SELECT * FROM items")
	for rows.Next(){
		var item models.Item
		rows.Scan(&item.Id, &item.Name, &item.Description)
		items =append(items, item)
	}
	return items, nil
}