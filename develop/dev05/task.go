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
	file := flag.Args()[1]

	lines, err := readFile(file)
	if err != nil {
		log.Fatalf("File reading error %s: %s", file, err)
	}
	counter, err := processFile(lines, pattern, *after, *before, *context, *ignoreCase, *fixed, *count, *invert, *lineNum)
	if err != nil {
		log.Fatalf("Error processing file %s: %s", file, err)
	}
	if *count {
		fmt.Println(counter)
	}
}

func readFile(file string) ([]*RowFlag, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []*RowFlag // Строки для печати

	// Проходим по каждой строке файла
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, &RowFlag{row: line})
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// Обрабатывает файл с учетом заданных параметров
func processFile(lines []*RowFlag, pattern string, after, before, context int, ignoreCase, fixed, count, invert, lineNum bool) (int, error) {
	// Создаем регулярное выражение на основе паттерна
	var re *regexp.Regexp
	switch {
	case fixed:
		re = regexp.MustCompile(regexp.QuoteMeta(pattern))
	case ignoreCase:
		re = regexp.MustCompile(`(?i)` + pattern)
	default:
		re = regexp.MustCompile(pattern)
	}
	var startIndex, stopIndex int
	var counter int
	for numLine, line := range lines {
		// Проверяем на соответствие паттерну
		match := re.MatchString(line.row)
		//если строка не должна быть выведена, то сразу переходим к следующей
		if (match && invert) || (!match && !invert) {
			continue
		}
		if count {
			counter++
			continue
		}
		switch {
		case context > 0:
			if context >= numLine {
				startIndex = 0
			} else {
				startIndex = numLine - context
			}
			if numLine+context >= len(lines)-1 {
				stopIndex = len(lines) - 1
			} else {
				stopIndex = numLine + context
			}
		case before > 0:
			if before >= numLine {
				startIndex = 0
			} else {
				startIndex = numLine - before
			}
			stopIndex = numLine
		case after > 0:
			if numLine+after >= len(lines)-1 {
				stopIndex = len(lines) - 1
			} else {
				stopIndex = numLine + after
			}
			startIndex = numLine
		default:
			startIndex = numLine
			stopIndex = numLine
		}
		printLines(lines[startIndex:stopIndex+1], lineNum)
	}
	return counter, nil
}

// Печатает строки с номерами (если необходимо)
func printLines(lines []*RowFlag, lineNum bool) {
	for i, line := range lines {
		if line.output {
			continue
		}
		line.output = true
		if lineNum {
			fmt.Printf("%d: %s\n", i+1, line.row)
		} else {
			fmt.Println(line.row)
		}
	}
}
