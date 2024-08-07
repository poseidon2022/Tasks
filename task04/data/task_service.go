package data

import (
	"errors"
	task "task04/models"
)

func SearchByID(id string) (task.Task, error) {
	taskCollection := task.Pre
	for _, value := range taskCollection {
		if value.ID == id {
			return value, nil
		}
	}
	return task.Task{}, errors.New("task not found") 
}

func DeleteByID(id string) (error) {
	taskCollection := task.Pre
	for index, value := range taskCollection {
		if value.ID == id {
			task.Pre = append(task.Pre[:index], task.Pre[index+1:]...)
			return nil
		}
	}

	return errors.New("book not found")
}

func ModifyTask(modified task.Task, id string) error {
	taskCollection := task.Pre
	for index, value := range taskCollection {
		if value.ID == id {
			task.Pre[index].Description = modified.Description
			task.Pre[index].Title = modified.Title
			return nil
		}
	}

	return errors.New("task with the specified ID not found")
}

func AddTask(newTask task.Task) error {
	taskCollection := task.Pre
	for _, value := range taskCollection {
		if value.ID == newTask.ID {
			return errors.New("task with the same ID exists")
		}
	}
	task.Pre = append(task.Pre, newTask)
	return nil 
}