package database

import (
	"database/sql" // Import package sql untuk koneksi ke database

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB { // Fungsi untuk inisiasi koneksi ke database mySQL XAMPP
	dsn := "root@tcp(localhost:3306)/crud-employee"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping() // Mengecek koneksi ke database
	if err != nil {
		panic(err)
	}
	return db
}
