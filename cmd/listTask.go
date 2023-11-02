package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listTask = &cobra.Command{
	Use:   "list",
	Short: "List tasks command",
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

		fmt.Println(tasks)

	},
}
