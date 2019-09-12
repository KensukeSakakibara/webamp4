/*
migration.go
@package github.com/KensukeSakakibara/webamp4/presentation
@author Kensuke Sakakibara
@since 2019.08.28
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note 起動時にGORMを利用してDBのマイグレーションを行います。
各テーブルの定義をinfrastructure/persistence/table以下で行うため例外的にtableに依存しています。
*/
package presentation

import (
	"github.com/KensukeSakakibara/webamp4/infrastructure/persistence/table"
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
