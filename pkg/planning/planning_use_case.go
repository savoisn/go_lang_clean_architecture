package planning_use_case

import (
	"errors"

	"gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/entity"
)

func addTaskToPlanning(planning entity.Planning, newTask entity.Task) (err error, ret entity.Planning) {
	err = isOverLapping(planning.Tasks, newTask)

	if err == nil {
		planning.Tasks = append(planning.Tasks, newTask)
	}

	return err, planning
}

func isTaskOverlapping(newTask entity.Task, task entity.Task) bool {
	isAfter := (newTask.StartDate.After(task.StartDate) && newTask.StartDate.Before(task.EndDate))
	isBefore := (newTask.EndDate.Before(task.EndDate) && newTask.EndDate.After(task.StartDate))
	isAround := false
	isInside := false
	return isAfter || isBefore || isAround || isInside
}

func isTaskAfterTask() {

}

func isOverLapping(tasks []entity.Task, newTask entity.Task) (err error) {
	for _, task := range tasks {
		if isTaskOverlapping(newTask, task) {
			err = errors.New("Ajout impossible : une tâche est déjà prévue à ce moment là")
		}
	}
	return err
}
