package util

import (
	"bufio"
	"errors"
	"os"
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

// ReadLinesFromBottom 从下往上顺序读取文件内容并返回指定数量的行数
func ReadLinesFromBottom(filePath string, numLines int) ([]string, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容并反转行顺序
	total := 0
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		total++
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}

	// 反转行顺序
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	// 返回指定数量的行数
	if len(lines) < numLines {
		numLines = len(lines)
	}
	return lines[:numLines], total, nil
}
