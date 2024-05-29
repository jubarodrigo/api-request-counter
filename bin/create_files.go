package bin

import (
	"fmt"
	"os"
	"time"

	"counter/config"
)

const tickerTime = 1 * time.Minute

func createFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
}

func RunTicker(fileConfig config.FileConfig) {
	createFile(fileConfig.FilePath)
	ticker := time.NewTicker(tickerTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			createFile(fileConfig.FilePath)
		}
	}
}
