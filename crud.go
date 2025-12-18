package main

import (
	"net/http"

	"github.com/Dryluigi/crud-employee-go/database"
	"github.com/Dryluigi/crud-employee-go/routes" // Import dari github untuk mengarahkan routes dan database
)

func main() {
	db := database.InitDatabase()        // Membuat koneksi ke database mySQL XAMPP
	server := http.NewServeMux()         // Membuat server baru yaitu localhost
	routes.MapRoutes(server, db)         // Mengarahkan routes ke server dan database
	http.ListenAndServe(":8080", server) // Memulai inisiasi pada server di port 8080
}
