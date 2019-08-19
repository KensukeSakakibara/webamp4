package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	//"os"
	"io/ioutil"
	//"path/filepath"
	"time"
)

func connectGorm() *gorm.DB {
	// コンフィグの読み込み
	raw, err := ioutil.ReadFile("./config/db.json")
	if err != nil {
		log.Fatal(err)
	}

	var dbconfig DbConfig
	json.Unmarshal(raw, &dbconfig)
	
	// DB接続
	connect := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", dbconfig.Master.Username, dbconfig.Master.Password, dbconfig.Master.Protocol, dbconfig.Master.Database)
	db, err := gorm.Open(dbconfig.Master.Dialect, connect)
	if err != nil {
		log.Println(err.Error())
	}

	return db
}

type TAlbums struct {
	ID           int        `json:"id"`                                    // id
	UserId       int        `json:"user_id"`                               // user_id
	OrderNum     int        `json:"order_num"`                             // order_num
	Name         string     `json:"name"`                                  // name
	CreatedAt    time.Time  `json:"create_date" gorm:"column:create_date"` // create_date
	CreateUserID int        `json:"create_user_id"`                        // create_user_id
	UpdatedAt    time.Time  `json:"update_date" gorm:"column:update_date"` // update_date
	UpdateUserID int        `json:"update_user_id"`                        // update_user_id
	DeletedAt    *time.Time `json:"delete_date" gorm:"column:delete_date"` // delete_date
	DeleteUserID int        `json:"delete_user_id"`                        // delete_user_id
	DeleteFlg    bool       `json:"delete_flg"`                            // delete_flg
}

type Adapter struct {
	Dialect  string `json:"dialect"`
	Username string `json:"username"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Database string `json:"database"`
}

type DbConfig struct {
	Master Adapter   `json:"master"`
	Slaves []Adapter `json:"slaves"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	db := connectGorm()
	defer db.Close()

	router.GET("/", func(ctx *gin.Context) {
		var talbums TAlbums
		db.Find(&talbums, "id = ?", 1)
		fmt.Println()

		ctx.HTML(200, "index.html", gin.H{"data": talbums.Name})
	})

	router.Run(":33333")
}
