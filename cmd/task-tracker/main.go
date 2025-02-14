package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ivansevryukov1995/Task-Tracker/internal"
)

const jsonFile = "task-tracker.json"

func CreateFileJSON(nameFile string) error {
	_, err := os.Stat(nameFile)
	if os.IsNotExist(err) {
		file, err := os.Create(nameFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		return nil

	} else if err != nil {
		return err
	} else {
		return nil
	}
}

func main() {
	var tasks internal.Tasks

	if err := CreateFileJSON(jsonFile); err != nil {
		log.Fatal(err)
	}

	if err := tasks.Load(jsonFile); err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)
}
