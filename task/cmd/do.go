package cmd

import (
	"fmt"
	"strconv"

	"github.com/moficodes/gophercises/task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse arg", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something wen wrong : ", err)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as complete. Error: %s\n", id, err)
			}
			fmt.Printf("Task %s marked complete\n", task.Value)
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
