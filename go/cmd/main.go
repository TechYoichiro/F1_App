package main

import (
	"github.com/TechYoichiro/F1_App/controllers"
	"github.com/TechYoichiro/F1_App/interfaces"
	"github.com/TechYoichiro/F1_App/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginフレームワークのデフォルトの設定を使用してルータを作成
	router := gin.Default()

	raceRepo := interfaces.NewRaceRepository()
	raceUsecase := usecase.NewRaceUsecase(raceRepo)
	raceController := controllers.NewRaceController(raceUsecase)

	// ルートハンドラの定義
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// レースデータ取得エンドポイントの定義
	router.GET("/result", raceController.GetRaceData)

	// サーバー起動
	router.Run(":8080")
}
