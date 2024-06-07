package model

import (
	"ginChart/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time `gorm:"column:login_time" json:"login_time"`
	HeartbeatTime time.Time `gorm:"column:heartbeat_time" json:"heartbeat_time"`
	LogOutTime    time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLoginOut    bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

/*
*
通过手机号查找用户
*/
func FindUserByPhone(phone string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("phone = ?", phone).First(&user)
	return user
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}
/**
获取用户列表
 */
func GetUserList(p int, ps int) ([]UserBasic, error) {
	var data []UserBasic
	offset := (p - 1) * ps
	result := utils.DB.Offset(offset).Limit(ps).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

/*
*
添加用户
*/
func CreateUser(user *UserBasic) error {
	result := utils.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/*
*
删除用户
*/
func DeleteUser(id int) error {
	result := utils.DB.Delete(&UserBasic{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/*
*
修改用户
*/
func UpdateUser(user *UserBasic) error {
	result := utils.DB.Updates(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
