package cmd

type Task struct { // TODO: add dates after
	Name string `json:"name"`
	Type string `json:"type"`
}

func init() {
	rootCmd.AddCommand(addTask, listTask, deleteTask)
}
