package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenbt456/todolist-go/model"
	"github.com/nguyenbt456/todolist-go/store"
	"github.com/nguyenbt456/todolist-go/util"
)

// Register new user
func Register(c *gin.Context) {
	user := &model.User{}
	c.ShouldBind(user)

	userStore := store.NewUserStore()
	user, err := userStore.Create(user.Name, user.Username, user.Password, user.Email)
	if err != nil {
		log.Println(err.Error())
		util.HandleError(c, 1001, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Login to app
func Login(c *gin.Context) {
	user := &model.User{}
	c.ShouldBind(user)

	userStore := store.NewUserStore()
	user, err := userStore.FindByUsernameAndPassword(user.Username, user.Password)
	if err != nil {
		log.Println(err.Error())
		util.HandleError(c, 1000, err)
		return
	}

	if user == nil {
		util.HandleError(c, 1001, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
