package planning_use_case

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/entity"
)

func TestCreationAgent(t *testing.T) {
	// var agent entity.Agent = entity.NewAgent("Nicolas")
	agent := entity.Agent{Name: "Nicolas"}
	fmt.Printf("%+v\n", agent)
	assert.Equal(t, "Nicolas", agent.Name)
}

func TestPlanningCreation(t *testing.T) {
	agent := entity.Agent{Name: "Diane"}
	fmt.Printf("%+v\n", agent)
	startDate := time.Now()
	fmt.Printf("%s\n", startDate)
	endDate := startDate.AddDate(0, 0, 7)
	fmt.Printf("%+s\n", endDate)
	planning := entity.Planning{Agent: agent, StartDate: startDate, EndDate: endDate}

	assert.Equal(t, "Diane", planning.Agent.Name)
	assert.Equal(t, startDate, planning.StartDate)
	assert.Equal(t, endDate, planning.EndDate)
}

func TestAddTaskToPlanning(t *testing.T) {

	agent := entity.Agent{Name: "Mathilde"}
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 7)
	planning := entity.Planning{Agent: agent, StartDate: startDate, EndDate: endDate}

	taskStartDate := time.Now()
	taskEndDate := startDate.AddDate(0, 0, 1)
	taskType := "Manuelle"

	task := entity.Task{
		Description: "sortir les poubelles",
		StartDate:   taskStartDate,
		EndDate:     taskEndDate,
		Type:        taskType}

	assert.Equal(t, 0, len(planning.Tasks))
	err, planning := addTaskToPlanning(planning, task)

	assert.Nil(t, err)

	assert.Equal(t, 1, len(planning.Tasks))
	assert.Equal(t, task, planning.Tasks[0])
}

func TestAddTaskToPlanningThatStartDateOverlapsPlanningTask(t *testing.T) {

	agent := entity.Agent{Name: "Mathilde"}
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 7)
	planning := entity.Planning{Agent: agent, StartDate: startDate, EndDate: endDate}

	taskType := "Manuelle"

	task1 := entity.Task{
		Description: "sortir les poubelles",
		StartDate:   startDate.AddDate(0, 0, 1),
		EndDate:     startDate.AddDate(0, 0, 4),
		Type:        taskType}
	task2 := entity.Task{
		Description: "tondre la pelouse",
		StartDate:   startDate.AddDate(0, 0, 2),
		EndDate:     startDate.AddDate(0, 0, 3),
		Type:        taskType}

	assert.Equal(t, 0, len(planning.Tasks))
	err, planning := addTaskToPlanning(planning, task1)

	assert.Nil(t, err)

	err, planning = addTaskToPlanning(planning, task2)

	assert.NotNil(t, err)

	assert.Equal(t, 1, len(planning.Tasks))
	assert.Equal(t, task1, planning.Tasks[0])
}

func TestAddTaskToPlanningThatEndDateOverlapsPlanningTask(t *testing.T) {

	agent := entity.Agent{Name: "Éloïse"}
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 7)
	planning := entity.Planning{Agent: agent, StartDate: startDate, EndDate: endDate}

	taskType := "Manuelle"

	task1 := entity.Task{
		Description: "sortir les poubelles",
		StartDate:   startDate.AddDate(0, 0, 2),
		EndDate:     startDate.AddDate(0, 0, 4),
		Type:        taskType}
	task2 := entity.Task{
		Description: "tondre la pelouse",
		StartDate:   startDate.AddDate(0, 0, 1),
		EndDate:     startDate.AddDate(0, 0, 3),
		Type:        taskType}

	assert.Equal(t, 0, len(planning.Tasks))
	err, planning := addTaskToPlanning(planning, task1)

	assert.Nil(t, err)

	err, planning = addTaskToPlanning(planning, task2)

	assert.NotNil(t, err)

	assert.Equal(t, 1, len(planning.Tasks))
	assert.Equal(t, task1, planning.Tasks[0])
}
