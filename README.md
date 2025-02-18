# TaskTracker

This is a simple program for keeping a task list. With it you can add, delete, update, change task status and output both simple and by status.

To start using it you need to have Golang installed.

There are a few steps you need to take to start the task:

```
make 

./task-cli <commands>
```

Available commands:

```
"To add new task"
task-cli add "Task"

"To delete the task"
task-cli delete <id>

"To update the task"
task-cli update <id> "Task"

"To change status"
task-cli mark-in-progress <id>
task-cli mark-done <id>

"To output all tasks"
task-cli list

"To output tasks by status"
task-cli list todo
task-cli list is-progress
task-cli list done
```

Link to task on roadmap.sh:
https://roadmap.sh/projects/task-tracker
