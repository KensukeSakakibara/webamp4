/*
user_model.go
@import github.com/KensukeSakakibara/webamp4/domain/model
@author Kensuke Sakakibara
@since 2019.08.28
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note ユーザーデータ操作に関するビジネスロジック。
domainに実装を寄せて実装する必要がありますが、なるべくステートレスにするためセッションは扱いません。
*/
package model

import (
	"github.com/KensukeSakakibara/webamp4/domain/repository"
)

// インターフェイス定義
type UserModelInterface interface {
	FetchOperatorByEmail(string) (*repository.User, error)
	CheckRegistByEmail(string) (bool, error)
	CheckLogin(string, string) (bool, error)
	AddAdmin(string, string, string) (uint, error)
}

type UserModel struct {
	tableRepository repository.TableRepositoryInterface
	userRepository  repository.UserRepositoryInterface
}

// コンストラクタ
func NewUserModel(tableRepositoryInterface repository.TableRepositoryInterface,
	userRepositoryInterface repository.UserRepositoryInterface) UserModelInterface {
	return &UserModel{
		tableRepository: tableRepositoryInterface,
		userRepository:  userRepositoryInterface,
	}
}

// E-mailを元にオペレータのデータを取得する
func (this *UserModel) FetchOperatorByEmail(email string) (*repository.User, error) {
	return this.userRepository.FetchUserByEmail(email)
}

// E-mailを元にユーザーが登録されているかを確認する
func (this *UserModel) CheckRegistByEmail(email string) (bool, error) {
	user, err := this.userRepository.FetchUserByEmail(email)
	if user.ID == 0 {
		return false, err
	} else {
		return true, err
	}
}

// ログインチェック処理
func (this *UserModel) CheckLogin(email string, password string) (bool, error) {
	user, err := this.userRepository.FetchUserByEmail(email)
	if err != nil {
		return false, err
	}
	if Sha256(password) == user.Password {
		return true, nil
	} else {
		return false, nil
	}
}

// 管理者を一件追加する
func (this *UserModel) AddAdmin(email string, password string, name string) (uint, error) {
	// 書き込むデータを用意する
	user := repository.User{
		UserType: repository.USER_TYPE_ADMIN,
		Email:    email,
		Password: Sha256(password),
		Name:     name,
	}

	// トランザクションを利用してユーザーデータを書き込む
	var userId uint
	err := this.tableRepository.Transaction(func() (err error) {
		userId, err = this.userRepository.AddUser(&user)
		return err
	})

	return userId, err
}
