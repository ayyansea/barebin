package main

import (
	"github.com/labstack/echo/v4"

	"github.com/ayyansea/barebin/users"
)

func main() {
	e := echo.New()

	e.GET("/users/:id", users.GetUser).Name = "get-user"
	e.POST("/users", users.CreateUser).Name = "create-user"
	e.PATCH("/users/:id", users.UpdateUser).Name = "update-user"
	e.DELETE("/users/:id", users.DeleteUser).Name = "delete-user"

	e.Logger.Fatal(e.Start(":1323"))
}