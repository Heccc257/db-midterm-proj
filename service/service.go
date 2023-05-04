package service

import (
	"log"
	"net/http"
	"server/service/dal"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Handler func(c *gin.Context) error

var db *gorm.DB

func StartUp(user, password, database_name string) error {
	tem, err := dal.NewDB(user, password, database_name)
	if err != nil {
		return err
	}
	db = tem
	return nil
}

func TestHandler(c *gin.Context) {
	log.Println("23333333333")
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}
