package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type bestUsers []bestUser

// Request
type bestUser struct {
	Uid       int64   `gorm:"column:user_id" json:"user_id"`
	NickName  string  `gorm:"column:nick_name" json:"nick_name"`
	AvgRating float64 `gorm:"column:avg_rating" json:"avg_rating"`
}

// 返回task_count>=1且平均分数最高的10个用户
func BestUsers(c *gin.Context) {
	var best_users bestUsers

	// 先选出平均分数最高的最多10个用户
	subQuery1 := db.Table("user_rating").
		Select("rated_user_id AS user_id, AVG(rating) AS avg_rating").
		Group("rated_user_id").
		Order("avg_rating DESC").
		Limit(10)

	// 与user_info表进行连接
	// 子查询，自然连接
	// 等价于以下sql
	/*
		select user_id, nick_name, avg_rating from
		(select rated_user_id as user_id, avg(rating) as avg_rating from user_rating
		group by rated_user_id
		order by avg_rating DESC
		LIMIT 10) AS t1
		NATURAL JOIN user_info AS t2;
	*/
	db.Table("(?) as t1", subQuery1).
		Joins("NATURAL JOIN user_info AS t2").
		Select("user_id, nick_name, avg_rating").Scan(&best_users)
	c.JSON(http.StatusOK, best_users)
}
