package router

import (
	"ginChart/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/getUserList",service.GetUserList)
	r.POST("/createUser",service.CreateUser)
	r.POST("/deleteUser",service.DeleteUser)
	r.POST("/updateUser",service.UpdateUser)
	return r
}
