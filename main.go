// package main

// import (
// 	_ "building-service/docs"
// 	"building-service/internal/handlers"
// 	"building-service/internal/repository"
// 	"building-service/pkg/database"
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/lib/pq"
// 	"github.com/spf13/viper"
// 	swaggerFiles "github.com/swaggo/files"
// 	ginSwagger "github.com/swaggo/gin-swagger"
// )

// // @title Building Service API
// // @version 1.0
// // @description Service for managing buildings
// // @host localhost:8080
// // @BasePath /api/v1
// func main() {
// 	// Загрузка конфигурации
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath("./config")

// 	if err := viper.ReadInConfig(); err != nil {
// 		log.Fatalf("Error reading config file: %s", err)
// 	}

// 	// Подключение к базе данных
// 	db, err := database.NewPostgresDB(database.Config{
// 		Host:     viper.GetString("database.host"),
// 		Port:     viper.GetInt("database.port"),
// 		User:     viper.GetString("database.user"),
// 		Password: viper.GetString("database.password"),
// 		DBName:   viper.GetString("database.dbname"),
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed to initialize db: %s", err.Error())
// 	}
// 	defer db.Close()

// 	// Настройка маршрутов
// 	router := gin.Default()
// 	api := router.Group("/api/v1")
// 	{
// 		buildingRepo := repository.NewBuildingRepository(db)
// 		buildingHandler := handlers.NewBuildingHandler(buildingRepo)

// 		api.POST("/buildings", buildingHandler.Create)
// 		api.GET("/buildings", buildingHandler.List)

// 		// Добавление маршрута для эндпоинта /docs
// 		api.GET("/docs", func(c *gin.Context) {
// 			c.Redirect(302, "/swagger/doc.json")
// 		})

// 	}

// 	// Настройка Swagger
// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// 	// Запуск сервера
// 	if err := router.Run(":8080"); err != nil {
// 		log.Fatalf("Failed to run server: %s", err)
// 	}
// }

package main // Пакет для основной программы
// Подчеркивание позволяет импортировать пакет исключительно для его побочных эффектов, не создавая переменную для использования его содержимого.
import (
	_ "building-service/docs"              // Импортирует пакет документации (нужен для работы Swaggo)
	"building-service/internal/handlers"   // Импортирует пакет обработчиков
	"building-service/internal/repository" // Импортирует пакет репозитория
	"building-service/pkg/database"        // Импортирует пакет для работы с базой данных
	"log"                                  // Импортирует стандартный пакет для логирования

	"github.com/gin-gonic/gin"                 // Импортирует библиотеку Gin для обработки HTTP-запросов
	_ "github.com/lib/pq"                      // Импортирует драйвер PostgreSQL (нужен для работы с базой данных)
	"github.com/spf13/viper"                   // Импортирует библиотеку Viper для работы с конфигурацией
	swaggerFiles "github.com/swaggo/files"     // Импортирует файлы Swagger для генерации документации
	ginSwagger "github.com/swaggo/gin-swagger" // Импортирует middleware для интеграции Swagger с Gin
)

// @title Building Service API           // Аннотация для указания заголовка API-документации
// @version 1.0                          // Аннотация для указания версии API
// @description Service for managing buildings // Аннотация для указания описания API
// @host localhost:8080                  // Аннотация для указания хоста API
// @BasePath /api/v1                     // Аннотация для указания базового пути API
func main() {
	// Загрузка конфигурации
	viper.SetConfigName("config")   // Устанавливает имя конфигурационного файла
	viper.SetConfigType("yaml")     // Устанавливает тип конфигурационного файла
	viper.AddConfigPath("./config") // Указывает путь к директории с конфигурационным файлом

	if err := viper.ReadInConfig(); err != nil { // Читает конфигурационный файл и проверяет на ошибки
		log.Fatalf("Error reading config file: %s", err) // Логирует и завершает программу в случае ошибки
	}

	// Подключение к базе данных
	db, err := database.NewPostgresDB(database.Config{ // Создает новое подключение к базе данных PostgreSQL
		Host:     viper.GetString("database.host"),     // Получает значение хоста из конфигурационного файла
		Port:     viper.GetInt("database.port"),        // Получает значение порта из конфигурационного файла
		User:     viper.GetString("database.user"),     // Получает значение пользователя из конфигурационного файла
		Password: viper.GetString("database.password"), // Получает значение пароля из конфигурационного файла
		DBName:   viper.GetString("database.dbname"),   // Получает значение имени базы данных из конфигурационного файла
	})
	if err != nil { // Проверяет на ошибки при подключении к базе данных
		log.Fatalf("Failed to initialize db: %s", err.Error()) // Логирует и завершает программу в случае ошибки
	}
	defer db.Close() // Откладывает закрытие подключения к базе данных

	// Настройка маршрутов
	router := gin.Default()        // Создает новый роутер Gin с настройками по умолчанию
	api := router.Group("/api/v1") // Создает группу маршрутов с базовым путём "/api/v1"
	{
		buildingRepo := repository.NewBuildingRepository(db)         // Создает новый репозиторий строений
		buildingHandler := handlers.NewBuildingHandler(buildingRepo) // Создает новый обработчик строений

		api.POST("/buildings", buildingHandler.Create) // Регистрирует маршрут для создания строений (POST /buildings)
		api.GET("/buildings", buildingHandler.List)    // Регистрирует маршрут для получения списка строений (GET /buildings)

		// Добавление маршрута для эндпоинта /docs
		api.GET("/docs", func(c *gin.Context) { // Регистрирует маршрут для документации API
			c.Redirect(302, "/swagger/doc.json") // Перенаправляет запрос на URL документации Swagger
		})

	}

	// Настройка Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // Регистрирует маршрут для Swagger-документации

	// Запуск сервера
	if err := router.Run(":8080"); err != nil { // Запускает HTTP-сервер на порту 8080
		log.Fatalf("Failed to run server: %s", err) // Логирует и завершает программу в случае ошибки
	}
}
