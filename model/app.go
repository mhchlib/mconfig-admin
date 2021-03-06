package model

import (
	"time"
)

// App ...
type App struct {
	Id          int    `gorm:"primary_key,column:id"`
	Name        string `gorm:"column:app_name" `
	Key         string `gorm:"column:app_key"`
	Description string `gorm:"column:description"`
	CreateUser  int    `gorm:"column:create_user"`
	UpdateUser  int    `gorm:"column:update_user"`
	CreateTime  int64  `gorm:"column:create_time"`
	UpdateTime  int64  `gorm:"column:update_time"`
}

// TableName ...
func (App) TableName() string {
	return "m_app"
}

// InsertApp ...
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

// CheckAppKeyUnique ...
func CheckAppKeyUnique(key string) bool {
	first := db.Where(&App{Key: key}).First(&App{})
	if first.Error == nil {
		return false
	}
	return true
}

// ListApps ...
func ListApps(filter string, limit int, offset int) ([]*App, error) {
	apps := make([]*App, 0)
	fields := []string{"id", "app_name", "description", "app_key", "create_time", "update_time"}
	find := db.Select(fields).Where("app_name LIKE ? or description LIKE ?", "%"+filter+"%", "%"+filter+"%").Order("update_time desc").Limit(limit).Offset(offset).Find(&apps)
	if find.Error != nil {
		return nil, find.Error
	}
	return apps, nil
}

// DeleteApp ...
func DeleteApp(id int) error {
	app := &App{}
	app.Id = id
	d := db.Delete(app)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

// UpdateApp ...
func UpdateApp(id int, name string, desc string) error {
	app := &App{
		Id:          id,
		Name:        name,
		Description: desc,
		UpdateTime:  time.Now().Unix(),
	}
	update := db.Model(app).Update(app)
	if update.Error != nil {
		return update.Error
	}
	return nil
}

// GetApp ...
func GetApp(id int) (*App, error) {
	f := &App{
		Id: id,
	}
	data := db.Where("id = ?", f.Id).Find(f)
	if data.Error != nil {
		return nil, data.Error
	}
	return f, nil
}

// CountApp ...
func CountApp() (interface{}, error) {
	var count int
	c := db.Model(&App{}).Count(&count)
	if c.Error != nil {
		return nil, c.Error
	}
	return count, nil
}
