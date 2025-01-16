
-- CREATE TABLE IF NOT EXISTS buildings (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(255) NOT NULL,
--     city VARCHAR(255) NOT NULL,
--     year_built INTEGER NOT NULL,
--     floors_count INTEGER NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

CREATE TABLE IF NOT EXISTS buildings (  -- Создаёт таблицу 'buildings', если она ещё не существует
    id SERIAL PRIMARY KEY,              -- Создаёт колонку 'id' с автогенерируемыми значениями и устанавливает её как первичный ключ
    name VARCHAR(255) NOT NULL,         -- Создаёт колонку 'name' с типом VARCHAR(255) и указывает, что это поле обязательно для заполнения
    city VARCHAR(255) NOT NULL,         -- Создаёт колонку 'city' с типом VARCHAR(255) и указывает, что это поле обязательно для заполнения
    year_built INTEGER NOT NULL,        -- Создаёт колонку 'year_built' с типом INTEGER и указывает, что это поле обязательно для заполнения
    floors_count INTEGER NOT NULL,      -- Создаёт колонку 'floors_count' с типом INTEGER и указывает, что это поле обязательно для заполнения
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Создаёт колонку 'created_at' с типом TIMESTAMP и устанавливает значение по умолчанию как текущее время
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   -- Создаёт колонку 'updated_at' с типом TIMESTAMP и устанавливает значение по умолчанию как текущее время
);
