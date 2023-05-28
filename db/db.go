package db

import (
	"balancer/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	host := config.Config("HOST")
	db, err := sql.Open("mysql", "root:@tcp("+host+":3306)/mysql")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query("CREATE DATABASE IF NOT EXISTS balancer;")
	if err != nil {
		panic(err)
	}

	db, err = sql.Open("mysql", "root:@tcp("+host+":3306)/balancer")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS `servers` (`id` text NOT NULL, `domain_id` text NOT NULL, `address` text NOT NULL, `label` text NOT NULL, `weight` int(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")
	if err != nil {
		panic(err)
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS `domains` ( `id` text NOT NULL, `label` text NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")
	if err != nil {
		panic(err)
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS `users` ( `id` text NOT NULL, `email` text NOT NULL, `hash` text NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")
	if err != nil {
		panic(err)
	}

	return db
}
