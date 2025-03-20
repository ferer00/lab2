package main

import (
	"flag"
	"fmt"
	"io"
	"lab2/pkg/handler" // Імпорт пакету handler
	"os"
	"strings"
)

func main() {
	expressionFlag := flag.String("e", "", "Вираз для обробки")
	fileFlag := flag.String("f", "", "Файл з виразом")
	outputFlag := flag.String("o", "", "Файл для результату")
	flag.Parse()

	if *expressionFlag != "" && *fileFlag != "" {
		fmt.Fprintln(os.Stderr, "Помилка: не можна використовувати одночасно -e та -f")
		os.Exit(1)
	}

	var input io.Reader
	if *expressionFlag != "" {
		input = strings.NewReader(*expressionFlag)
	} else if *fileFlag != "" {
		file, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Помилка: не вдалося відкрити файл: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "Помилка: не вказано вхідний вираз (використовуйте -e або -f)")
		os.Exit(1)
	}

	var output io.Writer
	if *outputFlag != "" {
		file, err := os.Create(*outputFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Помилка: не вдалося створити файл: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	computeHandler := handler.ComputeHandler{
		Input:  input,
		Output: output,
	}

	if err := computeHandler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "Помилка:", err)
		os.Exit(1)
	}
}
