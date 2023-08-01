package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

type RowFlag struct {
	row    string
	output bool
}

func main() {
	// Объявляем флаги
	after := flag.Int("A", 0, "Print N lines after the match")
	before := flag.Int("B", 0, "Print N lines before the match")
	context := flag.Int("C", 0, "Print N lines around the match")
	count := flag.Bool("c", false, "Print only the count of matching lines") //выводим только количество, для нас важны только флаги
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert the match")
	fixed := flag.Bool("F", false, "Fixed string match")
	lineNum := flag.Bool("n", false, "Print line numbers")

	// Парсим флаги командной строки
	flag.Parse()

	// Получаем аргументы командной строки (паттерн и файлы)
	pattern := flag.Arg(0)
	log.Println(pattern)
	file := flag.Args()[1]

	// Создаем регулярное выражение на основе паттерна
	var re *regexp.Regexp
	if *fixed {
		re = regexp.MustCompile(regexp.QuoteMeta(pattern))
	} else {
		if *ignoreCase {
			re = regexp.MustCompile(`(?i)` + pattern)
		} else {
			re = regexp.MustCompile(pattern)
		}
	}
	lines, err := readFile(file)
	if err != nil {
		log.Fatalf("File reading error %s: %s", file, err)
	}
	err = processFile(lines, re, *after, *before, *context, *count, *invert, *lineNum)
	if err != nil {
		log.Fatalf("Error processing file %s: %s", file, err)
	}
}

func readFile(file string) ([]RowFlag, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []RowFlag // Строки для печати

	// Проходим по каждой строке файла
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, RowFlag{row: line})
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// Обрабатывает файл с учетом заданных параметров
func processFile(lines []RowFlag, re *regexp.Regexp, after, before, context int, count, invert, lineNum bool) error {
	// Проходим по каждой строке файла
	for numLine, line := range lines {
		// Проверяем на соответствие паттерну
		match := re.MatchString(line.row)

		// Определяем, должна ли эта строка быть напечатана
		printLine := false

		if ((!invert && match) || (invert && !match)) && !line.output {
			printLine = true
		}
		if context > 0 {
			if len(lines) > context {
				lines = lines[1:]
			}
			lines = append(lines, line)
			if printLine {
				printLines(lines, lineNum)
				lines = []string{}
			}
		} else if before > 0 {
			if len(lines) == before {
				printLines(lines, lineNum)
				lines = []string{}
			}
			if printLine {
				lines = []string{}
			}
			lines = append(lines, line)
		} else if after > 0 {
			if printLine {
				lines = append(lines, line)
				printLines(lines, lineNum)
				lines = []string{}
			} else if len(lines) > 0 {
				lines = append(lines, line)
			}
		} else if count {
			if printLine {
				lines = append(lines, line)
			}
		} else {
			if printLine {
				fmt.Println(line)
			}
		}
	}

	// Проверяем наличие ошибок в процессе сканирования файла
	if err := scanner.Err(); err != nil {
		return err
	}

	// Если нужно вывести количество совпадений
	if count {
		fmt.Printf("Count: %d\n", len(lines))
	}

	return nil
}

// Печатает строки с номерами (если необходимо)
func printLines(lines []string, lineNum bool) {
	for i, line := range lines {
		if lineNum {
			fmt.Printf("%d: %s\n", i+1, line)
		} else {
			fmt.Println(line)
		}
	}
}
