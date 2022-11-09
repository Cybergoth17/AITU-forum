package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"net/smtp"
	"time"
)

func SeeUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		const AdminA = "admin"
		a, _ := c.Cookie("username")
		switch a {
		default:
			users, _ := FindAllUsers(context.Background())
			c.HTML(http.StatusOK, "seeUsers.html", users)
		case AdminA:
			users, _ := FindAllUsers(context.Background())
			c.HTML(http.StatusOK, "seeUsers.html", users)
		case "":
			c.JSON(http.StatusOK, "Please signup or login first")
		}
	}
}
func Notice() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.HTML(http.StatusOK, "Notification.html", nil)
	}
}
func CreateNotification() gin.HandlerFunc {
	return func(c *gin.Context) {
		a, _ := FindAllSubs(context.Background())
		var k int
		for _, _ = range a {
			k++
		}
		x := make([]string, k)
		for q, w := range a {
			x[q] = *w.Email
		}
		auth := smtp.PlainAuth(
			"",
			"ddev05702@gmail.com",
			"pjudcojtdhlfpshs",
			"smtp.gmail.com",
		)
		fmt.Println(x)
		msg := c.PostForm("message")
		fmt.Println(msg)
		err := smtp.SendMail("smtp.gmail.com:587",
			auth,
			"akmagambetovaanel0@gmail.com",
			x,
			[]byte(msg),
		)
		if err != nil {
			fmt.Println(err)
		}
		c.Redirect(303, "/notice")
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		a, _ := c.Params.Get("username")
		resultInsertionNumber, insertErr := userCollection.DeleteOne(ctx, bson.M{"username": a})
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		fmt.Println(resultInsertionNumber)
		c.Redirect(303, "/seeUsers")
	}
}
func AdminLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "AdminReg.html", nil)
	}
}
