package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteTask = &cobra.Command{
	Use:   "del",
	Short: "Delete task command",
	Run: func(cmd *cobra.Command, args []string) {


		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to get user's home directory:", err)
			return
		}

		tasklistDir := homeDir + "/.tasklist"
		if _, err := os.Stat(tasklistDir); os.IsNotExist(err) {
			if err := os.Mkdir(tasklistDir, 0700); err != nil {
				fmt.Println("Failed to create .tasklist directory:", err)
				return
			}
		}

		taskFile := tasklistDir + "/tasks.json"
		//file, err := os.OpenFile(taskFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		file, err := os.ReadFile(taskFile)
		if err != nil {
			fmt.Println("Failed to open tasks file:", err)
			return
		}
		//defer file.Close()

		var tasks []Task

		if err := json.Unmarshal(file, &tasks); err != nil {
			fmt.Println(err)
			fmt.Println("Failed to Unmarshal json file")
			panic(err)
		}

		for index, value := range tasks {
			for _, task := range args {
				if value.Name == task {
					tasks = append(tasks[:index], tasks[index+1:]...)
				}
			}
		}
		taskJSON, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			fmt.Println("Failed to serialize tasks to JSON:", err)
			return
		}

		if err := os.WriteFile(taskFile, taskJSON, 0644); err != nil {
			fmt.Println("Failed to write tasks to file:", err)
			return
		}

		fmt.Println("Tasks removed successfully.")
	},
}
