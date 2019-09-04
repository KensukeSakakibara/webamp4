/*
session_repository.go
@import github.com/KensukeSakakibara/gin_gorm_skeleton/domain/repository
@author Kensuke Sakakibara
@since 2019.09.03
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note セッション情報を管理・操作します。インフラ層を隠蔽する役目があります。
*/
package repository

import (
	"github.com/KensukeSakakibara/gin_gorm_skeleton/infrastructure/persistence/session"
	"github.com/gin-gonic/gin"
)

// インターフェイス定義
type SessionRepositoryInterface interface {
	Clear(*gin.Context)
	GetEmail(*gin.Context) string
	SetEmail(*gin.Context, string)
	GetPassword(*gin.Context) string
	SetPassword(*gin.Context, string)
}

type SessionRepository struct {
	session session.SessionInterface
}

// コンストラクタ
func NewSessionRepository(sessionInterface session.SessionInterface) SessionRepositoryInterface {
	return &SessionRepository{session: sessionInterface}
}

// E-Mailの取得
func (this *SessionRepository) GetEmail(ctx *gin.Context) string {
	return this.session.GetString(ctx, "email")
}

// E-Mailの保持
func (this *SessionRepository) SetEmail(ctx *gin.Context, email string) {
	this.session.SetString(ctx, "email", email)
}

// パスワードの取得
func (this *SessionRepository) GetPassword(ctx *gin.Context) string {
	return this.session.GetString(ctx, "password")
}

// パスワードの保持
func (this *SessionRepository) SetPassword(ctx *gin.Context, password string) {
	this.session.SetString(ctx, "password", password)
}

// セッションのクリア
func (this *SessionRepository) Clear(ctx *gin.Context) {
	this.session.Clear(ctx)
}
