package data

import (
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestLoadPoemsWorksFromDifferentWorkingDirectory(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	defer func() {
		_ = os.Chdir(cwd)
	}()

	tempDir := t.TempDir()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	Instance = nil
	once = sync.Once{}

	poems := LoadPoems()
	if len(poems) == 0 {
		t.Fatalf("expected poems to load from project data file, got %d", len(poems))
	}
}

func TestLoadPoemsFallsBackWhenDataFileIsMissing(t *testing.T) {
	poemsPath := filepath.Join("..", "..", "db", "poems.json")
	backupPath := poemsPath + ".bak"

	if _, err := os.Stat(poemsPath); err != nil {
		t.Fatalf("expected project data file to exist at %s: %v", poemsPath, err)
	}

	if err := os.Rename(poemsPath, backupPath); err != nil {
		t.Fatalf("rename data file: %v", err)
	}
	defer func() {
		_ = os.Rename(backupPath, poemsPath)
	}()

	Instance = nil
	once = sync.Once{}

	poems := LoadPoems()
	if len(poems) == 0 {
		t.Fatalf("expected poems to load from embedded fallback data, got %d", len(poems))
	}
}
