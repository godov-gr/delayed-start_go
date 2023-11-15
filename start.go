package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Определение флагов для времени и пути.
	var startTime string
	var endTime string
	var executablePath string

	// Установка значений флагов по умолчанию.
	flag.StringVar(&startTime, "start", "16:00:00", "Время начала (в формате HH:MM:SS)")
	flag.StringVar(&endTime, "end", "17:00:00", "Время окончания (в формате HH:MM:SS)")
	flag.StringVar(&executablePath, "path", "", "Путь к файлу")

	// Парсинг аргументов командной строки.
	flag.Parse()

	// Если путь к файлу не был указан, запросите у пользователя ввод пути к файлу.
	if executablePath == "" {
		fmt.Print("Введите путь к файлу: ")
		fmt.Scanln(&executablePath)
	}

	scheduleExecutable(startTime, endTime, executablePath)
}

func scheduleExecutable(startTime, endTime, executablePath string) {
	firstRun := true

	// Извлекаем название процесса из пути к файлу
	processName := filepath.Base(executablePath)

	for {
		currentTime := time.Now().Format("15:04:05")

		if currentTime >= startTime && currentTime <= endTime && !processExists(processName) {
			if firstRun {
				fmt.Println("Program is running...")
				runExecutable(executablePath)
				firstRun = false
			}

		} else if currentTime > endTime {
			fmt.Println("Program is completed.")
			endExecutable(processName)
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func runExecutable(executablePath string) {
	cmd := exec.Command(executablePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func endExecutable(processName string) {
	cmd := exec.Command("taskkill", "/F", "/IM", processName)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func processExists(processName string) bool {
	cmd := exec.Command("tasklist")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	return strings.Contains(string(output), processName)
}
