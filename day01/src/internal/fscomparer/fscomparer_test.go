package fscomparer

import (
	"bytes"
	//"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// captureOutput захватывает вывод функции f
func captureOutput(f func()) string {
	var buf bytes.Buffer
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout
	io.Copy(&buf, r)

	return buf.String()
}

func createTempFile(t *testing.T, content string) string {
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Не удалось создать временный файл: %v", err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("Не удалось записать в файл: %v", err)
	}
	file.Close()
	return file.Name()
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected map[string]struct{}
	}{
		{
			name:    "Three lines",
			content: "file1\nfile2\nfile3\n",
			expected: map[string]struct{}{
				"file1": {},
				"file2": {},
				"file3": {},
			},
		},
		{
			name:     "Empty file",
			content:  "",
			expected: map[string]struct{}{},
		},
		{
			name:    "Single line",
			content: "singleline\n",
			expected: map[string]struct{}{
				"singleline": {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempFile := createTempFile(t, tt.content)
			defer os.Remove(tempFile)

			result, err := ReadFile(tempFile)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCompareFS(t *testing.T) {
	tests := []struct {
		name           string
		oldSet         map[string]struct{}
		newSet         map[string]struct{}
		expectedOutput string
	}{
		{
			name: "Added and removed files",
			oldSet: map[string]struct{}{
				"file1": {},
				"file2": {},
				"file3": {},
			},
			newSet: map[string]struct{}{
				"file2": {},
				"file3": {},
				"file4": {},
			},
			expectedOutput: "ADDED file4\nREMOVED file1\n",
		},
		{
			name: "No changes",
			oldSet: map[string]struct{}{
				"file1": {},
			},
			newSet: map[string]struct{}{
				"file1": {},
			},
			expectedOutput: "",
		},
		{
			name: "All files removed",
			oldSet: map[string]struct{}{
				"file1": {},
				"file2": {},
			},
			newSet:         map[string]struct{}{},
			expectedOutput: "REMOVED file1\nREMOVED file2\n",
		},
		{
			name:   "All files added",
			oldSet: map[string]struct{}{},
			newSet: map[string]struct{}{
				"file1": {},
				"file2": {},
			},
			expectedOutput: "ADDED file1\nADDED file2\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				CompareFS(tt.oldSet, tt.newSet)
			})
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
