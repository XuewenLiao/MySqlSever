package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

//创建Database实例
//参数：数据库用户名，密码，IP地址，数据库名
func (db *Database) Init(username string, password string, addr string, port int, name string) {
	DB = &Database{
		Self: GetSelfDB(username, password, addr, port, name),
	}
}

func GetSelfDB(username string, password string, addr string, port int, name string) *gorm.DB {
	return openDB(username, password, addr, port, name)
}

//建立数据库连接
func openDB(username string, password string, addr string, port int, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		port,
		name,
		true,
		"Local")

	//config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
	//	username,
	//	password,
	//	addr,
	//	port,
	//	name,
	//	)

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Fatal(err, "Database connection failed. Database name: %s", name)
	} else {
		log.Printf("Database connection Success !")
	}

	return db
}

func (db *Database) Close() {
	db.Self.Close()
}
