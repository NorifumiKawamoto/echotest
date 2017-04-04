package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Index home page default
func Index(c echo.Context) error {
	user := User{}
	users := user.GetAll()
	return UsersView(c, users)
}

// Show home page default
func Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(id)
	user := User{
		ID: id,
	}
	err = user.Get()
	if err != nil {
		log.Fatal(err)
	}
	return UserView(c, user)
}

// New is User create Form Page
func New(c echo.Context) error {
	return UserFormView(c)
}

// Create is create action
func Create(c echo.Context) error {
	u := User{}
	email := c.FormValue("email")
	u.Email = email

	if err := u.Validate(); err != nil {
		log.Fatal(err)
		return c.String(http.StatusOK, "NG")
	}

	if err := u.Create(); err != nil {
		log.Fatal(err)
	}
	return c.Redirect(301, "/users")
}

// Delete is delete action
func Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	user := User{
		ID: id,
	}
	err = user.Get()
	if err != nil {
		log.Fatal(err)
	}

	err = user.Delete()
	if err != nil {
		log.Fatal(err)
	}

	return c.NoContent(http.StatusNoContent)
}
