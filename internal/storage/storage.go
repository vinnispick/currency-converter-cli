package storage

import (
	"fmt"
	"os"
)

func OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(filePath)
			if err != nil {
				return nil, fmt.Errorf("failed to create file %s: %w", filePath, err)
			}
			return file, nil
		}
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	return file, nil
}

func MarshalFile(file *os.File) ([]byte, error) {
	data, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", file.Name(), err)
	}
	if len(data) == 0 {
		return nil, nil
	}
	return data, nil
}

func SaveFile(filePath string, data []byte) error {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filePath, err)
	}
	return nil
}
