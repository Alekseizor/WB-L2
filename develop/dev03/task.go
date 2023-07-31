package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filePath := flag.String("file", "", "Path to the input file")
	column := flag.Int("k", 0, "Column number for sorting")
	numeric := flag.Bool("n", false, "Sort by numeric value")
	reverse := flag.Bool("r", false, "Sort in reverse order")
	unique := flag.Bool("u", false, "Remove duplicate lines")

	flag.Parse()

	if *filePath == "" {
		log.Fatal("Please provide the path to the input file using the -file flag")
	}
	//сразу разбираемся с флагом u
	lines, err := readLines(*filePath, *unique)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	//разбираемся с флагами k и n
	lines = sortLine(lines, *column, *numeric)

	//разбираемся с флагом r
	err = writeLines(lines, *reverse)
	if err != nil {
		log.Fatalf("Failed to write output: %v", err)
	}
}

func readLines(filePath string, unique bool) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	uniqueRows := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	var lines [][]string
	for i := 0; scanner.Scan(); i++ {
		if unique == true && uniqueRows[scanner.Text()] {
			continue
		}
		lines = append(lines, strings.Split(scanner.Text(), " "))
		uniqueRows[scanner.Text()] = true
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func writeLines(lines [][]string, reverse bool) error {
	var content strings.Builder
	if reverse {
		for index := len(lines) - 1; index >= 0; index-- {
			finalLine := strings.Join(lines[index], " ")
			content.WriteString(finalLine)
			content.WriteString("\n")
		}
	} else {
		for _, line := range lines {
			finalLine := strings.Join(line, " ")
			content.WriteString(finalLine)
			content.WriteString("\n")
		}
	}
	fmt.Println(content.String())
	return nil
}

func sortLine(lines [][]string, k int, n bool) [][]string {
	lineWithoutNumber := make([][]string, 0)
	lineWithNumber := make([][]string, 0)
	if n {
		for _, line := range lines {
			if len(line) <= k {
				sort.Slice(lines, func(i, j int) bool {
					return lines[i][0] < lines[j][0]
				})
				return lines
			}
			_, err := strconv.Atoi(line[k])
			if err != nil {
				lineWithoutNumber = append(lineWithoutNumber, line)
				continue
			}
			lineWithNumber = append(lineWithNumber, line)
		}
		sort.Slice(lineWithNumber, func(i, j int) bool {
			numberFirst, err := strconv.Atoi(lineWithNumber[i][k])
			if err != nil {
				log.Println(err)
				return true
			}
			numberSecond, err := strconv.Atoi(lineWithNumber[j][k])
			if err != nil {
				log.Println(err)
				return true
			}
			if numberFirst < numberSecond {
				return true
			}
			return false
		})
		sort.Slice(lineWithoutNumber, func(i, j int) bool {
			return lineWithoutNumber[i][0] < lineWithoutNumber[j][0]
		})
		return append(lineWithoutNumber, lineWithNumber...)
	}
	sort.Slice(lines, func(i, j int) bool {
		if len(lines[i]) <= k || len(lines[j]) <= k {
			return lines[i][0] < lines[j][0]
		}
		return lines[i][k] < lines[j][k]
	})
	return lines
}
