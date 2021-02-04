package model

import (
	"time"
)

type Env struct {
	Id          int    `gorm:"primary_key,column:id"`
	App         int    `gorm:"column:app_id"`
	Name        string `gorm:"column:env_name" `
	Key         string `gorm:"column:env_key"`
	Description string `gorm:"column:description"`
	Filter      int    `gorm:"column:filter"`
	CreateUser  int    `gorm:"column:create_user"`
	UpdateUser  int    `gorm:"column:update_user"`
	CreateTime  int64  `gorm:"column:create_time"`
	UpdateTime  int64  `gorm:"column:update_time"`
}

func (Env) TableName() string {
	return "m_env"
}

func InsertEnv(app int, name, desc, key, filter string) error {
	var insertFilterId int
	insertFilterId = -1
	if filter != "" {
		id, err := InsertFilter(TYPE_FILTER_LUA, filter)
		if err != nil {
			return err
		}
		insertFilterId = id
	}
	create := db.Create(&Env{
		App:         app,
		Name:        name,
		Key:         key,
		Filter:      insertFilterId,
		Description: desc,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	})
	return create.Error
}

func CheckEnvKeyUnique(app int, key string) bool {
	first := db.Where(&Env{App: app, Key: key}).First(&Env{})
	if first.Error == nil {
		return false
	}
	return true
}

func ListEnvs(app int, filter string, limit int, offset int) ([]*Env, error) {
	Envs := make([]*Env, 0)
	fields := []string{"id", "env_name", "description", "env_key", "filter", "create_time", "update_time"}
	find := db.Select(fields).Where("app_id = ? and (env_name LIKE ? or description LIKE ?)", app, "%"+filter+"%", "%"+filter+"%").Order("update_time desc").Limit(limit).Offset(offset).Find(&Envs)
	if find.Error != nil {
		return nil, find.Error
	}
	return Envs, nil
}

func DeleteEnv(id int) error {
	env := &Env{}
	env.Id = id
	d := db.Delete(env)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func UpdateEnv(id int, name string, desc string) error {
	env := &Env{
		Id:          id,
		Name:        name,
		Description: desc,
		UpdateTime:  time.Now().Unix(),
	}
	update := db.Model(env).Update(env)
	if update.Error != nil {
		return update.Error
	}
	return nil
}

func UpdateEnvFilter(id int, filter int) error {
	env := &Env{
		Id:         id,
		Filter:     filter,
		UpdateTime: time.Now().Unix(),
	}
	update := db.Model(env).Update(env)
	if update.Error != nil {
		return update.Error
	}
	return nil
}
