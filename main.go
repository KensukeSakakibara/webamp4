/*
main.go
@import github.com/KensukeSakakibara/gin_gorm_skeleton
@author Kensuke Sakakibara
@since 2019.08.28
@copyright Copyright (c) 2019 Kensuke Sakakibara
*/
package main

import (
	"github.com/KensukeSakakibara/gin_gorm_skeleton/registry"
)

func main() {
	// DBマイグレーション
	migrationInterface := registry.DiMigration()
	migrationInterface.Run()

	// アプリケーションの初期化
	initInterface := registry.DiInit()
	initInterface.Run()

	// アプリ実行
	routerInterface := registry.DiRouter()
	routerInterface.Run()
}
