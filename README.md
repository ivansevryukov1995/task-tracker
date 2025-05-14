# Task-Tracker

<a id='anchor'></a>
## CLI Application Task Tracker
Task tracker is a [application](https://roadmap.sh/projects/task-tracker) used to track and manage your tasks.

## Features

- **Add a Task:** Add a new task with a description.
- **Update a Task:** Update the description of an existing task.
- **Delete a Task:** Remove a task by its ID.
- **Mark a Task:** Mark a task as "in-progress" or "done".
- **List Tasks:** List all tasks or filter them by status (e.g., `todo`, `in progress`, `done`).

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/ivansevryukov1995/task-tracker.git
cd task-tracker
```

Run the following command to build and run the project:

```bash
# Build the application in wsl
make build

# or

# Build the application in bash
go build -o task-tracker

# Run the application
./task-tracker --help # To see the list of available commands

# To add a task
./task-tracker add "Buy groceries"

# To update a task
./task-tracker update 1 "Buy groceries and cook dinner"

# To delete a task
./task-tracker delete 1

# To mark a task as in progress/done/todo
./task-tracker mark-in-progress 1
./task-tracker mark-done 1

# To list all tasks
./task-tracker list= 
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
```

[Up](#anchor)
