package data

import (
	"os"
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
