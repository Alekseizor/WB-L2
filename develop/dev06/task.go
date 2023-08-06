package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	ErrFields = fmt.Errorf("you must specify a list of fields")
)

func main() {
	// Определение флагов
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	if *fields == "" {
		log.Println(ErrFields)
	}
	err := Cut(os.Stdin, *fields, *delimiter, *separated)
	if err != nil {
		log.Println(err)
	}
}

// Разбор запрошенных полей и возврат списка индексов
func parseFieldIndexes(fields string, maxIndex int) ([]int, error) {
	fieldIndexes := make([]int, 0)

	// Разделение запрошенных полей по запятой
	fieldsList := strings.Split(fields, ",")
	for _, field := range fieldsList {
		// Обработка диапазона полей (например, 1-3)
		rangeIndexes := strings.Split(field, "-")
		if len(rangeIndexes) == 2 {
			startIndex, err := parseFieldIndex(rangeIndexes[0], maxIndex)
			if err != nil {
				return nil, err
			}
			endIndex, err := parseFieldIndex(rangeIndexes[1], maxIndex)
			if err != nil {
				return nil, err
			}
			for i := startIndex; i <= endIndex; i++ {
				fieldIndexes = append(fieldIndexes, i)
			}
		} else {
			// Обработка отдельного поля
			fieldIndex, err := parseFieldIndex(field, maxIndex)
			if err != nil {
				return nil, err
			}
			fieldIndexes = append(fieldIndexes, fieldIndex)
		}
	}
	return fieldIndexes, nil
}

// Разбор индекса поля и возврат соответствующего индекса
func parseFieldIndex(fieldIndexStr string, maxIndex int) (int, error) {
	fieldIndex, err := strconv.Atoi(fieldIndexStr)
	if err != nil {
		return 0, err
	}
	fieldIndex--
	switch {
	case fieldIndex <= 0:
		return 0, nil
	case fieldIndex >= maxIndex:
		return maxIndex, nil
	default:
		return fieldIndex, nil
	}
}

func Cut(r io.Reader, fields, delimiter string, separated bool) error {
	// Чтение строк из STDIN
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		// Проверка наличия разделителя в строке, если включен флаг -s
		if separated && !strings.Contains(line, delimiter) {
			continue
		}

		// Разделение строки на поля
		fieldsList := strings.Split(line, delimiter)

		fieldIndexes, err := parseFieldIndexes(fields, len(fieldsList)-1)
		if err != nil {
			return err
		}
		var selectedFields []string
		for _, index := range fieldIndexes {
			selectedFields = append(selectedFields, fieldsList[index])
		}
		fmt.Println(strings.Join(selectedFields, delimiter))
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
