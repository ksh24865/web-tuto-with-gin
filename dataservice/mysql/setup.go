package mysql

import (
	"fmt"

	"github.com/KumKeeHyun/web-tuto-with-gin/domain/model"
	"github.com/jinzhu/gorm"
)

const (
	dbms = "mysql"
	user = "root"
	pass = "68gkrqjs"
	db   = "webtuto"
)

var dbConn *gorm.DB

func Setup() {
	var err error
	conn := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, pass, db)
	dbConn, err = gorm.Open(dbms, conn)
	if err != nil {
		panic(err)
	}

	dbConn.AutoMigrate(
		&model.Article{},
		&model.User{},
	)
	dbConn.Model(&model.Article{}).AddForeignKey("writer_id", "users(id)", "CASCADE", "CASCADE")
}
