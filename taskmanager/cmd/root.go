package cmd

import (
	"fmt"
	"log"
	"os"
	"shaowei/gophercises_with_tests/taskmanager/db"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	//Add all the commands before program starts
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(doCmd)
}

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage todos",
	Long:  "A CLI to manage your todos",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here to manage your TODOs.")
	},
}

//Add Todo
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		//Join the args into a String
		taskName := strings.Join(args, " ")
		//Create and add it into the db
		id, err := db.CreateTask(taskName)
		if err != nil {
			fmt.Println(err)
		}
		//Print info that it has been added
		fmt.Printf(`Added "%s" to your task list, the id is %d`, taskName, id)
	},
}

//List all todos
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your tasks in your todolist.",
	Run: func(cmd *cobra.Command, args []string) {
		//Get all the tasks from db
		tasks, err := db.ListAllTasks()
		if err != nil {
			fmt.Println(err)
		}
		for _, task := range tasks {
			fmt.Printf("%d : %v\n", task.Key, task.Value)
		}
	},
}

//Mark a todo as complete
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a todo on your list as complete.",
	Run: func(cmd *cobra.Command, args []string) {
		//convert args[0] to string
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("given ID is not a number")
		}
		//Get the full task
		task, err := db.GetTaskByID(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = db.DeleteTaskByID(id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(`Marked "%s" as done.`, task.Value)
	},
}

//Execute the root cmd
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
