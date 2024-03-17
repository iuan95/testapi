package main

import (
	"context"
	"testapi/db"

	"github.com/gofiber/fiber/v2"
)


func main() {
    app := fiber.New()
	err:=db.DbConnection()
	if err!= nil {
		panic("can not connection to DB")
	}

	app.Get("/users", func(c *fiber.Ctx) error {
       items,err:= db.Querier.GetAllUsers(context.Background())
	   if(err!=nil){
		c.Status(401).JSON(&fiber.Map{
			"message": "can not find users",
		})
		return err
	   }
	   c.Status(200).JSON(&fiber.Map{
		"data": items,
	})
	   return nil
    })


    app.Listen(":3000")
}