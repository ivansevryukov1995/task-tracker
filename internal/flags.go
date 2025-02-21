package internal

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type CmdFlags struct {
	Add            string
	Update         uint
	Delete         uint
	MarkInProgress uint
	MarkDone       uint
	List           string
}

func NewCmdFlags() *CmdFlags {
	cmdFlags := &CmdFlags{}
	flag.StringVar(&cmdFlags.Add, "add", "", "Adding a new task")
	flag.UintVar(&cmdFlags.Update, "update", 0, "Updating task. Use format: ID new description")
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

	case cf.Update != 0:
		tasks.UpdateDescription(cf.Update, strings.Join(flag.Args(), " "))
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
