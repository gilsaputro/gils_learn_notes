package main

import (
	"log"
	"net/http"
)

// CORSMiddleware menangani CORS untuk setiap permintaan.
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Membuat router
	mux := http.NewServeMux()

	// Menggunakan middleware CORS untuk setiap permintaan
	handler := CORSMiddleware(mux)

	// Menambahkan handler untuk rute "/index"
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	// Menjalankan server
	log.Println("Starting app at :9000")
	http.ListenAndServe(":9000", handler)
}
