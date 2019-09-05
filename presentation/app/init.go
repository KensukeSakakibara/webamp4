/*
init.go
@package github.com/KensukeSakakibara/webamp4/interfaces/app
@author Kensuke Sakakibara
@since 2019.08.28
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note 起動時にアプリケーションの初期化を行います。
*/
package app

import (
	"github.com/KensukeSakakibara/webamp4/domain/model"
	"github.com/KensukeSakakibara/webamp4/infrastructure/config"
)

// インターフェイス定義
type InitInterface interface {
	Run()
}

type Init struct {
	config    *config.Config
	userModel model.UserModelInterface
}

// コンストラクタ
func NewInit(configInstance *config.Config, userModelInterface model.UserModelInterface) InitInterface {
	return &Init{config: configInstance, userModel: userModelInterface}
}

func (this *Init) Run() {
	// コンフィグから管理者情報を取得する
	adminEmail := this.config.App.AdminEmail
	adminPassword := this.config.App.AdminPassword
	adminName := this.config.App.AdminName

	// 管理者が登録されて無ければ作成する
	isRegist, _ := this.userModel.CheckRegistByEmail(adminEmail)
	if !isRegist {
		this.userModel.AddAdmin(adminEmail, adminPassword, adminName)
	}
}
