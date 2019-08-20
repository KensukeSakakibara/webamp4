package database

import (
	"github.com/KensukeSakakibara/webamp4/domain/model"
	"github.com/KensukeSakakibara/webamp4/domain/repository"
)

type TAlbumsRepository struct {
	*Db
}

func NewTAlbumsRepository() repository.TAlbumsRepository {
	db := NewDb()
	return &TAlbumsRepository{db}
}
