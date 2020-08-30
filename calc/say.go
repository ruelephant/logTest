package calc

import (
	db "database/sql"
	"fmt"
)


type Calc struct {
	log Logger
}

func (ml *Calc) Sum(a, b int) {
	// Логируем что приложение сказало calc
	ml.log.Sum(a, b)

	fmt.Print("Hi all!")
}

func New(l Logger, db *db.DB) *Calc {
	return &Calc{
		log: l,
	}
}