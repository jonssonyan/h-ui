package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func RemoveFile(fileName string) error {
	if Exists(fileName) {
		if err := os.Remove(fileName); err != nil {
			return errors.New("failed to delete file")
		}
	}
	return nil
}

// ReadLinesFromBottom Read the file contents sequentially from bottom to top and return the specified number of lines
func ReadLinesFromBottom(filePath string, numLines int) ([]string, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// Read the file contents line by line and reverse the order of the lines
	total := 0
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		total++
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}

	// Reverse row order
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	// Returns the specified number of rows
	if len(lines) < numLines {
		numLines = len(lines)
	}
	return lines[:numLines], total, nil
}

func FindFile(dir, filename string) (string, error) {
	var result string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == filename {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			result = absPath
			return errors.New("file found")
		}
		return nil
	})
	if err != nil && err.Error() != "file found" {
		return "", err
	}
	if result == "" {
		return "", fmt.Errorf("file %s not found in directory %s", filename, dir)
	}
	return result, nil
}
