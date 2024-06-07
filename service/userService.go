package service

import (
	"fmt"
	"ginChart/model"
	"ginChart/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

/*
*
获取用户列表
*/
func GetUserList(c *gin.Context) {
	page := c.DefaultQuery("page","1")
	pageSize := c.DefaultQuery("pageSize","20")
	currectPage,_ := strconv.Atoi(page)
	currectPageSize,_ := strconv.Atoi(pageSize)
	list, err := model.GetUserList(currectPage,currectPageSize)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "获取用户失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": list,
	})
}

/*
*
新增用户
*/
func CreateUser(c *gin.Context) {
	user := model.UserBasic{}
	username := c.DefaultPostForm("name", " ")
	password := c.DefaultPostForm("password", " ")
	repassword := c.DefaultPostForm("repassword", " ")
	phone := c.DefaultPostForm("phone", "")
	if password != repassword {
		c.JSON(500, gin.H{"message": "两次密码不一致"})
		return
	}
	userPhone := model.FindUserByPhone(phone)
	if userPhone.Phone != "" {
		c.JSON(500, gin.H{"message": "手机号已注册"})
		return
	}
	userName := model.FindUserByName(username)
	if userName.Name != "" {
		c.JSON(500, gin.H{"message": "用户名已注册"})
		return
	}
	//加密
	md5_pass_word := utils.MD5Encode(password)
	user.Name = username
	user.PassWord = md5_pass_word
	user.Phone = phone
	user.LoginTime = time.Now()
	user.HeartbeatTime = time.Now()
	user.LogOutTime = time.Now()
	resutl := model.CreateUser(&user)
	if resutl != nil {
		c.JSON(200, gin.H{"message": "新增用户失败"})
		return
	}
	c.JSON(200, gin.H{"message": "新增用户成功"})
}

/*
*
删除用户
*/
func DeleteUser(c *gin.Context) {
	userId := c.PostForm("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(500, gin.H{"message": "用户Id无效"})
		return
	}
	delerr := model.DeleteUser(id)
	if delerr != nil {
		c.JSON(500, gin.H{"message": "删除失败"})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功"})
}

/*
*
修改用户
*/
func UpdateUser(c *gin.Context) {
	user := model.UserBasic{}
	userid := c.PostForm("userid")
	username := c.PostForm("username")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"message": "参数不匹配"})
		return
	}
	ids, err := strconv.Atoi(userid)
	if err != nil {
		c.JSON(500, gin.H{"message": "用户Id无效"})
		return
	}
	user.ID = uint(ids)
	user.Name = username
	updateErr := model.UpdateUser(&user)
	if updateErr != nil {
		c.JSON(500, gin.H{"message": "修改失败"})
		return
	}
	c.JSON(200, gin.H{"message": "修改成功"})
}
