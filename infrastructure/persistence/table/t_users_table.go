/*
t_users_table.go
@import github.com/KensukeSakakibara/gin_gorm_skeleton/infrastructure/persistence/table
@author Kensuke Sakakibara
@since 2019.08.28
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note GORMを利用してt_usersテーブルを管理・操作します。
*/
package table

import (
	"github.com/jinzhu/gorm"
)

// t_usersのテーブル構成定義
type TUser struct {
	gorm.Model
	UserType int    `gorm:"type:smallint(6);not null;column:user_type;comment:'1:Admin,2:General'"`
	Email    string `gorm:"type:varchar(255);not null;column:email;unique_index"`
	Password string `gorm:"type:varchar(255);not null;column:password;comment:'SHA256 hash'"`
	Name     string `gorm:"type:varchar(255);not null;column:name"`
}

// インターフェイス定義
type TUsersTableInterface interface {
	SelectRowById(uint) (*TUser, error)
	SelectRowsByLimitOffset(uint, uint) (*[]TUser, error)
	SelectRowByEmail(string) (*TUser, error)
	InsertRow(*TUser) (uint, error)
}

type TUsersTable struct {
	table TableInterface
}

// コンストラクタ
func NewTUsersTable(tableInterface TableInterface) TUsersTableInterface {
	return &TUsersTable{table: tableInterface}
}

// IDを元にデータを一件取得する
func (this *TUsersTable) SelectRowById(id uint) (*TUser, error) {
	var tUser TUser
	err := this.table.Select(func(db *gorm.DB) error {
		return db.Find(&tUser, "id = ?", id).Error
	})
	return &tUser, err
}

// LimitとOffsetからデータを取得する
func (this *TUsersTable) SelectRowsByLimitOffset(limit uint, offset uint) (*[]TUser, error) {
	var tUsers []TUser
	err := this.table.Select(func(db *gorm.DB) error {
		return db.Limit(limit).Offset(offset).Find(&tUsers).Error
	})
	return &tUsers, err
}

// E-mailからデータを一件取得する
func (this *TUsersTable) SelectRowByEmail(email string) (*TUser, error) {
	var tUser TUser
	err := this.table.Select(func(db *gorm.DB) error {
		return db.Find(&tUser, "email = ?", email).Error
	})
	return &tUser, err
}

// 任意のデータを一件追加する
func (this *TUsersTable) InsertRow(tUser *TUser) (uint, error) {
	err := this.table.Insert(func(db *gorm.DB) error {
		return db.Create(tUser).Error
	})
	return tUser.ID, err
}
