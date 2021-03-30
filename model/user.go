package model

import (
	"github.com/mhchlib/mconfig-admin/pkg/tools"
	"github.com/spf13/viper"
	"time"
)

// User ...
type User struct {
	Id         int    `gorm:"primary_key;AUTO_INCREMENT"  json:"id"`
	Name       string `gorm:"type:varchar(128)" json:"name"`
	Salt       string `gorm:"type:varchar(255)" json:"salt"`
	Password   string `gorm:"type:varchar(255)" json:"password"`
	CreateUser int    `gorm:"column:create_user" json:"create_user"`
	UpdateUser int    `gorm:"column:update_user" json:"update_user"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
}

// TableName ...
func (User) TableName() string {
	return "m_user"
}

// UserPayload ...
type UserPayload struct {
	UserId int
}

// InsertUser ...
func InsertUser(name, passwd string) error {
	salt := viper.GetString("user.salt")
	create := db.Create(&User{
		Name:       name,
		Password:   tools.Md5Crypt(passwd, salt),
		Salt:       salt,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	})
	return create.Error
}

// GetUserByName ...
func GetUserByName(name string) (*User, error) {
	f := &User{
		Name: name,
	}
	data := db.Where("name = ?", f.Name).Find(f)
	if data.Error != nil {
		return nil, data.Error
	}
	return f, nil
}

// CountUser ...
func CountUser() (interface{}, error) {
	var count int
	c := db.Model(&User{}).Count(&count)
	if c.Error != nil {
		return nil, c.Error
	}
	return count, nil
}
