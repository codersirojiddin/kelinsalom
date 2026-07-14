package main

import (
	"fmt"
	handler "kelinsalom/api" // Vercel handlerlarini import qilamiz
	"log"
	"net/http"
)

func main() {
	// 1. Static resurslarni ulash (CSS va JS fayllar uchun)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. Vercel handlerlarini tegishli URL'larga bog'laymiz
	http.HandleFunc("/ads.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("google.com, pub-4819817034021416, DIRECT, f08c47fec0942fa0\n"))
	})
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/poem/", handler.PoemHandler)
	http.HandleFunc("/submit", handler.SubmitHandler)

	// 3. Serverni 8080-portda yoqamiz
	port := "8080"
	fmt.Printf("🚀 KelinSalom loyihasi localhost:%s da yonmoqda...\n", port)
	fmt.Printf("👉 Brauzerda oching: http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
