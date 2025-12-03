package utils

import (
	"encoding/json"
	"os"
	"time"
	"tt/enums"
	"tt/structures"
)

var jsonFilePath = ".tt/tasks.json"

func ReadJsonFile() []structures.FormatStored {
	var jsonTask []structures.FormatStored
	file, _ := os.ReadFile(jsonFilePath)
	json.Unmarshal(file, &jsonTask)
	return jsonTask
}

func AddNewTask(task string) {
	newTask := structures.FormatStored{}
	newTask.Task = task
	newTask.Status = enums.ToDo.ToString()
	newTask.CreatedAt = time.Now()
	newTask.Id = len(ReadJsonFile()) + 1
	appendJsonFile(newTask)
}

func appendJsonFile(newTask structures.FormatStored) {
	existingTasks := ReadJsonFile()
	existingTasks = append(existingTasks, newTask)
	OverwriteJsonFile(existingTasks)
}

func OverwriteJsonFile(tasks []structures.FormatStored) {
	jsonString, _ := json.Marshal(tasks)
	os.WriteFile(jsonFilePath, jsonString, os.ModePerm)
}
