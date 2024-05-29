package bin

import (
	"fmt"
	"os"
	"time"
)

func createFile() {
	fileName := "./storage/files/data.txt"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()
	now := time.Now().Format("2006-01-02 15:04:05")
	_, err = file.WriteString(fmt.Sprintf("Current time: %s", now))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func RunTicker() {
	createFile()
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			createFile()
		}
	}
}
