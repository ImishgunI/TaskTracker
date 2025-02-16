package main

import (
	"TaskTracker/enums"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You should use command(add update delete list)\n")
	}
	command := os.Args[1]
	result := checkCommand(command)
	fmt.Print(result)
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
