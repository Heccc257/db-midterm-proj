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
	flag.StringVar(&addr, "addr", "0.0.0.0", "address to listen")
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

	S.POST("/login", service.Login)

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
		offerList.GET("/:uid", service.OfferListByUser)
	}

	S.GET("/categories", service.Categories)

	S.GET("/location_list", service.LocationList)

	S.POST("/accept_offer/:token", service.AcceptOffer)
	S.PUT("/complete_offer/:token", service.CompleteOffer)
	S.GET("accept_offer_list/:uid", service.AcceptOfferList)

	S.POST("/user_rating/:uid/:token", service.UserRating)

	S.GET("/best_users", service.BestUsers)

	//
	S.GET("/offer_list_by_cat/:category", service.OfferListByCat) //按分类获取
	S.PUT("/change_usr_info/:user_id", service.ChangeUserInfo) //修改用户信息


	err := S.Run(addr + ":" + port)
	if err != nil {
		fmt.Println("服务器启动失败! ", err)
	}
}
