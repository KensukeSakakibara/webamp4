package repository

import (
	"github.com/KensukeSakakibara/webamp4/domain/model"
)

type TAlbumRepository interface {
	GetRowById(id int) (*model.TAlbum, error)
}
