/*
app_router.go
@package github.com/KensukeSakakibara/webamp4/presentation/router
@author Kensuke Sakakibara
@since 2019.08.30
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note Ginを利用してWEBのルーティングを行います。usecaseと表示する画面を結びつけます。
*/
package router

import (
	"github.com/KensukeSakakibara/webamp4/application/usecase"
	"github.com/gin-gonic/gin"
)

// インターフェイス定義
type AppRouterInterface interface {
	Index(*gin.Context)
	Login(*gin.Context)
	Logout(*gin.Context)
}

type AppRouter struct {
	commonRouter    CommonRouterInterface
	usecase         usecase.UsecaseInterface
	appIndexUsecase usecase.AppIndexUsecaseInterface
}

// コンストラクタ
func NewAppRouter(commonRouterInterface CommonRouterInterface, usecaseInterface usecase.UsecaseInterface,
	appIndexUsecaseInterface usecase.AppIndexUsecaseInterface) AppRouterInterface {
	return &AppRouter{
		commonRouter:    commonRouterInterface,
		usecase:         usecaseInterface,
		appIndexUsecase: appIndexUsecaseInterface,
	}
}

// トップ表示
func (this *AppRouter) Index(ctx *gin.Context) {
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

	data, err := this.appIndexUsecase.IndexAction(ctx)
	if err != nil {
		this.commonRouter.ShowError(ctx, err)
		return
	}

	this.commonRouter.ShowHtml(ctx, "index_index.html", data)
}

// ログイン表示
func (this *AppRouter) Login(ctx *gin.Context) {
	loginFlg, data, err := this.appIndexUsecase.LoginAction(ctx)
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
func (this *AppRouter) Logout(ctx *gin.Context) {
	this.appIndexUsecase.LogoutAction(ctx)
	this.commonRouter.Redirect(ctx, "/")
}
