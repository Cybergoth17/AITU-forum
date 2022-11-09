package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Observer interface {
	Subscribe(User)
	Unsubscribe(User)
	Notify()
}

type Admin struct {
	Visitor
	Username string
	Password string
}

func NewAdmin(item *Admin, usertype string) *Admin {
	if item == nil {
		return &Admin{}
	}
	fmt.Println("Admin is already created")
	return item
}
func (a *Admin) Subscribe(reader User) {
	//TODO
}
func (a *Admin) Unsubscribe(reader User) {
	//TODO
}
func (a *Admin) LogAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.PostForm("username") == "Admin" && c.PostForm("password") == "admin" {
			c.SetCookie("username", c.PostForm("username"), 3600, "/", "localhost", false, false)
			c.Redirect(303, "/seeUsers")
		} else {
			c.Redirect(303, "/users/login")
		}
	}
}

type AdminInterface interface {
	LogAdmin()
}
