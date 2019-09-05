/*
usecase.go
@import github.com/KensukeSakakibara/webamp4/application/usecase
@author Kensuke Sakakibara
@since 2019.08.30
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note usecaseの汎用処理
*/
package usecase

import (
	"github.com/KensukeSakakibara/webamp4/domain/model"
	"github.com/KensukeSakakibara/webamp4/domain/repository"
	"github.com/gin-gonic/gin"
)

// インターフェイス定義
type UsecaseInterface interface {
	SessionLogin(*gin.Context) (bool, error)
}

type Usecase struct {
	session   repository.SessionRepositoryInterface
	userModel model.UserModelInterface
}

// コンストラクタ
func NewUsecase(sessionRepositoryInterface repository.SessionRepositoryInterface,
	userModelInterface model.UserModelInterface) UsecaseInterface {
	return &Usecase{
		session:   sessionRepositoryInterface,
		userModel: userModelInterface,
	}
}

// セッションに保持されたE-mailとパスワードからログインを試みる
func (this *Usecase) SessionLogin(ctx *gin.Context) (bool, error) {
	// セッションからE-mailとパスワードを取り出す
	email := this.session.GetEmail(ctx)
	password := this.session.GetPassword(ctx)

	// ログインを試行する
	return this.userModel.CheckLogin(email, password)
}
