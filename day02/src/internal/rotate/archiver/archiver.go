package archiver

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"day02/internal/rotate/flags"
)

// archiveFile создает архив для указанного файла.
func archiveFile(filePath, archiveDir string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Генерируем имя архива с использованием метки времени
	timestamp := info.ModTime().Unix()
	archivePath := fmt.Sprintf("%s_%d.tar.gz", filepath.Base(filePath), timestamp)

	// Если указана директория для архивации, добавляем ее к пути архива
	if archiveDir != "" {
		archivePath = filepath.Join(archiveDir, archivePath)
	}

	// Создаем файл архива
	archiveFile, err := os.Create(archivePath)
	if err != nil {
		return err
	}
	defer archiveFile.Close()

	// Инициализируем gzip писатель
	gw := gzip.NewWriter(archiveFile)
	defer gw.Close()

	// Инициализируем tar писатель
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Создаем заголовок для tar архива
	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}

	// Устанавливаем имя файла в заголовке
	header.Name = filepath.Base(filePath)

	// Записываем заголовок в tar архив
	if err := tw.WriteHeader(header); err != nil {
		return err
	}

	// Копируем содержимое файла в tar архив
	if _, err := io.Copy(tw, file); err != nil {
		return err
	}

	return nil
}

// ProcessFiles архивирует несколько файлов параллельно.
func ProcessFiles(flags *flags.Flags) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(flags.FilePaths))

	for _, filePath := range flags.FilePaths {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			if err := archiveFile(filePath, flags.ArchiveDir); err != nil {
				errCh <- err
			}
		}(filePath)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}
