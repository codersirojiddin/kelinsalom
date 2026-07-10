package data

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
)

type Poem struct {
	ID      int    `json:"id"`
	Region  string `json:"region"` // Faqat hudud filtri qoladi (Samarqandcha, Vodiycha...)
	Title   string `json:"title"`
	Content string `json:"content"`
	Views   int    `json:"views"`
}

var (
	Instance []Poem
	once     sync.Once
)

// LoadPoems JSON faylni RAMga yuklaydi
func LoadPoems() []Poem {
	once.Do(func() {
		file, err := os.ReadFile("db/poems.json")
		if err != nil {
			Instance = []Poem{}
			return
		}
		_ = json.Unmarshal(file, &Instance)
	})
	return Instance
}

// FilterAndSearch Endi faqat hudud va qidiruv so'zi bo'yicha ishlaydi
func FilterAndSearch(region, query string) []Poem {
	allPoems := LoadPoems()
	var result []Poem

	for _, p := range allPoems {
		// Viloyat filtri (chunki hududiy urf-odatlar baribir farq qiladi)
		if region != "" && p.Region != region {
			continue
		}
		// Matnli tezkor qidiruv
		if query != "" {
			q := strings.ToLower(query)
			if !strings.Contains(strings.ToLower(p.Title), q) && !strings.Contains(strings.ToLower(p.Content), q) {
				continue
			}
		}
		result = append(result, p)
	}
	return result
}