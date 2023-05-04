package service

import (
	"log"
	"net/http"
	"server/service/dal/crud"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	phone_number := c.Query("phone_number")
	log.Println("phone number = ", phone_number)
	u := crud.FindUserByPhongNumber(db, phone_number)
	if u != nil {
		c.String(http.StatusOK, "user exist")
	}

}
