package archiver

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"day02/internal/rotate/flags"
	"github.com/stretchr/testify/assert"
)

func TestArchiveFile(t *testing.T) {
	// Создаем временный файл
	tmpFile, err := ioutil.TempFile("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Записываем данные в временный файл
	_, err = tmpFile.WriteString("Hello, World!")
	assert.NoError(t, err)

	// Закрываем файл, чтобы использовать его для архивации
	err = tmpFile.Close()
	assert.NoError(t, err)

	// Создаем временную директорию для архива
	tmpDir, err := ioutil.TempDir("", "archive")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Архивируем файл
	err = archiveFile(tmpFile.Name(), tmpDir)
	assert.NoError(t, err)

	// Проверяем, что архив создан
	files, err := ioutil.ReadDir(tmpDir)
	assert.NoError(t, err)
	assert.Len(t, files, 1)

	// Проверяем содержимое архива
	archivePath := filepath.Join(tmpDir, files[0].Name())
	archiveFile, err := os.Open(archivePath)
	assert.NoError(t, err)
	defer archiveFile.Close()

	gzr, err := gzip.NewReader(archiveFile)
	assert.NoError(t, err)
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	header, err := tr.Next()
	assert.NoError(t, err)
	assert.Equal(t, filepath.Base(tmpFile.Name()), header.Name)

	var buf bytes.Buffer
	_, err = io.Copy(&buf, tr)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!", buf.String())
}

func TestProcessFiles(t *testing.T) {
	// Создаем временные файлы
	tmpFile1, err := ioutil.TempFile("", "testfile1")
	assert.NoError(t, err)
	defer os.Remove(tmpFile1.Name())
	_, err = tmpFile1.WriteString("Hello, File 1!")
	assert.NoError(t, err)
	err = tmpFile1.Close()
	assert.NoError(t, err)

	tmpFile2, err := ioutil.TempFile("", "testfile2")
	assert.NoError(t, err)
	defer os.Remove(tmpFile2.Name())
	_, err = tmpFile2.WriteString("Hello, File 2!")
	assert.NoError(t, err)
	err = tmpFile2.Close()
	assert.NoError(t, err)

	// Создаем временную директорию для архива
	tmpDir, err := ioutil.TempDir("", "archive")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Устанавливаем флаги
	flags := &flags.Flags{
		FilePaths:  []string{tmpFile1.Name(), tmpFile2.Name()},
		ArchiveDir: tmpDir,
	}

	// Архивируем файлы
	err = ProcessFiles(flags)
	assert.NoError(t, err)

	// Проверяем, что архивы созданы
	files, err := ioutil.ReadDir(tmpDir)
	assert.NoError(t, err)
	assert.Len(t, files, 2)

	// Проверяем содержимое первого архива
	archivePath1 := filepath.Join(tmpDir, files[0].Name())
	archiveFile1, err := os.Open(archivePath1)
	assert.NoError(t, err)
	defer archiveFile1.Close()

	gzr1, err := gzip.NewReader(archiveFile1)
	assert.NoError(t, err)
	defer gzr1.Close()

	tr1 := tar.NewReader(gzr1)
	header1, err := tr1.Next()
	assert.NoError(t, err)
	assert.Equal(t, filepath.Base(tmpFile1.Name()), header1.Name)

	var buf1 bytes.Buffer
	_, err = io.Copy(&buf1, tr1)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, File 1!", buf1.String())

	// Проверяем содержимое второго архива
	archivePath2 := filepath.Join(tmpDir, files[1].Name())
	archiveFile2, err := os.Open(archivePath2)
	assert.NoError(t, err)
	defer archiveFile2.Close()

	gzr2, err := gzip.NewReader(archiveFile2)
	assert.NoError(t, err)
	defer gzr2.Close()

	tr2 := tar.NewReader(gzr2)
	header2, err := tr2.Next()
	assert.NoError(t, err)
	assert.Equal(t, filepath.Base(tmpFile2.Name()), header2.Name)

	var buf2 bytes.Buffer
	_, err = io.Copy(&buf2, tr2)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, File 2!", buf2.String())
}
