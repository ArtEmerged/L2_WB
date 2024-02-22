package wget

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// ErrMissURL - ошибка при отсутствии URL
var (
	ErrMissURL = errors.New("wget: missing URL")
)

// Run - функция для запуска wget
func Run(args []string) error {
	var outputFile string
	fs := flag.NewFlagSet("wget", flag.ExitOnError)
	fs.StringVar(&outputFile, "o", "", "output file")
	fs.Parse(args)

	// Проверяем на наличие URL в аргументах
	if len(fs.Args()) == 0 {
		return ErrMissURL
	}

	url := fs.Args()[0]

	// Получение ответа от сервера по URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверяем статуса ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("wget: ERROR %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// Определение имени файла для сохранения
	if outputFile == "" {
		outputFile = generateFileName(url, resp.Header.Get("Content-Type"))
	}
	fmt.Println(outputFile)

	// Создание файла
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Копирование содержимого ответа в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// generateFileName - функция для генерации имени файла на основе URL и типа содержимого
func generateFileName(url, contentType string) string {
	var extension string

	// Получаем расширения файла из типа содержимого
	if contentType != "" {
		contentType = strings.Split(contentType, ";")[0]
		if typeFile := strings.Split(contentType, "/"); len(typeFile) == 2 {
			extension = "." + typeFile[1]
		}
	}

	// Удаление последнего слэша из URL
	url = strings.TrimSuffix(url, "/")
	nameSplit := strings.Split(url, "/")

	// Получаем имени файла из URL
	fileName := nameSplit[len(nameSplit)-1]

	// Добавление расширения, если оно не совпадает
	if !strings.HasSuffix(fileName, extension) {
		fileName += extension
	}

	return fileName
}
