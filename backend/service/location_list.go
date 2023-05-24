package service

import (
	"server/service/dal/model"

	"github.com/gin-gonic/gin"
)

func LocationList(c *gin.Context) {
	var locations []model.Location
	db.Find(&locations)
	// fmt.Println(locations)
	responseOK(c, locations)
	// c.JSON(http.StatusOK, locations)
}
