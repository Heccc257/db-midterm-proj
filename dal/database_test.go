package dal

import (
	"server/dal/model"
	"testing"
)

func TestNewDB(t *testing.T) {
	NewDB("root", "123456", "pku_mutualhelper")
	// db.Exec("drop database if exists bank;")
}

func TestCreateTables(t *testing.T) {
	db, _ := NewDB("root", "123456", "pku_mutualhelper")
	db.AutoMigrate(&model.User{})

}
