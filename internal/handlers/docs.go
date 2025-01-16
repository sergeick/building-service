// package handlers

// import (
// 	"github.com/gin-gonic/gin"
// 	swaggerFiles "github.com/swaggo/files"
// 	ginSwagger "github.com/swaggo/gin-swagger"
// )

// // @Summary Swagger API documentation
// // @Description Get Swagger API documentation
// // @Tags docs
// // @Produce html
// // @Router /docs [get]
//
//	func SwaggerHandler() gin.HandlerFunc {
//		return ginSwagger.WrapHandler(swaggerFiles.Handler)
//	}
package handlers // Пакет для обработчиков запросов

import (
	"github.com/gin-gonic/gin"                 // Импортирует библиотеку Gin для обработки HTTP-запросов
	swaggerFiles "github.com/swaggo/files"     // Импортирует файлы Swagger для генерации документации
	ginSwagger "github.com/swaggo/gin-swagger" // Импортирует middleware для интеграции Swagger с Gin
)

// @Summary Swagger API documentation       // Аннотация для генерации краткого описания эндпоинта
// @Description Get Swagger API documentation // Аннотация для генерации детального описания эндпоинта
// @Tags docs                               // Аннотация для группировки эндпоинтов по тегу "docs"
// @Produce html                            // Аннотация для указания формата выходных данных (HTML)
// @Router /docs [get]                      // Аннотация для указания пути эндпоинта и HTTP-метода (GET)
func SwaggerHandler() gin.HandlerFunc { // Функция для предоставления Swagger-документации
	return ginSwagger.WrapHandler(swaggerFiles.Handler) // Возвращает middleware для обработки запросов к документации Swagger
}
