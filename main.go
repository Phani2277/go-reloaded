package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"go_reloaded/text_processing"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Использование: go run main.go input.txt output.txt")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Чтение входного файла
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла %s: %v\n", inputFile, err)
		os.Exit(1)
	}

	// Применение модификаций
	modifiedText := text_processing.ProcessText(string(content))

	// Запись в выходной файл
	err = ioutil.WriteFile(outputFile, []byte(modifiedText), 0644)
	if err != nil {
		fmt.Printf("Ошибка при записи в файл %s: %v\n", outputFile, err)
		os.Exit(1)
	}

	fmt.Printf("Файл успешно обработан и сохранен в %s\n", outputFile)
}
