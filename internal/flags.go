package internal

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add         string
	Update      string
	Delete      uint
	Mark_status string
	List        string
}

func NewCmdFlags() *CmdFlags {
	cmdFlags := &CmdFlags{}
	flag.StringVar(&cmdFlags.Add, "add", "", "Adding a new task")
	flag.StringVar(&cmdFlags.Update, "update", "", "Updating task.  Use format: ID \"new description\" or ID new description")
	flag.UintVar(&cmdFlags.Delete, "delete", 0, "Deleting task")
	flag.StringVar(&cmdFlags.Mark_status, "mark", "", "Marking a task as in progress or done")
	flag.StringVar(&cmdFlags.List, "list", "", "Listing all tasks or listing tasks by status")

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

		return tasks.UpdateDescription(uint(id), strings.Join(desc, " "))

	case cf.Delete != 0:
		tasks.Delete(cf.Delete)

		return nil

	case cf.Mark_status != "":
		args := strings.SplitN(cf.Mark_status, " ", 2)
		if len(args) < 2 {
			return errors.New("mark_status requires task ID and status")
		}

		ind, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		return tasks.UpdateStatus(uint(ind), args[1])

	case cf.List != "":
		tasks.List(cf.List)

		return nil

	default:
		flag.PrintDefaults()

		return errors.New("no command specified")
	}
}
