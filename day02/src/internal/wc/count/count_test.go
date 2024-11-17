package count

import (
	"os"
	"testing"

	"day02/internal/wc/flags"
	"github.com/stretchr/testify/assert"
)

func TestCountLines(t *testing.T) {
	content := "line1\nline2\nline3\n"
	file := createTempFile(t, content)
	defer os.Remove(file.Name())

	count, err := countLines(file)
	assert.NoError(t, err)
	assert.Equal(t, 3, count)
}

func TestCountChars(t *testing.T) {
	content := "hello world"
	file := createTempFile(t, content)
	defer os.Remove(file.Name())

	count, err := countChars(file)
	assert.NoError(t, err)
	assert.Equal(t, 11, count)
}

func TestCountWords(t *testing.T) {
	content := "hello world"
	file := createTempFile(t, content)
	defer os.Remove(file.Name())

	count, err := countWords(file)
	assert.NoError(t, err)
	assert.Equal(t, 2, count)
}

func TestProcessFile_CountLines(t *testing.T) {
	content := "line1\nline2\nline3\n"
	file := createTempFile(t, content)
	defer os.Remove(file.Name())

	f := &flags.Flags{CountLines: true}
	result, err := processFile(file.Name(), f)
	assert.NoError(t, err)
	assert.Equal(t, 3, result.Count)
}

func TestProcessFile_CountChars(t *testing.T) {
	content := "hello world"
	file := createTempFile(t, content)
	defer os.Remove(file.Name())

	f := &flags.Flags{CountChars: true}
	result, err := processFile(file.Name(), f)
	assert.NoError(t, err)
	assert.Equal(t, 11, result.Count)
}

func TestProcessFile_CountWords(t *testing.T) {
	content := "hello world"
	file := createTempFile(t, content)
	defer os.Remove(file.Name())

	f := &flags.Flags{CountWords: true}
	result, err := processFile(file.Name(), f)
	assert.NoError(t, err)
	assert.Equal(t, 2, result.Count)
}

func TestProcessFiles(t *testing.T) {
	content1 := "hello world"
	content2 := "line1\nline2\n"
	file1 := createTempFile(t, content1)
	file2 := createTempFile(t, content2)
	defer os.Remove(file1.Name())
	defer os.Remove(file2.Name())

	filePaths := []string{file1.Name(), file2.Name()}
	f := &flags.Flags{CountWords: true}
	results, err := ProcessFiles(filePaths, f)
	assert.NoError(t, err)
	assert.Equal(t, 2, results[0].Count)
	assert.Equal(t, 2, results[1].Count)
}

// createTempFile создает временный файл с указанным содержимым и возвращает указатель на файл.
func createTempFile(t *testing.T, content string) *os.File {
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		t.Fatalf("Failed to seek to start of temp file: %v", err)
	}

	return file
}
