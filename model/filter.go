package model

import (
	"time"
)

type Mode_FILTER int

const (
	Mode_FILTER_LUA Mode_FILTER = 1
)

type Filter struct {
	Id         int         `gorm:"primary_key,column:id" json:"id"`
	Mode       Mode_FILTER `gorm:"column:mode" json:"mode"`
	Filter     string      `gorm:"column:filter" json:"filter"`
	CreateUser int         `gorm:"column:create_user" json:"create_user"`
	UpdateUser int         `gorm:"column:update_user" json:"update_user"`
	CreateTime int64       `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64       `gorm:"column:update_time" json:"update_time"`
}

type FilterModel struct {
	Id       int    `gorm:"primary_key,column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Template string `gorm:"column:template" json:"template"`
	Note     string `gorm:"column:note" json:"note"`
}

func (Filter) TableName() string {
	return "m_filter"
}

func (FilterModel) TableName() string {
	return "m_filter_mode"
}

func InsertFilter(filterMode Mode_FILTER, filter string) (int, error) {
	f := &Filter{
		Mode:       filterMode,
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

func UpdateFilter(id int, filterType Mode_FILTER, filter string) error {
	f := &Filter{
		Id:         id,
		Mode:       filterType,
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

func GetFilterModes() ([]*FilterModel, error) {
	modes := make([]*FilterModel, 0)
	find := db.Find(&modes)
	if find.Error != nil {
		return nil, find.Error
	}
	return modes, nil
}

func GetFilterMode(id int) (string, error) {
	mode := &FilterModel{
		Id: id,
	}
	find := db.Find(&mode)
	if find.Error != nil {
		return "", find.Error
	}
	return mode.Name, nil
}
