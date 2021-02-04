package model

import (
	"time"
)

type TYPE_FILTER int

const (
	TYPE_FILTER_LUA TYPE_FILTER = 1
)

type Filter struct {
	Id         int         `gorm:"primary_key,column:id"`
	Type       TYPE_FILTER `gorm:"column:type"`
	Filter     string      `gorm:"column:filter"`
	CreateUser int         `gorm:"column:create_user"`
	UpdateUser int         `gorm:"column:update_user"`
	CreateTime int64       `gorm:"column:create_time"`
	UpdateTime int64       `gorm:"column:update_time"`
}

func (Filter) TableName() string {
	return "m_filter"
}

func InsertFilter(filterType TYPE_FILTER, filter string) (int, error) {
	f := &Filter{
		Type:       filterType,
		Filter:     filter,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	create := db.Create(f)
	if create.Error != nil {
		return -1, create.Error
	}
	return f.Id, nil
}

func DeleteFilter(id int) error {
	filter := &Filter{}
	filter.Id = id
	d := db.Delete(filter)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func UpdateFilter(id int, filterType TYPE_FILTER, filter string) error {
	f := &Filter{
		Id:         id,
		Type:       filterType,
		Filter:     filter,
		UpdateTime: time.Now().Unix(),
	}
	update := db.Model(f).Update(f)
	if update.Error != nil {
		return update.Error
	}
	return nil
}

func GetFilter(id int) (*Filter, error) {
	f := &Filter{
		Id: id,
	}
	data := db.Where("id = ?", f.Id).Find(f)
	if data.Error != nil {
		return nil, data.Error
	}
	return f, nil
}
