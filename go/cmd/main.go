package main

import (
	"log"

	"github.com/TechYoichiro/F1_App/controllers"
	"github.com/TechYoichiro/F1_App/interfaces"
	"github.com/TechYoichiro/F1_App/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginフレームワークのデフォルトの設定を使用してルータを作成
	router := gin.Default()

	// CORSを設定
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// レースリポジトリ、ユースケース、コントローラを初期化
	raceRepo := interfaces.NewRaceRepository()
	raceUsecase := usecase.NewRaceUsecase(raceRepo)
	raceController := controllers.NewRaceController(raceUsecase)

	// ルートハンドラの定義（"/" エンドポイント）
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// 最新のレースデータを取得するエンドポイントの定義（"/result" エンドポイント）
	router.GET("/result", raceController.GetRaceData)

	// サーバー起動
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not run the server: %v", err)
	}
}
