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

// Load reads task data from a JSON file and populates the provided Tasks collection.
// It attempts to read the file specified by the Storage's FileName and unmarshal
// the JSON data into the Tasks object passed as a parameter.
//
// Parameters:
//   - tasks: A pointer to the Tasks collection that will be populated with the loaded data.
//
// Returns:
//   - An error if the file cannot be read or if the unmarshaling fails; otherwise, it returns nil.
func (s *Storage) Load(tasks *Tasks) error {
	data, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, tasks)
}

// Save writes the provided Tasks collection to a JSON file.
// It marshals the Tasks into JSON format and writes the resulting data
// to the file specified by the Storage's FileName.
//
// Parameters:
//   - tasks: A pointer to the Tasks collection that will be saved to the file.
//
// Returns:
//   - An error if the marshaling fails or if the file cannot be written; otherwise, it returns nil.
func (s *Storage) Save(tasks *Tasks) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, data, 0777)
}
