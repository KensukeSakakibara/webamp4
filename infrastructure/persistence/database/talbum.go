package database

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/KensukeSakakibara/webamp4/domain/model"
	"github.com/KensukeSakakibara/webamp4/domain/repository"
	"github.com/KensukeSakakibara/webamp4/infrastructure/config"
)

type TAlbumRepository struct {
	db *gorm.DB
}

func NewTAlbumRepository() repository.TAlbumRepository {
	db := NewDb()
	return &TAlbumRepository{db}
}

func NewDb() *gorm.DB {
	// コンフィグの読み込み
	dbconfig := config.GetConfig()

	// DB接続
	connect := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", dbconfig.Master.Username, dbconfig.Master.Password, dbconfig.Master.Protocol, dbconfig.Master.Database)
	db, err := gorm.Open(dbconfig.Master.Dialect, connect)
	if err != nil {
		log.Println(err.Error())
	}

	return db
}

func (r *TAlbumRepository) GetRowById(id int) (*model.TAlbum, error) {
	var talbum model.TAlbum
	r.db.Find(&talbum, "id = ?", id)
	return &talbum, nil
}
