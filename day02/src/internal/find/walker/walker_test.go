package walker

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"day02/internal/find/flags"
	"github.com/stretchr/testify/assert"
)

// setupTestEnvironment создает временную файловую структуру для тестов.
func setupTestEnvironment() (string, error) {
	tmpDir, err := os.MkdirTemp("", "testenv")
	if err != nil {
		return "", err
	}

	// Создание тестовой файловой структуры
	os.Mkdir(filepath.Join(tmpDir, "dir1"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "dir1", "file1.txt"), []byte("test file 1"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "dir1", "file2.go"), []byte("test file 2"), 0644)
	os.Symlink(filepath.Join(tmpDir, "dir1", "file1.txt"), filepath.Join(tmpDir, "dir1", "symlink1"))
	os.Symlink(filepath.Join(tmpDir, "nonexistent"), filepath.Join(tmpDir, "dir1", "broken_symlink"))

	return tmpDir, nil
}

// teardownTestEnvironment удаляет временную файловую структуру после тестов.
func teardownTestEnvironment(tmpDir string) {
	os.RemoveAll(tmpDir)
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	buf.ReadFrom(r)
	return buf.String()
}

func TestWalkFileSystem_FindAllFiles(t *testing.T) {
	tmpDir, err := setupTestEnvironment()
	assert.NoError(t, err)
	defer teardownTestEnvironment(tmpDir)

	f := &flags.Flags{
		Path:      tmpDir,
		ShowFiles: true,
	}

	output := captureOutput(func() {
		err := WalkFileSystem(f)
		assert.NoError(t, err)
	})

	expectedOutput := []string{
		filepath.Join(tmpDir, "dir1/file1.txt"),
		filepath.Join(tmpDir, "dir1/file2.go"),
	}
	for _, expected := range expectedOutput {
		assert.Contains(t, output, expected)
	}
}

func TestWalkFileSystem_FindAllDirectories(t *testing.T) {
	tmpDir, err := setupTestEnvironment()
	assert.NoError(t, err)
	defer teardownTestEnvironment(tmpDir)

	f := &flags.Flags{
		Path:     tmpDir,
		ShowDirs: true,
	}

	output := captureOutput(func() {
		err := WalkFileSystem(f)
		assert.NoError(t, err)
	})

	expectedOutput := []string{
		filepath.Join(tmpDir, "dir1"),
	}
	for _, expected := range expectedOutput {
		assert.Contains(t, output, expected)
	}
}
