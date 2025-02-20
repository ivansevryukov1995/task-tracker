package internal

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add            string
	Update         string
	Delete         uint
	MarkInProgress uint
	MarkDone       uint
	List           string
}

func NewCmdFlags() *CmdFlags {
	cmdFlags := &CmdFlags{}
	flag.StringVar(&cmdFlags.Add, "add", "", "Adding a new task")
	flag.StringVar(&cmdFlags.Update, "update", "", "Updating task. Use format: ID \"new description\" or ID new description")
	flag.UintVar(&cmdFlags.Delete, "delete", 0, "Deleting task")
	flag.UintVar(&cmdFlags.MarkInProgress, "mark-in-progress", 0, "Marking a task as in progress")
	flag.UintVar(&cmdFlags.MarkDone, "mark-done", 0, "Marking a task as done")
	flag.StringVar(&cmdFlags.List, "list", "", "Listing all tasks or listing tasks by status. Use \"\", \"done\", \"todo\" or \"in-progress\"")

	flag.Parse()

	return cmdFlags
}

func (cf *CmdFlags) ExecuteCmd(tasks *Tasks) error {
	switch {
	case cf.Add != "":
		fmt.Printf("Task added successfully (ID: %v)\n", tasks.Add(cf.Add))
		return nil

	case cf.Update != "":
		parts := strings.SplitN(cf.Update, " ", 2) // Split into two parts maximum
		desc := flag.Args()

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("invalid ID: %w", err)
		}

		tasks.UpdateDescription(uint(id), strings.Join(desc, " "))
		return nil

	case cf.Delete != 0:
		tasks.Delete(cf.Delete)
		return nil

	case cf.MarkInProgress != 0:
		tasks.UpdateStatus(cf.MarkInProgress, "in-progress")
		return nil

	case cf.MarkDone != 0:
		tasks.UpdateStatus(cf.MarkDone, "done")
		return nil

	case cf.List == "" || cf.List == "done" || cf.List == "todo" || cf.List == "in-progress":
		tasks.List(cf.List)
		return nil

	default:
		flag.PrintDefaults()
		return errors.New("no command specified")
	}
}
