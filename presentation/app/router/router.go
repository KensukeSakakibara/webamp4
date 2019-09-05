/*
router.go
@package github.com/KensukeSakakibara/webamp4/interfaces/app/router
@author Kensuke Sakakibara
@since 2019.08.28
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note Ginを利用してWEBのルーティングを行います。usecaseと表示する画面を結びつけます。
*/
package router

import (
	"github.com/KensukeSakakibara/webamp4/infrastructure/config"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// インターフェイス定義
type RouterInterface interface {
	Run()
}

type Router struct {
	config      *config.Config
	store       *sessions.RedisStore
	indexRouter IndexRouterInterface
}

// コンストラクタ
func NewRouter(configInstance *config.Config, redisStore *sessions.RedisStore,
	indexRouterInterface IndexRouterInterface) RouterInterface {
	return &Router{
		config:      configInstance,
		store:       redisStore,
		indexRouter: indexRouterInterface,
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
	defaultRoot := router.Group("/")
	{
		defaultRoot.GET("/", this.indexRouter.Index)
		defaultRoot.POST("/", this.indexRouter.Index)
	}

	index := router.Group("/index")
	{
		index.GET("/logout", this.indexRouter.Logout)
	}

	router.Run(":" + this.config.App.ApplicationPort)
}
