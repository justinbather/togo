package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Item struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var fileDir string

func validateArgs(args []string, expectedLen int, cmd string) {
	if len(args) != expectedLen {
		log.Fatalf("Incorrect call to `%s`. See help for usage", cmd)
	}
}

func print(tasks []Item) {
	width := 0
	for _, task := range tasks {
		if len(task.Task) > width {
			width = len(task.Task)
		}
	}

	for range width {
		fmt.Print("-")
	}
	fmt.Print("-------")
	fmt.Print("\n")
	for _, item := range tasks {
		if item.Completed == true {
			fmt.Printf("|- \u2705 %s\n", item.Task)
		} else {
			fmt.Printf("|- [] %s\n", item.Task)
		}
	}

	for range width {
		fmt.Print("-")
	}

	fmt.Print("-------")
}

func persist(fp string, data any) {
	res, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Fatalln("Error marshalling json: ", err)
	}

	err = os.WriteFile(fp, res, 0755)
	if err != nil {
		log.Fatalln("Error writing to file: ", err)
	}
}

func fetch(fp string) []Item {
	fileData, err := os.ReadFile(fp)
	if err != nil {
		log.Fatalln("Error reading stored tasks. Task not added. err: ", err)
	}

	var allTasks []Item

	err = json.Unmarshal(fileData, &allTasks)
	if err != nil {
		log.Fatalln("Error unmarshalling file data: ", err)
	}

	return allTasks
}

// given ["new", "task", "hello"]
// return ["new", "task hello"]
func parseStringArgs(args []string) []string {
	joined := strings.Join(args[1:], " ")

	return []string{args[0], joined}
}

func parseIntArgs(args []string) int {
	i, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalln("Must provide a valid integer value")
	}
	return i

}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Error finding user home: ", err)
	}

	err = os.Chdir(home)
	if err != nil {
		log.Fatalln("Error changing directory to home: ", err)
	}

	if _, err := os.Stat(filepath.Join(home, "togo", ".togo.json")); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(filepath.Join(home, "togo"), 0755)
		if err != nil {
			log.Fatalln("Error creating directory: ", err)
		}
		log.Println("Folder created successfully")

		_, err := os.Create(filepath.Join(home, "togo", ".togo.json"))
		if err != nil {
			log.Fatalln("Error creating togo.json: ", err)
		}
		log.Println("File created successfully")
	}

	fileDir = filepath.Join(home, "togo", ".togo.json")
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("Not enough arguments. run `togo help` for list of available commands")
	}

	switch args[0] {
	case "new":
		args = parseStringArgs(args)
		validateArgs(args, 2, "new")

		tasks := fetch(fileDir)
		newTask := Item{Task: args[1], Completed: false}
		tasks = append(tasks, newTask)

		persist(fileDir, tasks)
		return

	case "done":
		validateArgs(args, 2, "done")
		idx := parseIntArgs(args)

		taskList := fetch(fileDir)
		if idx > len(taskList) {
			log.Fatalln("Index out of range")
		}

		taskList[idx].Completed = true
		persist(fileDir, taskList)
		print(taskList)
		return

	case "clean":
		dirty := fetch(fileDir)
		cleaned := []Item{}
		numCleaned := 0
		for _, task := range dirty {
			if task.Completed == false {
				cleaned = append(cleaned, task)
			} else {
				numCleaned++
			}
		}

		persist(fileDir, cleaned)
		fmt.Printf("Cleaned %d tasks.", numCleaned)
		return

	case "list":
		validateArgs(args, 1, "list")
		tasks := fetch(fileDir)
		print(tasks)
		return

	case "del":
		return

	case "clear":
		persist(fileDir, []Item{})
		log.Println("Cleared all tasks")
		return

	case "help":
		log.Println("\n togo usage:\n `add <task>` - adds a new task to your list\n `done <index> - sets the task at the given integer index to complete\n `clean` - clears all completed tasks from your list\n `clear` - removes all tasks from your list\n `del <index> - removes a task at the given integer index")
		return

	default:
		log.Println("Incorrect usage.\n usage:\n `add <task>` - adds a new task to your list\n `done <index> - sets the task at the given integer index to complete\n `clean` - clears all completed tasks from your list\n `clear` - removes all tasks from your list\n `del <index> - removes a task at the given integer index")
		return

	}

}
