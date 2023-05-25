package service

import (
	"net/http"
	"server/service/dal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userRatingForm struct {
	Comment string `form:"comment"`
	OfferID int64  `form:"offer_id"`
	Rating  int64  `form:"rating"`
}

func UserRating(c *gin.Context) {
	var userRating model.UserRating
	if uid, err := strconv.Atoi(c.Param("uid")); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		userRating.RaterId = uint(uid)
	}
	token := c.Param("token")
	if token != "access" {
		// fmt.Println("token = ", token, tokens[token])
		if uid, ok := tokens[token]; !ok || uid != userRating.RaterId {
			responseBadRequest(c, "token验证失败")
			// c.String(http.StatusBadRequest, "token验证失败")
			return
		}
	}
	var rating_info userRatingForm
	c.ShouldBind(&rating_info)
	// fmt.Println("uid ", userRating.RaterId)
	// fmt.Printf("rating info: %+v\n", rating_info)
	// RatedUserId通过查表获取
	var accept_offer model.AcceptOffer
	db.Model(&model.AcceptOffer{}).Where("offer_id = ?", rating_info.OfferID).Find(&accept_offer)
	// fmt.Println("offer id = ", rating_info.OfferID, "accept_offer = ", accept_offer)
	if accept_offer.OfferId != uint(rating_info.OfferID) {
		responseBadRequest(c, "评论失败")
		// c.String(http.StatusBadRequest, "评论失败")
		return
	}
	userRating.Comment = rating_info.Comment
	userRating.OfferId = uint(rating_info.OfferID)
	userRating.Rating = int(rating_info.Rating)
	userRating.RatedUserId = accept_offer.UserId
	db.Create(&userRating)
	response(c, http.StatusOK, "评论成功", nil)
	// c.String(http.StatusOK, "评论成功")
}
