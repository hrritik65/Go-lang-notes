package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines(path string) ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println("Error reading file step 1:", err)
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file step 2:", err)
		return nil, err
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteJson(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		fmt.Println("Error encoding json:", err)
		return err
	}
	file.Close()
	return nil

}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
