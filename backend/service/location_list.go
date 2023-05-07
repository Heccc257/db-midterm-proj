package service

import (
	"fmt"
	"net/http"
	"server/service/dal/model"

	"github.com/gin-gonic/gin"
)

func LocationList(c *gin.Context) {
	var locations []model.Location
	db.Find(&locations)
	fmt.Println(locations)
	c.JSON(http.StatusOK, locations)
}
