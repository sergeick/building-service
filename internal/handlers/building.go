// package handlers

// import (
// 	"building-service/internal/models"
// 	"building-service/internal/repository"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type BuildingHandler struct {
// 	repo *repository.BuildingRepository
// }

// func NewBuildingHandler(repo *repository.BuildingRepository) *BuildingHandler {
// 	return &BuildingHandler{repo: repo}
// }

// // @Summary Create a new building
// // @Description Create a new building
// // @Tags buildings
// // @Accept json
// // @Produce json
// // @Param building body models.Building true "Building object"
// // @Success 201 {object} models.Building
// // @Failure 400 {object} models.ErrorResponse
// // @Failure 500 {object} models.ErrorResponse
// // @Router /buildings [post]
// func (h *BuildingHandler) Create(c *gin.Context) {
// 	var building models.Building
// 	if err := c.ShouldBindJSON(&building); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := h.repo.Create(&building); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, building)
// }

// // @Summary List buildings
// // @Description Get list of buildings
// // @Tags buildings
// // @Accept json
// // @Produce json
// // @Param city query string false "City"
// // @Param year query int false "Year built"
// // @Param floors query int false "Floors count"
// // @Success 200 {array} models.Building
// // @Failure 500 {object} models.ErrorResponse
// // @Router /buildings [get]
// func (h *BuildingHandler) List(c *gin.Context) {
// 	city := c.Query("city")
// 	year := c.DefaultQuery("year", "")
// 	floors := c.DefaultQuery("floors", "")
// 	buildings, err := h.repo.List(city, year, floors)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, buildings)
// }

package handlers // Пакет для обработчиков запросов

import (
	"building-service/internal/models"     // Импортирует пакет моделей
	"building-service/internal/repository" // Импортирует пакет репозитория
	"net/http"                             // Импортирует стандартный пакет для обработки HTTP-запросов

	"github.com/gin-gonic/gin" // Импортирует библиотеку Gin для обработки HTTP-запросов
)

type BuildingHandler struct { // Определяет структуру обработчика строений
	repo *repository.BuildingRepository // Поле для хранения репозитория строений
}

func NewBuildingHandler(repo *repository.BuildingRepository) *BuildingHandler { // Функция для создания нового обработчика строений
	return &BuildingHandler{repo: repo} // Возвращает новый экземпляр BuildingHandler с инициализированным репозиторием
}

// @Summary Create a new building       // Аннотация для генерации краткого описания эндпоинта
// @Description Create a new building   // Аннотация для генерации детального описания эндпоинта
// @Tags buildings                      // Аннотация для группировки эндпоинтов по тегу "buildings"
// @Accept json                         // Аннотация для указания формата входных данных (JSON)
// @Produce json                        // Аннотация для указания формата выходных данных (JSON)
// @Param building body models.Building true "Building object" // Аннотация для описания параметра запроса (тело запроса должно содержать объект модели Building)
// @Success 201 {object} models.Building // Аннотация для указания успешного ответа (201 Created с объектом модели Building)
// @Failure 400 {object} models.ErrorResponse // Аннотация для указания ответа в случае ошибки (400 Bad Request с объектом модели ErrorResponse)
// @Failure 500 {object} models.ErrorResponse // Аннотация для указания ответа в случае ошибки (500 Internal Server Error с объектом модели ErrorResponse)
// @Router /buildings [post]            // Аннотация для указания пути эндпоинта и HTTP-метода (POST)
func (h *BuildingHandler) Create(c *gin.Context) { // Функция для обработки создания нового строения
	var building models.Building                        // Объявляет переменную для хранения данных строения
	if err := c.ShouldBindJSON(&building); err != nil { // Привязывает данные JSON из запроса к переменной building, проверяет на ошибки
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Возвращает ответ с кодом 400 и сообщением об ошибке
		return
	}
	if err := h.repo.Create(&building); err != nil { // Сохраняет данные строения в базе данных через репозиторий, проверяет на ошибки
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращает ответ с кодом 500 и сообщением об ошибке
		return
	}
	c.JSON(http.StatusCreated, building) // Возвращает успешный ответ с кодом 201 и данными созданного строения
}

// @Summary List buildings              // Аннотация для генерации краткого описания эндпоинта
// @Description Get list of buildings   // Аннотация для генерации детального описания эндпоинта
// @Tags buildings                      // Аннотация для группировки эндпоинтов по тегу "buildings"
// @Accept json                         // Аннотация для указания формата входных данных (JSON)
// @Produce json                        // Аннотация для указания формата выходных данных (JSON)
// @Param city query string false "City" // Аннотация для описания параметра запроса (фильтрация по городу)
// @Param year query int false "Year built" // Аннотация для описания параметра запроса (фильтрация по году)
// @Param floors query int false "Floors count" // Аннотация для описания параметра запроса (фильтрация по количеству этажей)
// @Success 200 {array} models.Building  // Аннотация для указания успешного ответа (200 OK с массивом объектов модели Building)
// @Failure 500 {object} models.ErrorResponse // Аннотация для указания ответа в случае ошибки (500 Internal Server Error с объектом модели ErrorResponse)
// @Router /buildings [get]             // Аннотация для указания пути эндпоинта и HTTP-метода (GET)
func (h *BuildingHandler) List(c *gin.Context) { // Функция для обработки получения списка строений
	city := c.Query("city")                           // Получает значение параметра фильтрации по городу из запроса
	year := c.DefaultQuery("year", "")                // Получает значение параметра фильтрации по году из запроса, если параметр отсутствует, возвращает пустую строку
	floors := c.DefaultQuery("floors", "")            // Получает значение параметра фильтрации по количеству этажей из запроса, если параметр отсутствует, возвращает пустую строку
	buildings, err := h.repo.List(city, year, floors) // Получает список строений из базы данных через репозиторий с учетом параметров фильтрации
	if err != nil {                                   // Проверяет на ошибки при получении данных
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращает ответ с кодом 500 и сообщением об ошибке
		return
	}
	c.JSON(http.StatusOK, buildings) // Возвращает успешный ответ с кодом 200 и списком строений
}
