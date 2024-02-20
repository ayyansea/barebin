package users

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

type User struct {
	Name string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	u := &User{
		Name: id,
		Email: "random@email.com",
	}
	return c.JSON(http.StatusOK, u)
}

func CreateUser(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}

func UpdateUser(c echo.Context) (err error) {
	id := c.Param("id")
	return c.JSON(http.StatusOK, id)
}

func DeleteUser(c echo.Context) (err error) {
	id := c.Param("id")
	return c.JSON(http.StatusOK, id)
}