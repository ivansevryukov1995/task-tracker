package internal

import (
	"fmt"
	"strings"
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

// Add creates a new task with the specified description and adds it to the Tasks collection.
// It generates a unique ID for the new task, either by incrementing the ID of the last task
// in the collection or initializing it to 1 if the collection is empty.
// The new task is then appended to the collection, and the newly generated ID is returned.
//
// Parameters:
//   - desc: A string representing the description of the new task.
//
// Returns:
//   - The unique ID assigned to the newly created task.
func (t *Tasks) Add(desc string) uint {
	var id uint
	if len(*t) != 0 {
		id = (*t)[len(*t)-1].ID + 1
	} else {
		id = 1
	}

	task := NewTask(desc)
	task.ID = id

	*t = append(*t, *task)

	return id
}

// Delete removes the task with the specified ID from the Tasks collection.
// It searches for the task with the given ID and, if found, removes it from the list.
//
// Parameters:
//   - id: The ID of the task to be deleted.
//
// Returns:
//   - An error if the task with the specified ID does not exist; otherwise, it returns nil.
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

// UpdateStatus updates the status of the task with the specified ID.
// It searches through the Tasks collection to find the task matching the given ID
// and modifies its Status field with the provided new status. Additionally, it updates
// the UpdatedAt timestamp to the current time to reflect this change.
//
// Parameters:
//   - id: The ID of the task whose status needs to be updated.
//   - status: The new status to set for the task.
//
// If the task with the specified ID is not found, no changes are made.
func (t *Tasks) UpdateStatus(id uint, status string) {
	for i, task := range *t {
		if task.ID == id {
			(*t)[i].Status = status
			(*t)[i].UpdatedAt = time.Now()
			break
		}
	}
}

// UpdateDescription updates the description of the task with the specified ID.
// It searches for the task in the Tasks collection and modifies its Description field
// with the provided new description. Additionally, it updates the UpdatedAt timestamp
// to the current time to reflect the change.
//
// Parameters:
//   - id: The ID of the task whose description needs to be updated.
//   - desc: The new description to set for the task.
//
// If the task with the specified ID is not found, no changes are made.
func (t *Tasks) UpdateDescription(id uint, desc string) {
	for i, task := range *t {
		if task.ID == id {
			(*t)[i].Description = desc
			(*t)[i].UpdatedAt = time.Now()
			break
		}
	}
}

// List displays a list of tasks in a tabular format in the console.
// For each task, it shows the ID, description, status, creation date, and update date.
// If the `status` parameter is not empty, only tasks with the specified status are displayed.
// The description of each task is automatically wrapped to new lines if it exceeds 50 characters.
// For the first line with the description, all task data is displayed, while subsequent lines show only the description.
// Parameters:
//   - status: a string that filters tasks based on their status.
//     If empty, all tasks are displayed.
func (t *Tasks) List(status string) {
	fmt.Printf("%-5s %-30s %-15s %-15s %-15s\n", "ID", "Description", "Status", "CreatedAt", "UpdatedAt")
	for _, task := range *t {
		if task.Status == status || status == "" {
			descriptionLines := wrapText(task.Description, 50)
			for i, line := range descriptionLines {
				if i == 0 {
					fmt.Printf("%-5d %-30s %-15s %-15s %-15s\n",
						task.ID,
						line,
						task.Status,
						task.CreatedAt.Format("02.01.2006"),
						task.UpdatedAt.Format("02.01.2006"))
				} else {
					fmt.Printf("%-5s %-30s\n",
						"",
						line)
				}
			}
		}
	}
}

// wrapText wraps the input text to a specified line length limit.
// If the text exceeds the specified limit, it will be broken into multiple lines.
// Each line will be at most `limit` characters long.
// Words will not be split; if a word exceeds the limit, it will be placed entirely on the new line.
// This function helps in formatting text output for better readability.
// Parameters:
//   - text: The input string to be wrapped.
//   - limit: The maximum number of characters allowed per line.
//
// Returns:
//   - A string with the input text wrapped according to the specified limit.
func wrapText(text string, limit int) []string {
	var lines []string
	words := strings.Fields(text)

	var currentLine string
	for _, word := range words {
		if len(currentLine)+len(word)+1 > limit {
			if currentLine != "" {
				lines = append(lines, currentLine)
			}
			currentLine = word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}
