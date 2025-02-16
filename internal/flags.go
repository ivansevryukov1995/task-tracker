package internal

import "flag"

type CmdFlags struct {
	Add         string
	Update      string
	Delete      uint
	Mark_status uint
	List        string
}

func NewCmdFlags() *CmdFlags {
	return &CmdFlags{
		Add:         *flag.String("add", "", "Adding a new task"),
		Update:      *flag.String("update", "", "Updating task"),
		Delete:      *flag.Uint("delete", 0, "Deleting task"),
		Mark_status: *flag.Uint("mark", 0, ""),
		List:        *flag.String("list", "all", "Listing all tasks or listing tasks by status"),
	}
}
