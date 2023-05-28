package repositories

import (
	"balancer/db"
	"database/sql"
)

var DB *sql.DB = db.Init()
