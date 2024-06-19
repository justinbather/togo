package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type ItemList struct {
	Tasks []Item `json:"itemList"`
}

type Item struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var fileDir string

func init() {
	// TODO: Check if folder and json storage file have been created
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

	if len(args) > 2 {
		log.Fatalln("Too many arguments. run `togo help` for a list of available commands")
	}

	switch args[0] {
	case "new":
		if len(args) != 2 {
			log.Fatalln("Incorrect call to `new`. Usage: `new <task>`")
		}
		newTask := Item{Task: args[1], Completed: false}

		fileData, err := ioutil.ReadFile(fileDir)
		if err != nil {
			log.Fatalln("Error reading stored tasks. Task not added. err: ", err)
		}

		var allTasks ItemList

		err = json.Unmarshal(fileData, &allTasks.Tasks)
		if err != nil {
			log.Fatalln("Error unmarshalling file data: ", err)
		}

		newTaskList := append(allTasks.Tasks, newTask)

		data, err := json.MarshalIndent(newTaskList, "", "\t")
		if err != nil {
			log.Fatalln("Error marshalling json: ", err)
		}

		err = ioutil.WriteFile(fileDir, data, 0755)
		if err != nil {
			log.Fatalln("Error writing to file: ", err)
		}

		log.Println("Task added successfully")
		log.Println("Remaining Tasks\n", newTaskList)

	case "done":

	case "clean":

	case "del":

	case "clear":
	case "help":
		log.Println("\n togo usage:\n `add <task>` - adds a new task to your list\n `done <index> - sets the task at the given integer index to complete\n `clean` - clears all completed tasks from your list\n `clear` - removes all tasks from your list\n `del <index> - removes a task at the given integer index")

	default:
		log.Println("Incorrect usage.\n usage:\n `add <task>` - adds a new task to your list\n `done <index> - sets the task at the given integer index to complete\n `clean` - clears all completed tasks from your list\n `clear` - removes all tasks from your list\n `del <index> - removes a task at the given integer index")

	}

}
