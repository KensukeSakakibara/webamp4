/*
model.go
@import github.com/KensukeSakakibara/gin_gorm_skeleton/domain/model
@author Kensuke Sakakibara
@since 2019.08.29
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note ビジネスロジックの汎用処理。
domainに実装を寄せて実装する必要がありますが、なるべくステートレスにするためセッションは扱いません。
*/
package model

import (
	"crypto/sha256"
	"fmt"
)

// 文字列をSHA256でハッシュ化して文字列で返す
func Sha256(str string) string {
	var byteHash [32]byte
	byteHash = sha256.Sum256([]byte(str))
	return fmt.Sprintf("%x", byteHash)
}
