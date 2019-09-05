/*
index_usecase.go
@import github.com/KensukeSakakibara/webamp4/application/usecase
@author Kensuke Sakakibara
@since 2019.08.28
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note ユーザー操作に対応した必要な操作をmodelに問い合わせてプレゼンテーション層へ返します。
なるべくusecaseには処理を書かずmodelに書くようにしてください。またusecaseは画面とは結び付きません。
*/
package usecase

import (
	"github.com/KensukeSakakibara/webamp4/domain/model"
	"github.com/KensukeSakakibara/webamp4/domain/repository"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IndexData struct {
	Operator *repository.User
}

type LoginData struct {
	Email string
}

// インターフェイス定義
type IndexUsecaseInterface interface {
	IndexAction(*gin.Context) (*IndexData, error)
	LoginAction(*gin.Context) (bool, *LoginData, error)
	LogoutAction(*gin.Context)
}

type IndexUsecase struct {
	sessionRepository repository.SessionRepositoryInterface
	userModel         model.UserModelInterface
}

// コンストラクタ
func NewIndexUsecase(sessionRepositoryInterface repository.SessionRepositoryInterface,
	userModelInterface model.UserModelInterface) IndexUsecaseInterface {
	return &IndexUsecase{
		sessionRepository: sessionRepositoryInterface,
		userModel:         userModelInterface,
	}
}

// トップ
func (this *IndexUsecase) IndexAction(ctx *gin.Context) (*IndexData, error) {
	user, err := this.userModel.FetchOperatorByEmail(this.sessionRepository.GetEmail(ctx))
	if err != nil {
		return nil, err
	}
	return &IndexData{Operator: user}, nil
}

// ログイン
func (this *IndexUsecase) LoginAction(ctx *gin.Context) (bool, *LoginData, error) {
	// パラメータを取り出す
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	loginBtn, _ := strconv.ParseBool(ctx.PostForm("login_btn"))

	// ログインボタンを押したとき
	if loginBtn {
		// ログイン処理を実行する
		loginFlg, err := this.userModel.CheckLogin(email, password)
		if err != nil {
			return false, nil, err
		}

		// E-mailとパスワードをセッションに保持
		if loginFlg {
			this.sessionRepository.SetEmail(ctx, email)
			this.sessionRepository.SetPassword(ctx, password)

			// リダイレクトしてログイン後の画面を表示させる
			return true, &LoginData{Email: email}, nil

		} else {
			this.sessionRepository.Clear(ctx)
		}
	}

	return false, &LoginData{Email: email}, nil
}

// ログアウトを実行する
func (this *IndexUsecase) LogoutAction(ctx *gin.Context) {
	// セッションをクリア
	this.sessionRepository.Clear(ctx)
}
