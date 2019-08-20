package main

import (
	"github.com/gin-gonic/gin"
	"github.com/KensukeSakakibara/webamp4/infrastructure/persistence/database"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	talbumRepo := database.NewTAlbumRepository()
	talbum, err := talbumRepo.GetRowById(1)
	if err != nil {
		return
	}
	
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": talbum.Name})
	})

	router.Run(":33333")
}
