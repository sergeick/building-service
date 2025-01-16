// package models

// type ErrorResponse struct {
// 	Error   string `json:"error"`
// 	Code    int    `json:"code"`
// 	Message string `json:"message,omitempty"`
// }

//	type SuccessResponse struct {
//		Data    interface{} `json:"data"`
//		Message string      `json:"message,omitempty"`
//	}
package models // Пакет для моделей данных

type ErrorResponse struct { // Определяет структуру данных для ответа при ошибке
	Error   string `json:"error"`             // Сообщение об ошибке, будет представлено как "error" в JSON
	Code    int    `json:"code"`              // Код ошибки, будет представлен как "code" в JSON
	Message string `json:"message,omitempty"` // Дополнительное сообщение об ошибке, будет представлено как "message" в JSON, это поле является необязательным (omitempty)
}

type SuccessResponse struct { // Определяет структуру данных для успешного ответа
	Data    interface{} `json:"data"`              // Данные успешного ответа, будет представлено как "data" в JSON, тип данных может быть любым (interface{})
	Message string      `json:"message,omitempty"` // Дополнительное сообщение, будет представлено как "message" в JSON, это поле является необязательным (omitempty)
}
