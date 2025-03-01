package main

import (
	"log"

	"github.com/ivansevryukov1995/Task-Tracker/internal"
)

const jsonFileName = "task-tracker.json"

func main() {
	var tasks internal.Tasks

	storage := internal.NewStorage(jsonFileName)

	if err := storage.Load(&tasks); err != nil {
		log.Printf("Error load task: %v\nFile %v will be automatically created when you add the task with the --add command.",
			err, jsonFileName)
	}

	cmdFlags := internal.NewCmdFlags()
	if err := cmdFlags.ExecuteCmd(&tasks); err != nil {
		log.Fatalf("Flag not valid: %v", err)
	}

	if err := storage.Save(&tasks); err != nil {
		log.Fatalf("Error save task: %v", err)
	}

}
