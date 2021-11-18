## TodoApp - CLI task Manager
The TodoApp is a CLI Todo manager. We are going to create a Command(CLI) 'todo', which can be used to track Todo items.
We can Add, Delete and list Todo tasks.

We will use the following modules
- Cobra - For command line operations.
- BoltDB - Storing the tasks.


### Notes:
- Make sure to use the same name for the project as your cobra command. (Here I am creating a CLI command 'todo', So I named my project as 'todo')

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
