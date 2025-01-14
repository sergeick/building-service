package main

import (
	_ "building-service/docs"
	"building-service/internal/handlers"
	"building-service/internal/repository"
	"building-service/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Building Service API
// @version 1.0
// @description Service for managing buildings
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Загрузка конфигурации
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// Подключение к базе данных
	db, err := database.NewPostgresDB(database.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.dbname"),
	})
	if err != nil {
		log.Fatalf("Failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	// Настройка маршрутов
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		buildingRepo := repository.NewBuildingRepository(db)
		buildingHandler := handlers.NewBuildingHandler(buildingRepo)

		api.POST("/buildings", buildingHandler.Create)
		api.GET("/buildings", buildingHandler.List)

		// Добавление маршрута для эндпоинта /docs
		api.GET("/docs", func(c *gin.Context) {
			c.Redirect(302, "/swagger/doc.json")
		})

	}

	// Настройка Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}
}
