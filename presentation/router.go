/*
router.go
@package github.com/KensukeSakakibara/webamp4/presentation
@author Kensuke Sakakibara
@since 2019.08.28
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note Ginを利用してWEBのルーティングを行います。usecaseと表示する画面を結びつけます。
*/
package presentation

import (
	"github.com/KensukeSakakibara/webamp4/presentation/router"
	"github.com/KensukeSakakibara/webamp4/infrastructure/config"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// インターフェイス定義
type RouterInterface interface {
	Run()
}

type Router struct {
	config         *config.Config
	store          *sessions.RedisStore
	appRouter router.AppRouterInterface
	apiRouter router.ApiRouterInterface
}

// コンストラクタ
func NewRouter(configInstance *config.Config, redisStore *sessions.RedisStore,
	appRouterInterface router.AppRouterInterface,
	apiRouterInterface router.ApiRouterInterface) RouterInterface {
	return &Router{
		config:         configInstance,
		store:          redisStore,
		appRouter: appRouterInterface,
		apiRouter: apiRouterInterface,
	}
}

func (this *Router) Run() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.html")

	// 静的ファイル
	router.Static("/static", "./static")

	// セッション対応 (Redis)
	router.Use(sessions.Sessions(this.config.App.SessionName, *this.store))

	// ルーティング
	appRoot := router.Group("/app")
	{
		appRoot.GET("/", this.appRouter.Index)
		appRoot.POST("/", this.appRouter.Index)
		appRoot.GET("/index/logout", this.appRouter.Logout)
	}

	apiRoot := router.Group("/api")
	{
		apiRoot.GET("/users/:id", this.apiRouter.Users)
	}

	router.Run(":" + this.config.App.ApplicationPort)
}
