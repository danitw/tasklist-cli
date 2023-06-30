package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

type Task struct { // TODO: add dates after
	Name string `json:"name"`
	Type string `json:"type"`
}

var addTask = &cobra.Command{
	Use:   "add",
	Short: "Add task command",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := args

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
		file, err := os.OpenFile(taskFile, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Failed to open tasks file:", err)
			return
		}
		defer file.Close()

		decoder := json.NewDecoder(file)

		var taskList []Task

		decoder.Decode(&taskList)

		if err := decoder.Decode(&taskList); err != nil && err != io.EOF {
			fmt.Println("Failed to Unmarshal/Decode json file")
			panic(err)
		}

		for _, task := range tasks {
			taskObject := Task{
				Name: task,
				Type: "pending",
			}
			taskList = append(taskList, taskObject)
		}

		taskJSON, err := json.MarshalIndent(taskList, "", "  ")
		if err != nil {
			fmt.Println("Failed to serialize tasks to JSON:", err)
			return
		}

		if err := ioutil.WriteFile(taskFile,taskJSON , 0644); err != nil {
			fmt.Println("Failed to write tasks to file:", err)
			return
		}

		fmt.Println("Tasks added successfully.")
	},
}

func init() {
	rootCmd.AddCommand(addTask)
}
