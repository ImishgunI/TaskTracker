package main

import (
	"fmt"
	"log"
	"os"
)

const (
	add              = 1
	delete           = 2
	update           = 3
	mark_in_progress = 4
	mark_done        = 5
	listAll          = 6
	listTodo         = 7
	listInProgress   = 8
	listDone         = 9
)

const (
	commandAdd  = "add"
	commandDel  = "delete"
	commandUpd  = "update"
	commandMIP  = "mark-in-progress"
	commandMD   = "mark-done"
	commandList = "list"
)

const (
	done       = "done"
	todo       = "todo"
	inprogress = "in-progress"
)

func main() {
	command := os.Args[1]
	result := checkCommand(command)
	fmt.Print(result)
}

func checkCommand(command string) int {
	var result int
	switch command {
	case commandAdd:
		result = add
	case commandDel:
		result = delete
	case commandUpd:
		result = update
	case commandMIP:
		result = mark_in_progress
	case commandMD:
		result = mark_done
	case commandList:
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
		case done:
			result = listDone
		case todo:
			result = listTodo
		case inprogress:
			result = listInProgress
		default:
			log.Fatal("Unknown list command")
		}
	} else {
		result = listAll
	}
	return result
}
