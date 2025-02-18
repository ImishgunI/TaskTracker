package main

import (
	"TaskTracker/enums"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You should use command(add update delete list)")
	}
	filename := "tasks.json"
	var (
		tasks []Task
		id    int
	)
	command := os.Args[1]
	result := checkCommand(command)
	id, err := choseId(result)
	if err != nil {
		log.Fatal(err)
	}
	callFuncs(result, tasks, filename, id)
}

func callFuncs(result int, tasks []Task, filename string, id int) {
	if result == enums.Add {
		description := getDescriptionForAdd()
		addTask(tasks, filename, description)
	} else if result == enums.Delete {
		deleteTask(filename, id, tasks)
	} else if result == enums.Update {
		description := getDescriptionForUpdate()
		updateTask(id, tasks, filename, description)
	} else if result == enums.Mark_in_progress {
		markIP(id, tasks, filename)
	} else if result == enums.Mark_done {
		markDone(id, tasks, filename)
	} else if result == enums.ListAll {
		listAll(filename)
	} else if result == enums.ListTodo {
		listTodo(filename, tasks)
	} else if result == enums.ListInProgress {
		listInProgress(filename, tasks)
	} else if result == enums.ListDone {
		listDone(filename, tasks)
	}
}

/*** Pool of accessory functions ***/

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

func choseId(result int) (int, error) {
	if result == enums.Delete {
		return getId()
	} else if result == enums.Update {
		return getId()
	} else if result == enums.Mark_in_progress || result == enums.Mark_done {
		return getId()
	} else {
		return 0, nil
	}
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

func decode(tasks *[]Task, filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
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

func changeId(tasks []Task) []Task {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id != 1 {
			tasks[i].Id -= 1
		}
	}
	return tasks
}

/***  Pool of main functions  ***/

func addTask(tasks []Task, filename, desc string) {
	decode(&tasks, filename)
	t := time.Now()
	var task = Task{
		Id:          len(tasks) + 1,
		Description: desc,
		Status:      "todo",
		CreatedAt:   t.Format(time.RFC1123),
		UpdatedAt:   t.Format(time.RFC1123),
	}

	tasks = append(tasks, task)
	encode(tasks, filename)
}

func deleteTask(filename string, id int, tasks []Task) {
	decode(&tasks, filename)
	var i int
	if id > len(tasks) {
		log.Fatal("Task isn't exist")
	}
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
	t := time.Now()
	if id > len(tasks) {
		log.Fatal("Task isn't exist")
	}
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = t.Format(time.RFC1123)
			break
		}
	}

	encode(tasks, filename)
}

func markIP(id int, tasks []Task, filename string) {
	decode(&tasks, filename)
	if id > len(tasks) {
		log.Fatal("Task isn't exist")
	}
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
	if id > len(tasks) {
		log.Fatal("Task isn't exist")
	}
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			tasks[i].Status = enums.Done
			break
		}
	}
	encode(tasks, filename)
}

func listAll(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	info, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, info.Size())
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < count; i++ {
		fmt.Print(string(data[i]))
	}
}

func listTodo(filename string, tasks []Task) {
	decode(&tasks, filename)
	fmt.Print("[")
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Status == enums.Todo {
			fmt.Printf("\n%2c\n", '{')
			fmt.Printf("    \"id\": %d,\n    \"description\": \"%s\",\n    \"status\": \"%s\",\n    \"createdAt\": \"%s\",\n    \"updatedAt\": \"%s\"\n", tasks[i].Id,
				tasks[i].Description, tasks[i].Status,
				tasks[i].CreatedAt, tasks[i].UpdatedAt)
			fmt.Printf("%2c\n", '}')
		}
	}
	fmt.Println("]")
}

func listInProgress(filename string, tasks []Task) {
	decode(&tasks, filename)
	fmt.Print("[")
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Status == enums.Inprogress {
			fmt.Printf("\n%2c\n", '{')
			fmt.Printf("    \"id\": %d,\n    \"description\": \"%s\",\n    \"status\": \"%s\",\n    \"createdAt\": \"%s\",\n    \"updatedAt\": \"%s\"\n", tasks[i].Id,
				tasks[i].Description, tasks[i].Status,
				tasks[i].CreatedAt, tasks[i].UpdatedAt)
			fmt.Printf("%2c\n", '}')
		}
	}
	fmt.Println("]")
}

func listDone(filename string, tasks []Task) {
	decode(&tasks, filename)
	fmt.Print("[")
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Status == enums.Done {
			fmt.Printf("\n%2c\n", '{')
			fmt.Printf("    \"id\": %d,\n    \"description\": \"%s\",\n    \"status\": \"%s\",\n    \"createdAt\": \"%s\",\n    \"updatedAt\": \"%s\"\n", tasks[i].Id,
				tasks[i].Description, tasks[i].Status,
				tasks[i].CreatedAt, tasks[i].UpdatedAt)
			fmt.Printf("%2c\n", '}')
		}
	}
	fmt.Println("]")
}
