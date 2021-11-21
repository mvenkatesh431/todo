## TodoApp - CLI task Manager
The TodoApp is a CLI Todo manager. We are going to create a Command(CLI) 'todo', which can be used to track Todo items.
We can Add, Delete and list Todo tasks.

We will use the following modules
- Cobra - For command line operations.
- BoltDB - Storing the tasks.


### Notes:
- Make sure to use the same name for the project as your cobra command. (Here I am creating a CLI command 'todo', So I named my project as 'todo')
- This is my practise version of one of the exercises of gophercises

### Cobra setup

```
# Get the cobra from github
$ go get -u github.com/spf13/cobra/cobra

# Initialize your module
$ go mod init "your module Name"

# initialize the cobra application ('todo' is the CLI command here)
$ cobra init --pkg-name todo

# if needed run the 'go mod tidy'
$ go mod tidy

# add the necessary code, And you can install your CLI command with
$ go install

# Now your command is available, Check it
$ todo -h
todo - A CLI Command Line Todo task Manager

```

### BoltDB setup:

As per the bolt github page, It is not activly maintained so we are going to use the clone of bolt db - https://github.com/etcd-io/bbolt

```
# use the following command to get the bbolt db
$ go get go.etcd.io/bbolt/...

# importing
import	bolt "go.etcd.io/bbolt"
```

### Example Usage:

First of all install the command using the 'go install', Then 'todo' command is ready to use. 
You can add, list and delete todos. Please look at the below example for usage.

```
todo> go install .
todo> 
todo> todo -h
  todo [command]

  completion  generate the autocompletion script for the specified shell
  list        To List the tasks in the Todo List
todo> go install .
todo - A CLI Command Line Todo task Manager
todo>
todo - A CLI Command Line Todo task Manager
  todo [flags]
Available Commands:
  add         To Add a task to the Todo List
  completion  generate the autocompletion script for the specified shell
  list        To List the tasks in the Todo List
  -h, --help   help for todo

todo> todo list
todo>
todo> todo add "Commit this code"
Todo 'Commit this code' Added successfully to TodoList
Todo 'Drive to survive' Added successfully to TodoList
Todo 'Golang is Awesome' Added successfully to TodoList
todo>
2. Drive to survive
3. Golang is Awesome
todo>
todo>
2. Golang is Awesome
Todo 'Who is going to be World Champion?' Added successfully to TodoList
1. Commit this code
3. Who is going to be World Champion?
todo>
todo> todo delete 1
Todo 'Commit this code' marked as Completed
todo>
todo> todo list
1. Golang is Awesome
2. Who is going to be World Champion?
todo>
todo> todo delete 2
Todo 'Who is going to be World Champion?' marked as Completed
todo>
todo> todo list
1. Golang is Awesome
todo>
todo> todo delete 1
Todo 'Golang is Awesome' marked as Completed
todo>
todo> todo list
Your TodoList is Empty!. Add few todos :)
todo>
todo> 
```