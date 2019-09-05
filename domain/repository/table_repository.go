/*
table_repository.go
@import github.com/KensukeSakakibara/webamp4/domain/repository
@author Kensuke Sakakibara
@since 2019.09.03
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note DBセッションを管理・操作します。インフラ層を隠蔽する役目があります。
*/
package repository

import (
	"github.com/KensukeSakakibara/webamp4/infrastructure/persistence/table"
)

// インターフェイス定義
type TableRepositoryInterface interface {
	Close()
	Transaction(func() error) error
}

type TableRepository struct {
	table table.TableInterface
}

// コンストラクタ
func NewTableRepository(tableInterface table.TableInterface) TableRepositoryInterface {
	return &TableRepository{table: tableInterface}
}

// 開いているDBセッションを全て閉じる
func (this *TableRepository) Close() {
	this.table.CloseAllDbHandler()
}

// トランザクションを実行するラッパー
func (this *TableRepository) Transaction(txFunc func() error) error {
	return this.table.Transaction(txFunc)
}
