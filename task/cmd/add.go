package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/moficodes/gophercises/task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to our task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong : ", err)
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to your List\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
