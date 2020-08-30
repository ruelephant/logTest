package db

// Интерфейс логгера содержит только методы логирования которые использует данная библиотека
type Logger interface {
	Connect(host string)
}