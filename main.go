package main

import (
	"log"
	"os"
	"strconv"
)

type Item struct {
	task      string
	completed bool
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("Not enough arguments. run `togo help` for list of available commands")
	}

	if len(args) > 2 {
		log.Fatalln("Too many arguments. run `togo help` for a list of available commands")
	}

	taskList := []Item{{task: "Hello 1", completed: true}, {task: "Hello 2", completed: false}}

	switch args[0] {
	case "new":
		if len(args) != 2 {
			log.Fatalln("Incorrect call to `new`. Usage: `new <task>`")
		}
		task := Item{task: args[1], completed: false}
		taskList = append(taskList, task)
		log.Println("Saved tasks: ", taskList)

	case "done":
		if len(args) != 2 {
			log.Fatalln("Incorrect call to `done`. Usage: `done <index>`")
		}
		idx, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalln("Not a valid integer index")
		}

		if idx > len(taskList) {
			log.Fatalln("Provided index is out of range of stored tasks.")
		}

		taskList[idx].completed = true
		log.Printf("Completed task number %d. tasks:\n%v", idx, taskList)

	case "clean":
		count := 0
		for idx, task := range taskList {
			if task.completed == true {
				taskList = append(taskList[:idx], taskList[idx+1:]...)
				count++
			}
		}
		log.Printf("Cleaned %d tasks.", count)

	case "del":
		if len(args) != 2 {
			log.Fatalln("Incorrect call to `del`. Usage: `del <index>`")
		}
		idx, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalln("Not a valid integer index")
		}

		if idx > len(taskList) {
			log.Fatalln("Provided index is out of range of stored tasks.")
		}

		taskList = append(taskList[:idx], taskList[idx+1:]...)
		log.Printf("Removed task number %d. Remaining tasks:\n%v", idx, taskList)

	case "clear":
		taskList = []Item{}
		log.Println("Cleared all tasks.")
	case "help":
		log.Println("\n togo usage:\n `add <task>` - adds a new task to your list\n `done <index> - sets the task at the given integer index to complete\n `clean` - clears all completed tasks from your list\n `clear` - removes all tasks from your list\n `del <index> - removes a task at the given integer index")

	default:
		log.Println("Incorrect usage.\n usage:\n `add <task>` - adds a new task to your list\n `done <index> - sets the task at the given integer index to complete\n `clean` - clears all completed tasks from your list\n `clear` - removes all tasks from your list\n `del <index> - removes a task at the given integer index")

	}

}
