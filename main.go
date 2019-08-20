package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	//"os"
	//"path/filepath"

	"github.com/KensukeSakakibara/webamp4/infrastructure"
	"github.com/KensukeSakakibara/webamp4/domain/model"
)

func connectGorm() *gorm.DB {
	// コンフィグの読み込み
	dbconfig := config.GetConfig();
	
	// DB接続
	connect := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", dbconfig.Master.Username, dbconfig.Master.Password, dbconfig.Master.Protocol, dbconfig.Master.Database)
	db, err := gorm.Open(dbconfig.Master.Dialect, connect)
	if err != nil {
		log.Println(err.Error())
	}

	return db
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	db := connectGorm()
	defer db.Close()

	router.GET("/", func(ctx *gin.Context) {
		var talbums model.TAlbums
		db.Find(&talbums, "id = ?", 1)
		fmt.Println()

		ctx.HTML(200, "index.html", gin.H{"data": talbums.Name})
	})

	router.Run(":33333")
}
