package main

import (
	//Ginの実装
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//Ginの変数
	router := gin.Default()

	//html/css/jsファイルのディレクトリをロード
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"judge": true,
		})
	})

	//ポートを指定して実行
	router.Run(":8000")
}
