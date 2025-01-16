// package repository

// import (
// 	"building-service/internal/models"
// 	"database/sql"
// 	"strconv"
// )

// type BuildingRepository struct {
// 	db *sql.DB
// }

// func NewBuildingRepository(db *sql.DB) *BuildingRepository {
// 	return &BuildingRepository{db: db}
// }

// func (r *BuildingRepository) Create(building *models.Building) error {
// 	query := `
//         INSERT INTO buildings (name, city, year_built, floors_count)
//         VALUES ($1, $2, $3, $4)
//         RETURNING id`

// 	return r.db.QueryRow(query,
// 		building.Name,
// 		building.City,
// 		building.YearBuilt,
// 		building.FloorsCount,
// 	).Scan(&building.ID)
// }

// func (r *BuildingRepository) List(city, year, floors string) ([]models.Building, error) {
// 	query := `SELECT id, name, city, year_built, floors_count FROM buildings WHERE 1=1`
// 	params := []interface{}{}

// 	if city != "" {
// 		query += " AND city = $" + strconv.Itoa(len(params)+1)
// 		params = append(params, city)
// 	}
// 	if year != "" {
// 		query += " AND year_built = $" + strconv.Itoa(len(params)+1)
// 		params = append(params, year)
// 	}
// 	if floors != "" {
// 		query += " AND floors_count = $" + strconv.Itoa(len(params)+1)
// 		params = append(params, floors)
// 	}

// 	rows, err := r.db.Query(query, params...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var buildings []models.Building
// 	for rows.Next() {
// 		var b models.Building
// 		if err := rows.Scan(&b.ID, &b.Name, &b.City, &b.YearBuilt, &b.FloorsCount); err != nil {
// 			return nil, err
// 		}
// 		buildings = append(buildings, b)
// 	}
// 	return buildings, nil
// }

package repository // Пакет для репозитория

import (
	"building-service/internal/models" // Импортирует пакет моделей
	"database/sql"                     // Импортирует стандартный пакет для работы с SQL базами данных
	"strconv"                          // Импортирует пакет strconv для конвертации строк
)

type BuildingRepository struct { // Определяет структуру репозитория строений
	db *sql.DB // Поле для хранения объекта подключения к базе данных
}

func NewBuildingRepository(db *sql.DB) *BuildingRepository { // Функция для создания нового репозитория строений
	return &BuildingRepository{db: db} // Возвращает новый экземпляр BuildingRepository с инициализированным подключением к базе данных
}

func (r *BuildingRepository) Create(building *models.Building) error { // Функция для создания нового строения в базе данных
	query := `
        INSERT INTO buildings (name, city, year_built, floors_count)
        VALUES ($1, $2, $3, $4)
        RETURNING id` // SQL-запрос для вставки нового строения и возврата его ID

	return r.db.QueryRow(query, // Выполняет SQL-запрос, передавая параметры строения
		building.Name,        // Имя строения
		building.City,        // Город строения
		building.YearBuilt,   // Год постройки строения
		building.FloorsCount, // Количество этажей строения
	).Scan(&building.ID) // Сканирует возвращенный ID строения и сохраняет его в модели
}

func (r *BuildingRepository) List(city, year, floors string) ([]models.Building, error) { // Функция для получения списка строений с фильтрацией
	query := `SELECT id, name, city, year_built, floors_count FROM buildings WHERE 1=1` // SQL-запрос для выбора строений с начальным условием
	params := []interface{}{}                                                           // Слайс для хранения параметров запроса

	if city != "" { // Если параметр города не пуст
		query += " AND city = $" + strconv.Itoa(len(params)+1) // Добавляет условие для фильтрации по городу
		params = append(params, city)                          // Добавляет значение города в параметры запроса
	}
	if year != "" { // Если параметр года не пуст
		query += " AND year_built = $" + strconv.Itoa(len(params)+1) // Добавляет условие для фильтрации по году
		params = append(params, year)                                // Добавляет значение года в параметры запроса
	}
	if floors != "" { // Если параметр количества этажей не пуст
		query += " AND floors_count = $" + strconv.Itoa(len(params)+1) // Добавляет условие для фильтрации по количеству этажей
		params = append(params, floors)                                // Добавляет значение количества этажей в параметры запроса
	}

	rows, err := r.db.Query(query, params...) // Выполняет SQL-запрос с параметрами
	if err != nil {                           // Проверяет на ошибки при выполнении запроса
		return nil, err // Возвращает ошибку, если есть
	}
	defer rows.Close() // Откладывает закрытие результирующего набора строк

	var buildings []models.Building // Создает слайс для хранения списка строений
	for rows.Next() {               // Итерация по всем строкам результирующего набора
		var b models.Building                                                                    // Создает переменную для хранения данных строения
		if err := rows.Scan(&b.ID, &b.Name, &b.City, &b.YearBuilt, &b.FloorsCount); err != nil { // Сканирует данные строки в переменную строения
			return nil, err // Возвращает ошибку, если есть
		}
		buildings = append(buildings, b) // Добавляет строение в список
	}
	return buildings, nil // Возвращает список строений и nil, если ошибок нет
}
