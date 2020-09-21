package mysql

import (
	"fmt"

	"github.com/KumKeeHyun/web-tuto-with-gin/domain/model"
	"github.com/jinzhu/gorm"
)

const (
	dbms = "mysql"
	user = "root"
	pass = "password"
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
	)
}
