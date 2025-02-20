package main

import (
	"errors"
	"fmt"

	"github.com/ivansevryukov1995/Task-Tracker/internal"
)

const jsonFileName = "task-tracker.json"

func main() {
	var tasks internal.Tasks

	storage := internal.NewStorage(jsonFileName)

	if err := storage.Load(&tasks); err != nil {
		fmt.Println(errors.Unwrap(err))
		fmt.Printf("File %v will be created.\n", jsonFileName)
	}

	cmdFlags := internal.NewCmdFlags()
	cmdFlags.ExecuteCmd(&tasks)

	if err := storage.Save(&tasks); err != nil {
		fmt.Println(errors.Unwrap(err))
	}

}
