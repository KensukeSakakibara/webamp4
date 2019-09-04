/*
table.go
@import github.com/KensukeSakakibara/gin_gorm_skeleton/infrastructure/persistence/table
@author Kensuke Sakakibara
@since 2019.08.28
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note GORMを利用してテーブルを管理・操作する汎用処理です。
DBのレプリケーションの対応もこちらで行います。
*/
package table

import (
	"fmt"
	"github.com/KensukeSakakibara/gin_gorm_skeleton/infrastructure/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"math/rand"
	"time"
)

// インターフェイス定義
type TableInterface interface {
	MakeMigrateHandler() *gorm.DB
	UseMaster()
	Select(func(*gorm.DB) error) error
	Insert(func(*gorm.DB) error) error
	Update(func(*gorm.DB) error) error
	Delete(func(*gorm.DB) error) error
	Transaction(func() error) error
	CloseAllDbHandler()
}

type Table struct {
	isSet              bool
	config             *config.Config
	isMaster           bool
	masterHandler      *gorm.DB
	slaveHandlers      []*gorm.DB
	transactionHandler *gorm.DB
}

var table *Table = new(Table)

// コンストラクタ
func NewTable(configInstance *config.Config) TableInterface {
	if table.isSet {
		return table
	}
	table = &Table{isSet: true, config: configInstance, isMaster: false}
	return table
}

// Migrate用のDBアダプタを用意する
func (this *Table) MakeMigrateHandler() *gorm.DB {
	return this.makeAdapter(true)
}

// マスターを参照
func (this *Table) UseMaster() {
	this.isMaster = true
}

// Selectを実行する
func (this *Table) Select(selectFunc func(*gorm.DB) error) error {
	var dbHandler *gorm.DB
	if this.isMaster {
		dbHandler = this.makeAdapter(true)
		this.isMaster = false
	} else {
		dbHandler = this.makeAdapter(false)
	}

	// 単にレコードが見つからない場合はエラーにしない
	if err := selectFunc(dbHandler); gorm.IsRecordNotFoundError(err) {
		return nil
	} else {
		return err
	}
}

// Insertを実行する
func (this *Table) Insert(insertFunc func(*gorm.DB) error) error {
	var dbHandler *gorm.DB
	if this.transactionHandler != nil {
		dbHandler = this.transactionHandler
	} else {
		dbHandler = this.makeAdapter(true)
	}
	return insertFunc(dbHandler)
}

// Updateを実行する
func (this *Table) Update(updateFunc func(*gorm.DB) error) error {
	var dbHandler *gorm.DB
	if this.transactionHandler != nil {
		dbHandler = this.transactionHandler
	} else {
		dbHandler = this.makeAdapter(true)
	}
	return updateFunc(dbHandler)
}

// Deleteを実行する
func (this *Table) Delete(deleteFunc func(*gorm.DB) error) error {
	var dbHandler *gorm.DB
	if this.transactionHandler != nil {
		dbHandler = this.transactionHandler
	} else {
		dbHandler = this.makeAdapter(true)
	}
	return deleteFunc(dbHandler)
}

func (this *Table) makeAdapter(masterFlg bool) *gorm.DB {
	if masterFlg {
		// マスターの場合
		if this.masterHandler != nil {
			// 既に保持しているのであればリターンして終了
			return this.masterHandler
		}

		// DBハンドラを作成して保持する
		adapter := this.config.Db.Master
		dbHandler := this.makeDbHandler(adapter)
		this.masterHandler = dbHandler
		return dbHandler

	} else {
		// スレーブの場合
		slaveCount := len(this.config.Db.Slaves)
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(slaveCount)

		// DBハンドラを作成して保持する
		if len(this.slaveHandlers) != 0 && this.slaveHandlers[randNum] != nil {
			// 既に保持しているのであればリターンして終了
			return this.slaveHandlers[randNum]
		}

		// スレーブのDBハンドラを作成して保持する
		adapter := this.config.Db.Slaves[randNum]
		dbHandler := this.makeDbHandler(adapter)
		if len(this.slaveHandlers) == 0 {
			this.slaveHandlers = make([]*gorm.DB, slaveCount)
		}
		this.slaveHandlers[randNum] = dbHandler

		return dbHandler
	}
}

// The returned DB is safe for concurrent use by multiple goroutines
// and maintains its own pool of idle connections.
// Thus, the Open function should be called just once.
// It is rarely necessary to close a DB.
// 参照：https://golang.org/pkg/database/sql/#Open
func (this *Table) makeDbHandler(adapter config.Adapter) *gorm.DB {
	// DB接続
	username := adapter.Username
	password := adapter.Password
	protocol := adapter.Protocol
	database := adapter.Database
	connect := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", username, password, protocol, database)

	dbHandler, err := gorm.Open(adapter.Dialect, connect)
	if err != nil {
		log.Println(err.Error())
	}

	return dbHandler
}

// トランザクション用のラッパー関数
func (this *Table) Transaction(txFunc func() error) error {
	// トランザクション開始
	db := this.makeAdapter(true)
	this.transactionHandler = db.Begin()
	defer func() {
		if r := recover(); r != nil {
			this.transactionHandler.Rollback()
			this.transactionHandler = nil
		}
	}()

	// エラーがあればエラーを返す
	if err := this.transactionHandler.Error; err != nil {
		this.transactionHandler = nil
		return err
	}

	// ユーザーの関数を実行
	if err := txFunc(); err != nil {
		this.transactionHandler.Rollback()
		this.transactionHandler = nil
		return err
	}

	// コミット
	err := this.transactionHandler.Commit().Error
	this.transactionHandler = nil
	return err
}

// 全てのDBコネクションを終了させる（リクエスト処理の終了時に実行される）
func (this *Table) CloseAllDbHandler() {
	if this.masterHandler != nil {
		this.masterHandler.Close()
	}
	if len(this.slaveHandlers) != 0 {
		for i, slaveHandler := range this.slaveHandlers {
			if slaveHandler != nil {
				this.slaveHandlers[i].Close()
				this.slaveHandlers[i] = nil
			}
		}
	}
	this.isMaster = false
}
