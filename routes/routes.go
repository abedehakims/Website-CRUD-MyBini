package routes

import (
	"database/sql"
	"net/http"

	"github.com/Dryluigi/crud-employee-go/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.Controller())                                   // Setting server route untuk ke handler controller
	server.HandleFunc("/employee", controller.NewIndexEmployee(db))                   // Setting server route untuk ke halaman utama
	server.HandleFunc("/employee/create", controller.NewCreateEmployeeController(db)) // Setting server route untuk ke halaman create
	server.HandleFunc("/employee/update", controller.NewUpdateEmployeeController(db)) // Setting server route untuk ke halaman update
	server.HandleFunc("/employee/delete", controller.NewDeleteEmployeeController(db)) // Setting server route untuk ke halaman delete
}
