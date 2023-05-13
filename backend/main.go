package main

import (
	"flag"
	"fmt"
	"net/http"
	"server/service"

	"github.com/gin-gonic/gin"
)

var (
	port          string
	addr          string
	user          string
	password      string
	database_name string
)

func main() {
	flag.StringVar(&addr, "addr", "127.0.0.1", "address to listen")
	flag.StringVar(&port, "port", "9999", "port to listen")
	flag.StringVar(&user, "u", "root", "mysql user")
	flag.StringVar(&password, "p", "123456", "mysql password")
	flag.StringVar(&database_name, "db_name", "pku_mutualhelper", "mysql database name")
	flag.Parse()

	// gin.SetMode(gin.ReleaseMode)
	S := gin.Default()
	service.StartUp(user, password, database_name)

	S.GET("/", service.TestHandler)

	S.POST("/register", service.Register)

	S.GET("/login", service.Login)

	S.GET("/user_info/:uid", service.User_info)

	S.PUT("/clear_database", func(ctx *gin.Context) {
		service.StartUp(user, password, database_name)
		ctx.JSON(http.StatusOK, "database cleared")
	})

	offer := S.Group("/offer")
	{
		offer.POST("/post/:token", service.OfferPost)

		offerList := offer.Group("/offer_list")

		offerList.GET("", service.OfferListRandom)
		offerList.GET("/user", service.OfferListByUser)
	}

	S.GET("/categories", service.Categories)

	S.GET("/location_list", service.LocationList)

	S.POST("/accept_offer", service.AcceptOffer)
	S.PUT("/complete_offer", service.CompleteOffer)

	S.POST("/user_rating/:uid", service.UserRating)

	S.GET("/best_users", service.BestUsers)

	err := S.Run(addr + ":" + port)
	if err != nil {
		fmt.Println("服务器启动失败! ", err)
	}
}
