package internal

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status" validate:"oneof=todo in-progress done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks []Task

func NewTask(id uint, desc string, status string) *Task {
	return &Task{
		ID:          id,
		Description: desc,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Tasks) Load(nameFile string) error {
	data, err := os.ReadFile(nameFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	return nil
}

func (t *Tasks) Add(id uint, desc string) error {

	task := NewTask(id, desc, "todo")

	*t = append(*t, *task)
	return nil
}

func (t *Tasks) Delete(id uint) error {
	t[]
	*t = append(*t[:id], *t[id+1:]...)
	return nil
}

func (t *Tasks) Update(task Task) error {
	// TODO:
	return nil
}
