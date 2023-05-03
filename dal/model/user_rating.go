package model

/*
create table user_rating
(
   rating_id            int unsigned AUTO_INCREMENT not null,
   rater_id             int unsigned not null,
   rated_user_id        int unsigned not null,
   order_id             int unsigned not null,
   rating               int not null,
   comment              text,
   primary key (rating_id)
);
*/

type UserRating struct {
	RatingId    uint   `gorm:"column:rating_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"rating_id"`
	RaterId     uint   `gorm:"column:rater_id;type:int(11) unsigned;NOT NULL" json:"rater_id"`
	RatedUserId uint   `gorm:"column:rated_user_id;type:int(11) unsigned;NOT NULL" json:"rated_user_id"`
	OrderId     uint   `gorm:"column:order_id;type:int(11) unsigned;NOT NULL" json:"order_id"`
	Rating      int    `gorm:"column:rating;type:int(11);NOT NULL" json:"rating"`
	Comment     string `gorm:"column:comment;type:text" json:"comment"`
}

func (m *UserRating) TableName() string {
	return "user_rating"
}
