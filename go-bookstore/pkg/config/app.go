// gorm is an go ORM --> Object-relational mappers (ORMs) are tools that provide functionalities for interacting with SQL databases using the data types and syntax of the preferred
// programming language. You can use an ORM to operate on databases without the hassle of writing pure SQL because a typical ORM provides database abstraction
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// it will return a variable named db that will help other files to interact with DB
var (
	db *gorm.DB
)

// to open a connection with database
func Connect() {
	d, err := gorm.Open("mysql", "username:password@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// it will return db which will allow other files to talk to the database
func GetDB() *gorm.DB {
	return db
}
