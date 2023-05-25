package service

import (
	"fmt"
	"net/http"
	"server/service/dal/model"

	"github.com/gin-gonic/gin"
)

func Categories(c *gin.Context) {
	var categories []model.OfferCategory
	if err := db.Find(&categories).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(categories)
	responseOK(c, categories)
	// c.JSON(http.StatusOK, categories)
}
