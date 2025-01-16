// package models

// type Building struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name" validate:"required"`
// 	City        string `json:"city" validate:"required"`
// 	YearBuilt   int    `json:"year" validate:"required"`
// 	FloorsCount int    `json:"floors" validate:"required"`
// }

package models // Пакет для моделей данных

type Building struct { // Определяет структуру данных для строения
	ID          int    `json:"id"`                         // Уникальный идентификатор строения, будет представлен как "id" в JSON
	Name        string `json:"name" validate:"required"`   // Название строения, будет представлено как "name" в JSON, обязательно для заполнения
	City        string `json:"city" validate:"required"`   // Город, в котором находится строение, будет представлен как "city" в JSON, обязательно для заполнения
	YearBuilt   int    `json:"year" validate:"required"`   // Год постройки строения, будет представлен как "year" в JSON, обязательно для заполнения
	FloorsCount int    `json:"floors" validate:"required"` // Количество этажей в строении, будет представлено как "floors" в JSON, обязательно для заполнения
}
