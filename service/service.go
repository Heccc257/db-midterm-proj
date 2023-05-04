package service

import (
	"fmt"
	"net/http"
	"server/service/dal"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Handler func(c *gin.Context) error

var db gorm.DB

func StartUp(db *gorm.DB) error {
	tem, err := dal.NewDB("root", "123456", "pku_mutualhelper")
	if err != nil {
		return err
	}
	db = tem
	return nil
}

func TestHandler(c *gin.Context) {
	fmt.Println("23333333333")
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})

}
