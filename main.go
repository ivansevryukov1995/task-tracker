package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const nameFileJSON = "task-tracker.json"

type Task struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status" validate:"oneof=todo in-progress done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks []Task

func NewTask(id uint, desc string, status string) *Task {
	return &Task{
		Id:          id,
		Description: desc,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Tasks) Load(nameFileJSON string) error {
	_, err := os.Stat(nameFileJSON)
	if os.IsNotExist(err) {
		file, err := os.Create(nameFileJSON)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		return nil

	} else if err != nil {
		return err

	} else {
		data, err := os.ReadFile(nameFileJSON)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(data, t); err != nil {
			return err
		}
		return nil
	}
}

func main() {
	var tasks Tasks

	if err := tasks.Load(nameFileJSON); err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasks)
}
