/*
common_router.go
@package github.com/KensukeSakakibara/webamp4/interfaces/app/router
@author Kensuke Sakakibara
@since 2019.09.02
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note ルーティング関連の共通処理
*/
package router

import (
	"fmt"
	"github.com/KensukeSakakibara/webamp4/domain/repository"
	"github.com/gin-gonic/gin"
)

// インターフェイス定義
type CommonRouterInterface interface {
	ShowHtml(*gin.Context, string, interface{})
	ShowJson(*gin.Context, interface{})
	Redirect(*gin.Context, string)
	ShowError(*gin.Context, error)
}

type CommonRouter struct {
	tableRepository repository.TableRepositoryInterface
}

// コンストラクタ
func NewCommonRouter(tableRepositoryInterface repository.TableRepositoryInterface) CommonRouterInterface {
	return &CommonRouter{tableRepository: tableRepositoryInterface}
}

// HTML表示
func (this *CommonRouter) ShowHtml(ctx *gin.Context, templateName string, data interface{}) {
	ctx.HTML(200, templateName, gin.H{"data": data})
	this.tableRepository.Close()
}

// JSON表示
func (this *CommonRouter) ShowJson(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, data)
	this.tableRepository.Close()
}

// リダイレクト
func (this *CommonRouter) Redirect(ctx *gin.Context, redirectUrl string) {
	ctx.Redirect(302, redirectUrl)
	this.tableRepository.Close()
}

// Error表示
func (this *CommonRouter) ShowError(ctx *gin.Context, err error) {
	errorMsg := err.Error()
	fmt.Printf(errorMsg)
	ctx.HTML(200, "common_error.html", gin.H{"error_msg": errorMsg})
	this.tableRepository.Close()
}
