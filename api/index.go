package handler

import (
 "embed"
 "fmt"
 "html/template"
 "kelinsalom/pkg/data"
 "net/http"
)

type IndexData struct {
 Poems     []data.Poem
 Region    string
 Query     string
 PageTitle string
 MetaDesc  string
 URLPath   string
}

//go:embed templates/*.html
var templateFS embed.FS

func IndexHandler(w http.ResponseWriter, r *http.Request) {
 // 1. SITEMAP.XML
 if r.URL.Query().Get("sitemap") == "true" {
  w.Header().Set("Content-Type", "application/xml; charset=utf-8")
  fmt.Fprint(w, "<?xml version=\"1.0\" encoding=\"UTF-8\"?><urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"><url><loc>https://kelinsalom.uz/</loc><priority>1.0</priority></url>")
  
  allPoems := data.LoadPoems()
  for _, p := range allPoems {
   fmt.Fprintf(w, "<url><loc>https://kelinsalom.uz/poem/%d</loc><priority>0.8</priority></url>", p.ID)
  }
  fmt.Fprint(w, "</urlset>")
  return
 }

 // 2. ODDIY SO'ROV
 if r.URL.Path != "/" {
  http.NotFound(w, r)
  return
 }

 region := r.URL.Query().Get("region")
 query := r.URL.Query().Get("q")
 filteredPoems := data.FilterAndSearch(region, query)

 tmpl, err := template.ParseFS(templateFS, "templates/base.html", "templates/index.html")
 if err != nil {
  http.Error(w, "Shablon xatosi: "+err.Error(), http.StatusInternalServerError)
  return
 }

 pageTitle := "Kelin Salom Matnlari To'plami — To'liq va Sara She'rlar"
 metaDesc := "O'zbekona to'ylar uchun eng sara, to'liq va an'anaviy kelin salom matnlari. Viloyatlar bo'yicha qulay qidiruv va nusxalash."
 
 if region != "" {
  pageTitle = region + "cha Kelin Salom Matnlari — To'liq To'plam"
  metaDesc = region + " viloyatiga xos bo'lgan eng chiroyli va an'anaviy kelin salom she'rlari matni."
 }

 dataToSend := IndexData{
  Poems:     filteredPoems,
  Region:    region,
  Query:     query,
  PageTitle: pageTitle,
  MetaDesc:  metaDesc,
  URLPath:   "/",
 }

 w.Header().Set("Content-Type", "text/html; charset=utf-8")
 _ = tmpl.ExecuteTemplate(w, "base.html", dataToSend)
}