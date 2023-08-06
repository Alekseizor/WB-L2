package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Функция для обработки команды cd
func handleCd(args []string) {
	switch {
	case len(args) > 1:
		fmt.Println("Too many arguments")
	case len(args) <= 0:
		fmt.Println("Missing argument for cd command")
	default:
		err := os.Chdir(args[0])
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Функция для обработки команды pwd
func handlePwd() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cwd)
}

// Функция для обработки команды echo
func handleEcho(command string) {
	fmt.Println(command)
}

// Функция для обработки команды kill
func handleKill(args []string) {
	if len(args) > 0 {
		err := exec.Command("kill", args...).Run()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Missing argument for kill command")
	}
}

// Функция для обработки команды ps
func handlePs() {
	cmd := exec.Command("ps")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

// Функция для выполнения команды с поддержкой конвейера на пайпах
func executeCommandWithPipes(command string) {
	cmds := strings.Split(command, "|")
	var output []byte
	var err error
	for _, cmd := range cmds {
		cmd = strings.TrimSpace(cmd)
		parts := strings.Fields(cmd)
		if len(parts) > 0 {
			cmd := exec.Command(parts[0], parts[1:]...)
			if len(output) > 0 {
				cmd.Stdin = strings.NewReader(string(output))
			}
			output, err = cmd.Output()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fmt.Println(string(output))
}

func Shell(reader *bufio.Reader) error {
	for {
		fmt.Print(">> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		command = strings.TrimSuffix(command, "\n")
		command = strings.TrimSpace(command)

		switch {
		case command == "exit":
			return nil
		case strings.Contains(command, "|"):
			executeCommandWithPipes(command)
		case strings.HasPrefix(command, "cd "):
			args := strings.Split(command, " ")[1:]
			handleCd(args)
		case command == "pwd":
			handlePwd()
		case strings.HasPrefix(command, "echo "):
			handleEcho(command[5:])
		case strings.HasPrefix(command, "kill "):
			args := strings.Split(command, " ")[1:]
			handleKill(args)
		case command == "ps":
			handlePs()
		default:
			cmd := exec.Command("bash", "-c", command)
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Print(string(output))
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	err := Shell(reader)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
