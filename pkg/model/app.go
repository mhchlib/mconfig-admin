package model

import "time"

type App struct {
	Id          int    `gorm:"primary_key,column:id"`
	Name        string `gorm:"column:app_name"`
	Key         string `gorm:"column:app_key"`
	Description string `gorm:"column:description"`
	Filter      int    `gorm:"column:filter_id"`
	CreateUser  int    `gorm:"column:create_user"`
	UpdateUser  int    `gorm:"column:update_user"`
	CreateTime  int64  `gorm:"column:create_time"`
	UpdateTime  int64  `gorm:"column:update_time"`
}

func (App) TableName() string {
	return "m_app"
}

func InsertApp(name, desc, key string) error {
	create := db.Create(&App{
		Name:        name,
		Key:         key,
		Description: desc,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	})
	return create.Error
}

func CheckKeyUnique(key string) bool {
	first := db.Where(&App{Key: key}).First(&App{})
	if first.Error == nil {
		return false
	}
	return true
}

func ListApps(filter string, limit int, offset int) ([]*App, error) {
	apps := make([]*App, 0)
	fields := []string{"app_name", "description", "app_key", "create_time", "update_time"}
	find := db.Select(fields).Where("app_name LIKE ? or description LIKE ?", "%"+filter+"%", "%"+filter+"%").Limit(limit).Offset(offset).Find(&apps)
	if find.Error != nil {
		return nil, find.Error
	}
	return apps, nil
}
