/*
user_repository.go
@import github.com/KensukeSakakibara/gin_gorm_skeleton/domain/repository
@author Kensuke Sakakibara
@since 2019.09.03
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note ユーザーデータを管理・操作します。インフラ層を隠蔽する役目があります。
基本的にはインフラ層との橋渡し役になりますが、複数テーブルをひとつのリポジトリで操作する場合もあります。
*/
package repository

import (
	"github.com/KensukeSakakibara/gin_gorm_skeleton/infrastructure/persistence/table"
)

const USER_TYPE_ADMIN = 1
const USER_TYPE_GENERAL = 2

// ユーザーデータの構造定義
type User struct {
	ID       uint
	UserType int
	Email    string
	Password string
	Name     string
}

// インターフェイス定義
type UserRepositoryInterface interface {
	FetchUserById(uint) (*User, error)
	FetchUsersByLimitOffset(uint, uint) (*[]User, error)
	FetchUserByEmail(string) (*User, error)
	AddUser(*User) (uint, error)
}

type UserRepository struct {
	tUsersTable table.TUsersTableInterface
}

// コンストラクタ
func NewUserRepository(tUsersTableInterface table.TUsersTableInterface) UserRepositoryInterface {
	return &UserRepository{tUsersTable: tUsersTableInterface}
}

// t_userのデータからユーザデータを生成するメソッド
func (this *UserRepository) makeUserFromTUser(tUser table.TUser) *User {
	return &User{
		ID:       tUser.ID,
		UserType: tUser.UserType,
		Email:    tUser.Email,
		Password: tUser.Password,
		Name:     tUser.Name,
	}
}

// t_userのデータ複数件からユーザーデータ複数件を生成するメソッド
func (this *UserRepository) makeUsersFromTUsers(tUsers []table.TUser) *[]User {
	var users []User
	for _, tUser := range tUsers {
		users = append(users, *this.makeUserFromTUser(tUser))
	}
	return &users
}

// IDからユーザーデータを一件取得する
func (this *UserRepository) FetchUserById(id uint) (*User, error) {
	tUser, err := this.tUsersTable.SelectRowById(id)
	return this.makeUserFromTUser(*tUser), err
}

// リミットとオフセットを利用してユーザーデータを複数件取得する
func (this *UserRepository) FetchUsersByLimitOffset(limit uint, offset uint) (*[]User, error) {
	tUsers, err := this.tUsersTable.SelectRowsByLimitOffset(limit, offset)
	return this.makeUsersFromTUsers(*tUsers), err
}

// E-mailからユーザーデータを一件取得する
func (this *UserRepository) FetchUserByEmail(email string) (*User, error) {
	tUser, err := this.tUsersTable.SelectRowByEmail(email)
	return this.makeUserFromTUser(*tUser), err
}

// ユーザーデータを一件保存する
func (this *UserRepository) AddUser(user *User) (uint, error) {
	tUser := table.TUser{
		UserType: user.UserType,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
	}
	return this.tUsersTable.InsertRow(&tUser)
}
