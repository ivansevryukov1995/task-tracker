package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ivansevryukov1995/Task-Tracker/internal"
)

const jsonFile = "task-tracker.json"

func CreateFileJSON(nameFile string) error {
	_, err := os.Stat(nameFile)
	if os.IsNotExist(err) {
		file, err := os.Create(nameFile)
		if err != nil {
			errors.Unwrap(err)
		}
		defer file.Close()

		return nil

	} else if err != nil {
		return errors.Unwrap(err)
	} else {
		return nil
	}
}

func main() {
	var tasks internal.Tasks

	if err := CreateFileJSON(jsonFile); err != nil {
		fmt.Println(errors.Unwrap(err))
	}

	if err := tasks.Load(jsonFile); err != nil {
		fmt.Println(errors.Unwrap(err))
	}

	_, err := tasks.Add("Привет, Иван!!!")
	if err != nil {
		fmt.Println(errors.Unwrap(err))
	}

	if err := tasks.Unload(jsonFile); err != nil {
		fmt.Println(errors.Unwrap(err))
	}

	if err := tasks.Delete(8); err != nil {
		fmt.Println(err)
	}

	if err := tasks.Unload(jsonFile); err != nil {
		fmt.Println(errors.Unwrap(err))
	}

	if err := tasks.UpdateStatus(6, "done"); err != nil {
		fmt.Println(errors.Unwrap(err))
	}

	if err := tasks.Unload(jsonFile); err != nil {
		fmt.Println(errors.Unwrap(err))
	}

	fmt.Println(tasks)
}
