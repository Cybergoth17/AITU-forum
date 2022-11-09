package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Bone struct {
	Text1 string
	Text2 string
	Text3 string
}

func MainPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		n := &Noauth{}
		a := &Auth{Name: "Anel", Age: 18, Password: "1234567", Email: "211395@astanait.edu.kz", Wrapper: n}
		admin := &Admin{IsAdmin: true, Wrapper: a}
		bone := Bone{Text1: n.DoYourOwnJob(),
			Text2: a.DoYourOwnJob(),
			Text3: admin.DoYourOwnJob(),
		}
		fmt.Println(bone)
		c.HTML(http.StatusOK, "index.html", bone)
	}
}
