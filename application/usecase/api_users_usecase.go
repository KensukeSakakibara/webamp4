/*
api_users_usecase.go
@import github.com/KensukeSakakibara/webamp4/application/usecase
@author Kensuke Sakakibara
@since 2019.09.10
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note ユーザー操作に対応した必要な操作をmodelに問い合わせてプレゼンテーション層へ返します。
なるべくusecaseには処理を書かずmodelに書くようにしてください。またusecaseは画面とは結び付きません。
*/
package usecase

import (
	"github.com/KensukeSakakibara/webamp4/domain/model"
	"github.com/KensukeSakakibara/webamp4/domain/repository"
	"github.com/gin-gonic/gin"
)

type UserData struct {
	Operator *repository.User
}

// インターフェイス定義
type ApiUsersUsecaseInterface interface {
	UsersAction(*gin.Context) (*UserData, error)
}

type ApiUsersUsecase struct {
	sessionRepository repository.SessionRepositoryInterface
	userModel         model.UserModelInterface
}

// コンストラクタ
func NewApiUsersUsecase(sessionRepositoryInterface repository.SessionRepositoryInterface,
	userModelInterface model.UserModelInterface) ApiUsersUsecaseInterface {
	return &ApiUsersUsecase{
		sessionRepository: sessionRepositoryInterface,
		userModel:         userModelInterface,
	}
}

// トップ
func (this *ApiUsersUsecase) UsersAction(ctx *gin.Context) (*UserData, error) {
	user, err := this.userModel.FetchOperatorByEmail(this.sessionRepository.GetEmail(ctx))
	if err != nil {
		return nil, err
	}
	return &UserData{Operator: user}, nil
}
