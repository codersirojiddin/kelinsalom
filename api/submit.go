package handler

import (
 "fmt"
 "kelinsalom/pkg/telegram"
 "net/http"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
 if r.Method != http.MethodPost {
  http.Error(w, "Metodga ruxsat yo'q", http.StatusMethodNotAllowed)
  return
 }

 author := r.FormValue("author")
 text := r.FormValue("text")

 if author == "" || text == "" {
  http.Error(w, "Iltimos, hamma maydonlarni to'ldiring", http.StatusBadRequest)
  return
 }

 err := telegram.SendToAdmin(author, text)
 if err != nil {
  fmt.Println("Telegram jo'natishda xato:", err)
 }

 w.Header().Set("Content-Type", "text/html; charset=utf-8")
 fmt.Fprint(w, "<div style=\"text-align: center; font-family: Tahoma, Arial; padding: 50px;\"><h2 style=\"color: #3b5998;\">Katta rahmat!</h2><p>Siz yuborgan kelin salom matni administratorga yetkazildi. Tez orada tekshirib, saytga qo'shamiz.</p><p><i>Hozir bosh sahifaga qaytasiz...</i></p><script>setTimeout(function(){ window.location.href = \"/\"; }, 3000);</script></div>")
}