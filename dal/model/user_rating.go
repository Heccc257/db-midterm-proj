package model

/*
create table user_rating
(
   rating_id            int not null,
   rater_id             int not null,
   rated_user_id        int not null,
   order_id             int not null,
   rating               int not null,
   comment              text,
   primary key (rating_id)
);
*/

type UserRating struct {
	RatingId    int    `gorm:"column:rating_id;type:int(11);primary_key" json:"rating_id"`
	RaterId     int    `gorm:"column:rater_id;type:int(11);NOT NULL" json:"rater_id"`
	RatedUserId int    `gorm:"column:rated_user_id;type:int(11);NOT NULL" json:"rated_user_id"`
	OrderId     int    `gorm:"column:order_id;type:int(11);NOT NULL" json:"order_id"`
	Rating      int    `gorm:"column:rating;type:int(11);NOT NULL" json:"rating"`
	Comment     string `gorm:"column:comment;type:text" json:"comment"`
}

func (m *UserRating) TableName() string {
	return "user_rating"
}
