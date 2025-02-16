package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          uint      `json:"id,omitempty"`
	Description string    `json:"description"`
	Status      string    `json:"status" validate:"oneof=todo in-progress done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks []Task

func NewTask(desc string) *Task {
	return &Task{
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Tasks) Load(nameFile string) error {
	data, err := os.ReadFile(nameFile)
	if err != nil {
		return errors.Unwrap(err)
	}

	if err := json.Unmarshal(data, t); err != nil {
		return errors.Unwrap(err)
	}

	return nil
}

func (t *Tasks) Unload(nameFile string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return errors.Unwrap(err)
	}

	if err := os.WriteFile(nameFile, data, 0644); err != nil {
		return errors.Unwrap(err)
	}

	return nil
}

// The method adds a task to the general task list
func (t *Tasks) Add(desc string) (uint, error) {
	var id uint
	if len(*t) != 0 {
		id = (*t)[len(*t)-1].ID + 1
	} else {
		id = 1
	}

	task := NewTask(desc)

	task.ID = id

	*t = append(*t, *task)

	return id, nil
}

// The method deletes the issue by the specified ID.
func (t *Tasks) Delete(id uint) error {
	ind := -1
	for i, task := range *t {
		if task.ID == id {
			ind = i
			break
		}
	}
	if ind == -1 {
		return fmt.Errorf("Task ID=%v not exists", id)
	}
	*t = append((*t)[:ind], (*t)[ind+1:]...)
	return nil
}

// The method updates the task status by the id number.
// The todo, in-progress, and done statuses are available.
func (t *Tasks) UpdateStatus(id uint, status string) error {
	for i, task := range *t {
		if task.ID == id {
			(*t)[i].Status = status
			(*t)[i].UpdatedAt = time.Now()
			break
		}
	}

	return nil
}

// The method updates the task description by the identification number.
func (t *Tasks) UpdateDescription(id uint, desc string) error {
	for i, task := range *t {
		if task.ID == id {
			(*t)[i].Description = desc
			(*t)[i].UpdatedAt = time.Now()
			break
		}
	}

	return nil
}
