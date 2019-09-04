/*
index_router.go
@package github.com/KensukeSakakibara/gin_gorm_skeleton/interfaces/app/router
@author Kensuke Sakakibara
@since 2019.08.30
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note Ginを利用してWEBのルーティングを行います。usecaseと表示する画面を結びつけます。
*/
package router

import (
	"github.com/KensukeSakakibara/gin_gorm_skeleton/application/usecase"
	"github.com/gin-gonic/gin"
)

// インターフェイス定義
type IndexRouterInterface interface {
	Index(*gin.Context)
	Login(*gin.Context)
	Logout(*gin.Context)
}

type IndexRouter struct {
	commonRouter CommonRouterInterface
	usecase      usecase.UsecaseInterface
	indexUsecase usecase.IndexUsecaseInterface
}

// コンストラクタ
func NewIndexRouter(commonRouterInterface CommonRouterInterface, usecaseInterface usecase.UsecaseInterface,
	indexUsecaseInterface usecase.IndexUsecaseInterface) IndexRouterInterface {
	return &IndexRouter{
		commonRouter: commonRouterInterface,
		usecase:      usecaseInterface,
		indexUsecase: indexUsecaseInterface,
	}
}

// トップ表示
func (this *IndexRouter) Index(ctx *gin.Context) {
	loginFlg, err := this.usecase.SessionLogin(ctx)
	if err != nil {
		this.commonRouter.ShowError(ctx, err)
		return
	}

	// ログインできていなければログイン画面を表示する
	if !loginFlg {
		this.Login(ctx)
		return
	}

	data, err := this.indexUsecase.IndexAction(ctx)
	if err != nil {
		this.commonRouter.ShowError(ctx, err)
		return
	}

	this.commonRouter.ShowHtml(ctx, "index_index.html", data)
}

// ログイン表示
func (this *IndexRouter) Login(ctx *gin.Context) {
	loginFlg, data, err := this.indexUsecase.LoginAction(ctx)
	if err != nil {
		this.commonRouter.ShowError(ctx, err)
		return
	}

	if loginFlg {
		this.commonRouter.Redirect(ctx, "/")
	} else {
		this.commonRouter.ShowHtml(ctx, "index_login.html", data)
	}
}

// ログアウト
func (this *IndexRouter) Logout(ctx *gin.Context) {
	this.indexUsecase.LogoutAction(ctx)
	this.commonRouter.Redirect(ctx, "/")
}
