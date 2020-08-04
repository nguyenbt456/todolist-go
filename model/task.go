package model

import "time"

// Task represents tasks table
type Task struct {
	ID     string         `json:"id"`
	Name   string         `json:"name"`
	Status TaskStatusType `json:"status"`
	Type   TaskTypeType   `json:"type"`
	UserID string         `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TaskStatusType ...
type TaskStatusType int

type taskStatusValue struct {
	UnFinished TaskStatusType
	Finished   TaskStatusType
}

// TaskStatus is real value of status in task table
var TaskStatus = taskStatusValue{
	UnFinished: 0,
	Finished:   1,
}

// TaskTypeType ...
type TaskTypeType int

type taskTypeValue struct {
	Day  TaskTypeType
	Week TaskTypeType
}

// TaskType is value of type in task table
var TaskType = taskTypeValue{
	Day:  0,
	Week: 1,
}
