package repository

import (
	"github.com/KensukeSakakibara/webamp4/domain/model"
)

type TAlbumsRepository interface {
	SetMaster(isMaster bool)
}
