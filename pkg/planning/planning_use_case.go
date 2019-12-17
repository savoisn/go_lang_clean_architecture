package planning

import (
	"errors"

	"gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/entity"
)

type PlanningUseCase struct {
	persister PlanningPersister
}

func NewPlanningUseCase(p PlanningPersister) PlanningUseCase {
	return PlanningUseCase{persister: p}
}

func (p PlanningUseCase) SavePlanning(planning entity.Planning) (error, entity.Planning) {
	return p.persister.SavePlanning(planning)
}

func (p PlanningUseCase) GetPlannings() (error, []entity.Planning) {
	return p.persister.GetPlannings()
}

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
	isAround := (newTask.StartDate.Before(task.StartDate) && newTask.EndDate.After(task.EndDate))
	isInside := isAfter && isBefore
	isStartingAtSameTime := newTask.StartDate.Equal(task.StartDate)
	isEndingAtSameTime := newTask.EndDate.Equal(task.EndDate)
	// fmt.Printf("isAfter %t, isBefore %t, isAround %t \n", isAfter, isBefore, isAround)

	return isAfter || isBefore || isAround || isInside || isStartingAtSameTime || isEndingAtSameTime
}

func isOverLapping(tasks []entity.Task, newTask entity.Task) (err error) {
	for _, task := range tasks {
		if isTaskOverlapping(newTask, task) {
			err = errors.New("Ajout impossible : une tâche est déjà prévue à ce moment là")
		}
	}
	return err
}
