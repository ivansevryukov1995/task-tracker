package internal

import (
	"encoding/json"
	"os"
)

type Storage struct {
	FileName string
}

func NewStorage(fileName string) *Storage {
	return &Storage{
		FileName: fileName,
	}
}

func (s *Storage) Load(tasks *Tasks) error {
	data, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, tasks)
}

func (s *Storage) Save(tasks *Tasks) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, data, 0777)
}
