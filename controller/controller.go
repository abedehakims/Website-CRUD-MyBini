package controller

import "net/http"

func Controller() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ini database Employee"))
	}
}

// Testing untuk menampilkan teks sederhana pada halaman utama server.
// Bisa juga digunakan untuk menangani permintaan di HTTP dan melanjutkan ke handler lalu ke server
