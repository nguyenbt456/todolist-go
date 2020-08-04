package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/nguyenbt456/todolist-go/model"
	"github.com/nguyenbt456/todolist-go/util"

	"github.com/gin-gonic/gin"
	"github.com/nguyenbt456/todolist-go/store"
)

// GetTasks return tasks by query
func GetTasks(c *gin.Context) {
	dateQuery := c.Query("date")
	date := time.Now()

	if dateQuery != "" {
		dateInt, _ := strconv.Atoi(dateQuery)
		date = date.AddDate(0, 0, dateInt)
	}

	taskStore := store.NewTaskStore()
	tasks, err := taskStore.FindByDate(date)
	if err != nil {
		log.Println(err.Error())
		util.HandleError(c, 2000, err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// CreateTask create new task
func CreateTask(c *gin.Context) {
	userIDI, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", userIDI)

	task := &model.Task{}
	c.ShouldBind(task)

	taskStore := store.NewTaskStore()
	task, err := taskStore.Create(task.Name, model.TaskStatus.UnFinished, model.TaskType.Day, userID)
	if err != nil {
		log.Println(err.Error())
		util.HandleError(c, 2001, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask update status of task
func UpdateTask(c *gin.Context) {
	userIDI, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", userIDI)

	task := &model.Task{}
	c.ShouldBind(task)

	taskStore := store.NewTaskStore()
	err := taskStore.UpdateStatus(userID, task.ID, task.Status)
	if err != nil {
		log.Println(err.Error())
		util.HandleError(c, 2002, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})

}

// DeleteTask delete task in database
func DeleteTask(c *gin.Context) {
	userIDI, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", userIDI)

	task := &model.Task{}
	c.ShouldBind(task)

	taskStore := store.NewTaskStore()
	task, err := taskStore.FindByID(task.ID)
	if err != nil {
		log.Println(err.Error())
		util.HandleError(c, 2003, err)
		return
	}

	if task.UserID != userID {
		err = errors.New("User don't have permision")
		log.Println(err.Error())
		util.HandleError(c, 2004, err)
		return
	}

	err = taskStore.DeleteByID(task.ID)
	if err != nil {
		log.Println(err.Error())
		util.HandleError(c, 2005, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
