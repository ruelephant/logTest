package calc

// Интерфейс логгера содержит только методы логирования которые использует данная библиотека
type Logger interface {
	Sum(a, b int)
}