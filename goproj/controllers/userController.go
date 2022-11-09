package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	database "hello/db"
	"net/http"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

// HashPassword is used to encrypt the password before it is stored in the DB
func ViewProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		a, _ := c.Cookie("username")
		//x, _ := FindUserPosts(context.Background(), a)
		x, _ := FindOneUserClone(context.Background(), a)
		c.HTML(http.StatusOK, "Profile.html", x)
	}
}
func UnSub() gin.HandlerFunc {
	return func(c *gin.Context) {
		a, _ := c.Cookie("username")
		//x, _ := FindUserPosts(context.Background(), a)
		yesOrNo := c.PostForm("byebye")
		fmt.Println(yesOrNo)
		if yesOrNo == "Unsubscribe" {
			unsubscribe, err := FindOneUserAndUnsubscribe(context.Background(), a)
			unsub, err := UpdateUserInfoUnsub(context.Background(), a)
			if err != nil {
				return
			}
			if err != nil {
				return
			}
			c.Redirect(303, "/profile")
			fmt.Println(unsubscribe)
			fmt.Println(unsub)
		} else {
			subscribe, err := FindOneUserAndSubscribe(context.Background(), a)
			sub, err := UpdateUserInfoSub(context.Background(), a)
			if err != nil {
				return
			}
			if err != nil {
				return
			}
			c.Redirect(303, "/profile")
			fmt.Println(subscribe)
			fmt.Println(sub)
		}

	}
}
func BuyLux() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "buyLux.html", nil)
	}
}
func BuyLuxMoney() gin.HandlerFunc {
	return func(c *gin.Context) {
		x, _ := c.Cookie("username")
		a := c.PostForm("payment")
		if a == "10" {
			lux, err := ToLux(context.Background(), x)
			if err != nil {
				return
			}
			c.Redirect(303, "/profile")
			fmt.Println(lux)
		} else {
			c.JSON(http.StatusBadRequest, "Please enter 10$")
		}
	}
}
