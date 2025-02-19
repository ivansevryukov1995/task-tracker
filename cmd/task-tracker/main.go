package main

import (
	"errors"
	"fmt"

	"github.com/ivansevryukov1995/Task-Tracker/internal"
)

const jsonFile = "task-tracker.json"

func main() {
	var tasks internal.Tasks

	if err := tasks.Load(jsonFile); err != nil {
		fmt.Println(errors.Unwrap(err))
		fmt.Printf("File %v will be created.\n", jsonFile)
	}

	cmdFlags := internal.NewCmdFlags()
	cmdFlags.ExecuteCmd(&tasks)

	if err := tasks.Unload(jsonFile); err != nil {
		fmt.Println(errors.Unwrap(err))
	}

}
