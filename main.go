package main

import (
	"TaskTracker/enums"
	"encoding/json"
	"log"
	"os"
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
	var description string
	command := os.Args[1]
	result := checkCommand(command)
	if result == enums.Add {
		description = getDescription()
		addTask(filename, description)
	}
}

func getDescription() string {
	if len(os.Args) < 3 {
		log.Fatal("You must write a task description")
	}
	return os.Args[2]
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

func addTask(filename, desc string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
	var tasks []Task
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
		err = decoder.Decode(&tasks)
		if err != nil {
			log.Fatal(err)
		}
	}
	var task = Task{
		Id:          len(tasks) + 1,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)

	file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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
