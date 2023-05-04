package model

type UserRating struct {
	RatingId    uint   `gorm:"column:rating_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"rating_id"`
	RaterId     uint   `gorm:"column:rater_id;type:int(11) unsigned;NOT NULL" json:"rater_id"`
	RatedUserId uint   `gorm:"column:rated_user_id;type:int(11) unsigned;NOT NULL" json:"rated_user_id"`
	OfferId     uint   `gorm:"column:offer_id;type:int(11) unsigned;NOT NULL" json:"offer_id"`
	Rating      int    `gorm:"column:rating;type:int(11);NOT NULL" json:"rating"`
	Comment     string `gorm:"column:comment;type:text" json:"comment"`
}

func (m *UserRating) TableName() string {
	return "user_rating"
}
