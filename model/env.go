package model

import (
	"time"
)

// Env ...
type Env struct {
	Id          int    `gorm:"primary_key,column:id"`
	App         int    `gorm:"column:app_id"`
	Name        string `gorm:"column:env_name" `
	Key         string `gorm:"column:env_key"`
	Weight      int    `gorm:"column:weight"`
	Description string `gorm:"column:description"`
	Filter      int    `gorm:"column:filter"`
	DeployUser  int    `gorm:"column:deploy_user"`
	DeployTime  int64  `gorm:"column:deploy_time"`
	CreateUser  int    `gorm:"column:create_user"`
	UpdateUser  int    `gorm:"column:update_user"`
	CreateTime  int64  `gorm:"column:create_time"`
	UpdateTime  int64  `gorm:"column:update_time"`
}

// TableName ...
func (Env) TableName() string {
	return "m_env"
}

// InsertEnv ...
func InsertEnv(app int, name, desc, key, filter string, weight int) error {
	var insertFilterId int
	insertFilterId = -1
	if filter != "" {
		id, err := InsertFilter(Mode_FILTER_LUA, filter)
		if err != nil {
			return err
		}
		insertFilterId = id
	}
	create := db.Create(&Env{
		App:         app,
		Name:        name,
		Key:         key,
		Weight:      weight,
		Filter:      insertFilterId,
		Description: desc,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	})
	return create.Error
}

// CheckEnvKeyUnique ...
func CheckEnvKeyUnique(app int, key string) bool {
	first := db.Where(&Env{App: app, Key: key}).First(&Env{})
	if first.Error == nil {
		return false
	}
	return true
}

// ListEnvs ...
func ListEnvs(app int, filter string, limit int, offset int) ([]*Env, error) {
	envs := make([]*Env, 0)
	fields := []string{"id", "env_name", "description", "env_key", "weight", "filter", "deploy_time", "create_time", "update_time"}
	find := db.Select(fields).Where("app_id = ? and (env_name LIKE ? or description LIKE ?)", app, "%"+filter+"%", "%"+filter+"%").Order("update_time desc").Limit(limit).Offset(offset).Find(&envs)
	if find.Error != nil {
		return nil, find.Error
	}
	return envs, nil
}

// DeleteEnv ...
func DeleteEnv(id int) error {
	env := &Env{}
	env.Id = id
	d := db.Delete(env)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

// UpdateEnv ...
func UpdateEnv(id int, name string, desc string, weight int) error {
	env := &Env{
		Id:          id,
		Name:        name,
		Weight:      weight,
		Description: desc,
		UpdateTime:  time.Now().Unix(),
	}
	update := db.Model(env).Update(env)
	if update.Error != nil {
		return update.Error
	}
	return nil
}

// UpdateEnvFilter ...
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

// GetEnv ...
func GetEnv(id int) (*Env, error) {
	f := &Env{
		Id: id,
	}
	data := db.Where("id = ?", f.Id).Find(f)
	if data.Error != nil {
		return nil, data.Error
	}
	return f, nil
}

// UpdateEnvDeployData ...
func UpdateEnvDeployData(id int) error {
	env := &Env{
		Id:         id,
		DeployTime: time.Now().Unix(),
	}
	update := db.Model(env).Update(env)
	if update.Error != nil {
		return update.Error
	}
	return nil
}
