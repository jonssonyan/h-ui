package util

import (
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

const (
	echConfigsPemType = "ECH CONFIGS"
	echKeyFileMaxSize = 64 * 1024
)

func ParseEchConfigList(data []byte) (string, error) {
	rest := data
	for {
		var block *pem.Block
		block, rest = pem.Decode(rest)
		if block == nil {
			break
		}
		if block.Type == echConfigsPemType {
			if len(block.Bytes) == 0 {
				return "", errors.New("ECH CONFIGS block is empty")
			}
			return base64.StdEncoding.EncodeToString(block.Bytes), nil
		}
	}
	return "", errors.New("ECH CONFIGS block not found")
}

func ReadEchConfigList(keyPath string) (string, error) {
	info, err := os.Stat(keyPath)
	if err != nil {
		return "", fmt.Errorf("ECH key file not found: %s", keyPath)
	}
	if info.IsDir() || info.Size() > echKeyFileMaxSize {
		return "", fmt.Errorf("invalid ECH key file: %s", keyPath)
	}
	data, err := os.ReadFile(keyPath)
	if err != nil {
		return "", fmt.Errorf("read ECH key file err: %v", err)
	}
	return ParseEchConfigList(data)
}
