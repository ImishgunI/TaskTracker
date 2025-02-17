package enums

const (
	Add = iota + 1
	Delete
	Update
	Mark_in_progress
	Mark_done
	ListAll
	ListTodo
	ListInProgress
	ListDone
)

const (
	CommandAdd  = "add"
	CommandDel  = "delete"
	CommandUpd  = "update"
	CommandMIP  = "mark-in-progress"
	CommandMD   = "mark-done"
	CommandList = "list"
)

const (
	Done       = "done"
	Todo       = "todo"
	Inprogress = "in-progress"
)
