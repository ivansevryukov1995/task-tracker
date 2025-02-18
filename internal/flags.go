package internal

import (
	"errors"
	"flag"
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
	cmdFlags := &CmdFlags{
		Add:         *flag.String("add", "", "Adding a new task"),
		Update:      *flag.String("update", "", "Updating task"),
		Delete:      *flag.Uint("delete", 0, "Deleting task"),
		Mark_status: *flag.String("mark", "", "Marking a task as in progress or done"),
		List:        *flag.String("list", "", "Listing all tasks or listing tasks by status"),
	}

	flag.Parse()

	return cmdFlags
}

func (cf *CmdFlags) Parse(tasks *Tasks) error {
	switch {
	case cf.Add != "":
		tasks.Add(cf.Add)
	case cf.Update != "":
		args := strings.SplitN(cf.Update, " ", 0)

		ind, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.Unwrap(err)
		}

		tasks.UpdateDescription(uint(ind), args[1])
	case cf.Delete != 0:
		tasks.Delete(cf.Delete)
	case cf.Update != "":
		args := strings.SplitN(cf.Update, " ", 0)

		ind, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.Unwrap(err)
		}

		tasks.UpdateStatus(uint(ind), args[1])
	case cf.List != "":

		tasks.List(cf.List)

	}
	return nil
}
