/*
api_router.go
@package github.com/KensukeSakakibara/webamp4/presentation/router
@author Kensuke Sakakibara
@since 2019.09.10
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note Ginを利用してWEBのルーティングを行います。usecaseと表示する画面を結びつけます。
*/
package router

import (
	"github.com/KensukeSakakibara/webamp4/application/usecase"
	"github.com/gin-gonic/gin"
)

// インターフェイス定義
type ApiRouterInterface interface {
	Users(*gin.Context)
}

type ApiRouter struct {
	commonRouter CommonRouterInterface
	apiUsersUsecase usecase.ApiUsersUsecaseInterface
}

// コンストラクタ
func NewApiRouter(commonRouterInterface CommonRouterInterface,
	apiUsersUsecaseInterface usecase.ApiUsersUsecaseInterface) ApiRouterInterface {
	return &ApiRouter{
		commonRouter: commonRouterInterface,
		apiUsersUsecase: apiUsersUsecaseInterface,
	}
}

// ユーザに関するAPI
func (this *ApiRouter) Users(ctx *gin.Context) {
	data, err := this.apiUsersUsecase.UsersAction(ctx)
	if err != nil {
		this.commonRouter.ShowError(ctx, err)
		return
	}
	this.commonRouter.ShowJson(ctx, data)
}
