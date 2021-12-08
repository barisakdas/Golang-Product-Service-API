package entity

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Dsn string = "host=localhost user=admin password=Acaba1234!** dbname=hb port=5433 sslmode=disable TimeZone=Asia/Shanghai"

func DbMigration() {
	Campaign{}.Migrate()
	Order{}.Migrate()
	Product{}.Migrate()
}

func CreateDb() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5433 user=admin dbname=postgres password=Acaba1234!** sslmode=disable")
	db = db.Exec("CREATE DATABASE hb;")
	if db.Error != nil {
		fmt.Println("Unable to create DB test_db, attempting to connect assuming it exists...")
	}
	if err != nil {

	}
}
