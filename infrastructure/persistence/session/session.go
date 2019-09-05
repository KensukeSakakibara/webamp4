/*
session.go
@import github.com/KensukeSakakibara/webamp4/infrastructure/persistence/session
@author Kensuke Sakakibara
@since 2019.08.28
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note Sessionを管理する処理です。Sessionの保存はRedisで行います。
application/usecaseより上位のみで利用することを想定しています。
*/
package session

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// インターフェイス定義
type SessionInterface interface {
	GetString(*gin.Context, string) string
	SetString(*gin.Context, string, string)
	Clear(*gin.Context)
}

type Session struct{}

// コンストラクタ
func NewSession() SessionInterface {
	return &Session{}
}

// Keyを元にセッションから文字列を取得する
func (this *Session) GetString(ctx *gin.Context, key string) string {
	session := sessions.Default(ctx)
	var value string
	v := session.Get(key)
	if v == nil {
		value = ""
	} else {
		value = v.(string)
	}
	return value
}

// セッションに文字列を保存する
func (this *Session) SetString(ctx *gin.Context, key string, value string) {
	session := sessions.Default(ctx)
	session.Set(key, value)
	session.Save()
}

// セッションをクリアする
func (this *Session) Clear(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
}
