/*
migration.go
@package github.com/KensukeSakakibara/gin_gorm_skeleton/interfaces/app
@author Kensuke Sakakibara
@since 2019.08.28
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note 起動時にGORMを利用してDBのマイグレーションを行います。
各テーブルの定義をinfrastructure/persistence/table以下で行うため例外的にtableに依存しています。
*/
package app

import (
	"github.com/KensukeSakakibara/gin_gorm_skeleton/infrastructure/persistence/table"
)

// インターフェイス定義
type MigrationInterface interface {
	Run()
}

type Migration struct {
	table table.TableInterface
}

// コンストラクタ
func NewMigration(tableInterface table.TableInterface) MigrationInterface {
	return &Migration{table: tableInterface}
}

// DBマイグレーション
func (this *Migration) Run() {
	tableHandler := this.table.MakeMigrateHandler()
	tableHandler.Set("gorm:table_options", "ENGINE=InnoDB")

	// t_user table
	tableHandler.AutoMigrate(&table.TUser{})
}
