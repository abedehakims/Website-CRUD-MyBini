package controller

import (
	"database/sql"
	"net/http" // Mengelola request dan response di HTTP
)

func NewDeleteEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) { // Handler menghapus data
		id := r.URL.Query().Get("id")
		_, err := db.Exec("DELETE FROM employee WHERE id = ?", id) // Query / perintah menghapus data berdasarkan id (data dari input user)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/employee", http.StatusMovedPermanently) // Setelah data dihapus, kembali ke halaman utama
	}
}
