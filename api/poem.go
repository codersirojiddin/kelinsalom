package handler

import (
 "embed"
 "html/template"
 "kelinsalom/pkg/data"
 "net/http"
 "strconv"
 "strings"
)

type PoemData struct {
 Poem      data.Poem
 PageTitle string
 MetaDesc  string
 URLPath   string
}

//go:embed templates/*.html
var poemTemplateFS embed.FS

func PoemHandler(w http.ResponseWriter, r *http.Request) {
 idStr := r.URL.Query().Get("id")
 if idStr == "" {
  parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
  if len(parts) >= 2 {
   idStr = parts[1]
  }
 }

 id, err := strconv.Atoi(idStr)
 if err != nil {
  http.NotFound(w, r)
  return
 }

 allPoems := data.LoadPoems()
 var currentPoem data.Poem
 found := false

 for _, p := range allPoems {
  if p.ID == id {
   currentPoem = p
   found = true
   break
  }
 }

 if !found {
  http.NotFound(w, r)
  return
 }

 tmpl, err := template.ParseFS(poemTemplateFS, "templates/base.html", "templates/poem.html")
 if err != nil {
  http.Error(w, "Shablon xatosi: "+err.Error(), http.StatusInternalServerError)
  return
 }

 pageTitle := currentPoem.Title + " — KelinSalom.uz"
 
 runes := []rune(currentPoem.Content)
 length := len(runes)
 if length > 150 {
  length = 150
 }
 metaDesc := string(runes[0:length]) + "..."

 w.Header().Set("Content-Type", "text/html; charset=utf-8")
 _ = tmpl.ExecuteTemplate(w, "base.html", PoemData{
  Poem:      currentPoem,
  PageTitle: pageTitle,
  MetaDesc:  metaDesc,
  URLPath:   "/poem/" + idStr,
 })
}