package model

import "database/sql"

/*
CREATE TABLE `book`  (
`name` varchar(10) NOT NULL,
`type` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
`created_at` timestamp NULL DEFAULT NULL,
`updated_at` timestamp NULL DEFAULT NULL,
PRIMARY KEY (`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = DYNAMIC;

CREATE TABLE `borrow`  (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`bookname` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
`created_at` timestamp NULL DEFAULT NULL,
`updated_at` timestamp NULL DEFAULT NULL,
PRIMARY KEY (`id`) USING BTREE,
CONSTRAINT fk_borrow_book FOREIGN KEY(book_name) REFERENCES book(name) ON DELETE SET NULL
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = DYNAMIC;


*/
type Book struct {
	Name      string       `gorm:"column:name;type:varchar(20);primary_key" json:"name"`
	Type      string       `gorm:"column:type;type:varchar(20);NOT NULL" json:"type"`
	CreatedAt sql.NullTime `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

func (m *Book) TableName() string {
	return "book"
}

type Borrow struct {
	Id        uint         `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	BookName  string       `gorm:"column:book_name;type:varchar(10);NOT NULL" json:"book_name"`
	CreatedAt sql.NullTime `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

func (m *Borrow) TableName() string {
	return "borrow"
}
