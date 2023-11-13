package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Установка времени запуска и завершения экзешника.
	startTime := "16:11:00"
	endTime := "16:11:10"

	// Путь к экзеешнику.
	scheduleExecutable(startTime, endTime, "C:/Program Files/AIMP/AIMP.exe")
}

func scheduleExecutable(startTime, endTime, executablePath string) {
	firstRun := true

	for {
		currentTime := time.Now().Format("15:04:05") //Формат даты, можно любое значение.
		if currentTime >= startTime && currentTime <= endTime {
			if firstRun {
				fmt.Println("Program is running...")
				firstRun = false
			}
			runExecutable(executablePath)
		} else if currentTime > endTime {
			fmt.Println("Program is completed.")
			endExecutable(executablePath)
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

func endExecutable(executablePath string) {
	cmd := exec.Command("taskkill", "/F", "/IM", "AIMP.exe") // Вместо "AIMP.exe нужно использовать название процесса запущенного экзешника, либо путь к нему БЕЗ пробелов!"
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
