package main

import (
	"fmt"
	"log"
	"net/http"
	handler "kelinsalom/api" // Vercel handlerlarini import qilamiz
)

func main() {
	// 1. Static resurslarni ulash (CSS va JS fayllar uchun)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. Vercel handlerlarini tegishli URL'larga bog'laymiz
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