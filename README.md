
#### Usage

`togo <options>`

Initial start creates a .togo.json file at ~/togo/ to store tasks. 

###### Future feature
- Running `togo` without any flags or parameters launches interactive session to use all below features

##### Commands

`new <task> --p <int>`
Creates a new task with optional priority level

`done <index>`
Marks task at `<index>` as complete

`clean`
Cleans up completed tasks

`del <index>`
Deletes task at `<index>`

`clear` 
Clears all tasks

###### Future commands
- Support integration with git to store tasks in a repo for use on multiple computers

