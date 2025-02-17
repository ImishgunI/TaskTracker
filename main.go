package main

import (
	"TaskTracker/enums"
	"encoding/json"
	"errors"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You should use command(add update delete list)")
	}
	filename := "tasks.json"
	var (
		description string
		tasks       []Task
	)
	command := os.Args[1]
	result := checkCommand(command)
	if result == enums.Add {
		description = getDescriptionForAdd()
		addTask(tasks, filename, description)
	} else if result == enums.Delete {
		id, err := getId()
		if err != nil {
			log.Fatal("Need to write a number")
		}
		deleteTask(filename, id, tasks)
	} else if result == enums.Update {
		id, err := getId()
		if err != nil {
			log.Fatal(err)
		}
		description = getDescriptionForUpdate()
		updateTask(id, tasks, filename, description)
	} else if result == enums.Mark_in_progress {
		id, err := getId()
		if err != nil {
			log.Fatal(err)
		}
		markIP(id, tasks, filename)
	} else if result == enums.Mark_done {
		id, err := getId()
		if err != nil {
			log.Fatal(err)
		}
		markDone(id, tasks, filename)
	}
}

func getDescriptionForAdd() string {
	if len(os.Args) < 3 {
		log.Fatal("You must write a task description")
	}
	return os.Args[2]
}

func getDescriptionForUpdate() string {
	if len(os.Args) < 4 {
		log.Fatal("You must write a task description")
	}
	return os.Args[3]
}

func getId() (int, error) {
	if len(os.Args) < 3 {
		log.Fatal("You must write a task id")
	}
	return strconv.Atoi(os.Args[2])
}

func checkCommand(command string) int {
	var result int
	switch command {
	case enums.CommandAdd:
		result = enums.Add
	case enums.CommandDel:
		result = enums.Delete
	case enums.CommandUpd:
		result = enums.Update
	case enums.CommandMIP:
		result = enums.Mark_in_progress
	case enums.CommandMD:
		result = enums.Mark_done
	case enums.CommandList:
		result = checkCommandList()
	default:
		log.Fatal("Unknown command")
	}
	return result
}
func checkCommandList() int {
	var result int
	if len(os.Args) > 2 {
		switch os.Args[2] {
		case enums.Done:
			result = enums.ListDone
		case enums.Todo:
			result = enums.ListTodo
		case enums.Inprogress:
			result = enums.ListInProgress
		default:
			log.Fatal("Unknown list command")
		}
	} else {
		result = enums.ListAll
	}
	return result
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func decode(tasks *[]Task, filename string) {
	var flag int
	if fileExist(filename) {
		flag = os.O_CREATE | os.O_RDONLY
	} else {
		flag = os.O_RDONLY
	}
	file, err := os.OpenFile(filename, flag, 0644)
	if err != nil && err.Error() != "EOF" {
		log.Fatal(err)
	}
	defer file.Close()

	info, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	if info.Size() > 0 {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(tasks)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func encode(tasks []Task, filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(tasks)
	if err != nil {
		log.Fatal(err)
	}
}

func addTask(tasks []Task, filename, desc string) {
	decode(&tasks, filename)
	var task = Task{
		Id:          len(tasks) + 1,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)
	encode(tasks, filename)
}

func changeId(tasks []Task) []Task {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id != 1 {
			tasks[i].Id -= 1
		}
	}
	return tasks
}

func deleteTask(filename string, id int, tasks []Task) {
	decode(&tasks, filename)
	var i int
	for i = 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			break
		}
	}
	tasks = slices.Delete(tasks, i, i+1)
	tasks = changeId(tasks)
	encode(tasks, filename)
}

func updateTask(id int, tasks []Task, filename, description string) {
	decode(&tasks, filename)
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			tasks[i].Description = description
			break
		}
	}

	encode(tasks, filename)
}

func markIP(id int, tasks []Task, filename string) {
	decode(&tasks, filename)
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			tasks[i].Status = enums.Inprogress
			break
		}
	}
	encode(tasks, filename)
}

func markDone(id int, tasks []Task, filename string) {
	decode(&tasks, filename)
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			tasks[i].Status = enums.Done
			break
		}
	}
	encode(tasks, filename)
}
