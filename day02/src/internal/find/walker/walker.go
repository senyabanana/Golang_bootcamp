package walker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"day02/internal/find/flags"
)

// WalkFileSystem рекурсивно обходит файловую систему начиная с указанного пути.
func WalkFileSystem(f *flags.Flags) error {
	return filepath.Walk(f.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Если ошибка связана с разрешениями, просто пропускаем этот файл/директорию
			if os.IsPermission(err) {
				return nil
			}
			return err
		}

		// Проверка, является ли текущий элемент символической ссылкой
		if info.Mode()&os.ModeSymlink != 0 {
			if f.ShowSymlinks {
				target, err := os.Readlink(path)
				if err != nil {
					// Если символическая ссылка сломана (не удается разрешить), выводим [broken]
					fmt.Printf("%s -> [broken]\n", path)
				} else {
					// Если символическая ссылка разрешена успешно, выводим путь назначения
					fmt.Printf("%s -> %s\n", path, target)
				}
			}
			return nil
		}

		// Проверка, является ли текущий элемент обычным файлом
		if f.ShowFiles && info.Mode().IsRegular() {
			// Если указано расширение и файл соответствует этому расширению, выводим его
			if f.Ext == "" || strings.HasSuffix(path, "."+f.Ext) {
				fmt.Println(path)
			}
		}

		// Проверка, является ли текущий элемент директорией
		if f.ShowDirs && info.IsDir() {
			fmt.Println(path)
		}

		return nil
	})
}
