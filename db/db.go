package db

import db "database/sql"

func New(l Logger, host string) *db.DB {
	// Логируем попытку соединения, передаем в логгер host
	l.Connect(host)

	return &db.DB{}
}
