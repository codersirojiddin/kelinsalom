package data

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
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

func resolvePoemsPath() string {
	var candidates []string

	if cwd, err := os.Getwd(); err == nil {
		candidates = append(candidates, filepath.Join(cwd, "db", "poems.json"))
	}

	if _, file, _, ok := runtime.Caller(0); ok {
		baseDir := filepath.Dir(file)
		for dir := baseDir; ; dir = filepath.Dir(dir) {
			candidates = append(candidates, filepath.Join(dir, "db", "poems.json"))
			parent := filepath.Dir(dir)
			if parent == dir {
				break
			}
		}
	}

	if exe, err := os.Executable(); err == nil {
		candidates = append(candidates, filepath.Join(filepath.Dir(exe), "db", "poems.json"))
	}

	for _, candidate := range candidates {
		if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
			return candidate
		}
	}

	return ""
}

// LoadPoems JSON faylni RAMga yuklaydi
func LoadPoems() []Poem {
	once.Do(func() {
		path := resolvePoemsPath()
		if path == "" {
			Instance = []Poem{}
			return
		}

		file, err := os.ReadFile(path)
		if err != nil {
			Instance = []Poem{}
			return
		}
		if err := json.Unmarshal(file, &Instance); err != nil {
			Instance = []Poem{}
			return
		}
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
