package main

import (
	"logTest/calc"
	"logTest/db"
	"logTest/logger"
)

func main() {
	// Логгер реализует централизованное логирование,
	// если мы забудем реализовать логирование для какой то библиотеки - ошибка компиляции
	logger := logger.NewLogrusLogger()

	// Подключаемся к бд (фейк в данном случаи)
	db := db.New(logger, "localhost")

	// Что-то считаем
	lib1 := calc.New(logger, db)
	lib1.Sum(10, 12)
}

