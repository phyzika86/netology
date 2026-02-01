// Пакет main демонстрирует работу с файлами.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	_ = ReadProcessWrite("error_input_file_path.txt", "error_output_file_path.txt", Process)
	_ = ReadProcessWrite("input.txt", "output.txt", Process)
}

// CreateFile создает новый файл по указанному пути.
// Если файл уже существует, он будет перезаписан.
func CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании файла %s: %w", path, err)
	}
	fmt.Printf("Дескриптор созданного файла %s: %d\n", path, file.Fd())
	return file, nil
}

// ReadProcessWrite читает файл построчно, обрабатывает каждую строку
// и записывает результат в новый файл.
func ReadProcessWrite(inputPath string, outputPath string, process func(string) (string, error)) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла %s: %w", inputPath, err)
	}
	defer file.Close()

	fileOut, err := CreateFile(outputPath)
	if err != nil {
		return fmt.Errorf("ошибка создания файла для записи %s: %w", outputPath, err)
	}
	defer fileOut.Close()

	writer := bufio.NewWriter(fileOut)
	scanner := bufio.NewScanner(file)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		newLine, _ := process(line)
		fmt.Println("Обработано:", line)
		_, err := writer.WriteString(newLine + "\n")
		if err != nil {
			log.Println("Ошибка при записи строки в файл:", err)
		}
	}

	if err = scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла %s: %w", inputPath, err)
	}

	log.Printf("Успешно: %s → %s\n", inputPath, outputPath)
	return nil
}

// Process преобразует строку в верхний регистр.
func Process(str string) (string, error) {
	return strings.ToUpper(str), nil
}
