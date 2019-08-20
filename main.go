package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	"github.com/KensukeSakakibara/webamp4/domain/model"
	"github.com/KensukeSakakibara/webamp4/infrastructure/config"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	talbumRepo := NewTAlbumRepository()
	talbum := talbumRepo.GetRowById(1)

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": talbum.Name})
	})

	router.Run(":33333")
}
