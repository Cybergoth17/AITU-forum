package main

import (
	"github.com/gin-gonic/gin"
	"hello/controllers"
	"hello/middleware"
	"hello/models"
	"hello/routes"
	"log"
	"os"
)

var admin *models.Admin

func init() {
	admin = &models.Admin{Username: "admin"}
}

func main() {
	admin = models.NewAdmin(admin, "")
	admin = models.NewAdmin(admin, "")
	admin = models.NewAdmin(admin, "")
	admin = models.NewAdmin(admin, "")
	admin = models.NewAdmin(admin, "")
	admin = models.NewAdmin(admin, "")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/images", "./templates/images/")
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.Run(":" + port)

	a := false
	subs := controllers.NewSubs(10)
	err := subs.RequestSub(a)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = subs.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = subs.SuccessSubs()
	err = subs.InsertMoney(10)
	err = subs.SuccessSubs()

}
