package dal

import (
	"fmt"
	"server/dal/model"
	"testing"
)

func TestCreate(t *testing.T) {
	db, _ := NewDB("root", "123456", "bank")
	if err := db.Exec("create database if not exists bank;"); err != nil {
		fmt.Println(err)
	}
	db.Exec("use bank;")
	db.AutoMigrate(&model.Book{})
	// db.Exec("drop database if exists bank;")
}

func TestForeignKey(t *testing.T) {
	db, _ := NewDB("root", "123456", "bank")
	db.Exec("drop database if exists bank;")
	db.Exec("create database if not exists bank;")
	db.Exec("use bank;")
	db.AutoMigrate(&model.Book{})
	db.Create(&model.Book{
		Name: "abc",
		Type: "000",
	})
	db.AutoMigrate(&model.Borrow{})
	db.Create(&model.Borrow{
		Id:       1,
		BookName: "abc",
	})
	db.Exec(`
	ALTER TABLE borrow
	ADD CONSTRAINT fk_borrow_book
	FOREIGN KEY (book_name) REFERENCES book (name)
	ON DELETE CASCADE;
	`)
	fmt.Println(db)
}
