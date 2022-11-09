package routes

import (
	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"
	"hello/controllers"
	_ "hello/db"
	"hello/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	visit := models.Visitor{}
	incomingRoutes.GET("/", controllers.MainPage())
	incomingRoutes.GET("/forum", controllers.SeeAllPosts())

	incomingRoutes.GET("/posts/:title", controllers.SeeSinglePost())
	incomingRoutes.POST("/posts/:title", controllers.LeaveAnswer())

	incomingRoutes.GET("/buyLux", controllers.BuyLux())
	incomingRoutes.POST("/buyLux", controllers.BuyLuxMoney())

	incomingRoutes.GET("/post", controllers.Post())
	incomingRoutes.POST("/post", controllers.CreateUserPost())

	incomingRoutes.GET("/notice", controllers.Notice())
	incomingRoutes.POST("/notice", controllers.CreateNotification())

	incomingRoutes.GET("/seeUsers", controllers.SeeUsers())
	incomingRoutes.GET("/seeUsers/:username", controllers.DeleteUser())

	adm := models.Admin{}
	incomingRoutes.GET("/adminLogin", controllers.AdminLogin())
	incomingRoutes.POST("/adminLogin", adm.LogAdmin())

	incomingRoutes.GET("/profile", controllers.ViewProfile())
	incomingRoutes.POST("/profile", controllers.UnSub())

	incomingRoutes.POST("/users/signup", visit.Signup())
	incomingRoutes.POST("/users/login", visit.Login())
	incomingRoutes.GET("/logout", visit.Logout())

	incomingRoutes.GET("/users/signup", func(context *gin.Context) {
		context.HTML(http.StatusOK, "testRegister.html", gin.H{})
	})
	incomingRoutes.GET("/users/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{})
	})
	incomingRoutes.GET("/postTest", func(context *gin.Context) {
		context.HTML(http.StatusOK, "forum.html", gin.H{})
	})
}
