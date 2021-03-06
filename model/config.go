package model

import (
	"time"
)

// Config ...
type Config struct {
	Id          int    `gorm:"primary_key,column:id"`
	App         int    `gorm:"column:app_id"`
	Env         int    `gorm:"column:env_id"`
	Name        string `gorm:"column:config_name" `
	Key         string `gorm:"column:config_key"`
	Val         string `gorm:"column:config_value"`
	Schema      string `gorm:"column:config_schema"`
	Description string `gorm:"column:description"`
	DeployUser  int    `gorm:"column:deploy_user"`
	DeployTime  int64  `gorm:"column:deploy_time"`
	DeployTag   int    `gorm:"column:deploy_tag"`
	CreateUser  int    `gorm:"column:create_user"`
	UpdateUser  int    `gorm:"column:update_user"`
	CreateTime  int64  `gorm:"column:create_time"`
	UpdateTime  int64  `gorm:"column:update_time"`
}

// TableName ...
func (Config) TableName() string {
	return "m_config"
}

// InsertConfig ...
func InsertConfig(app int, env int, name, desc, key string) error {
	create := db.Create(&Config{
		App:         app,
		Env:         env,
		Name:        name,
		Key:         key,
		Description: desc,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	})
	return create.Error
}

// CheckConfigKeyUnique ...
func CheckConfigKeyUnique(app int, env int, key string) bool {
	first := db.Where(&Config{App: app, Env: env, Key: key}).First(&Config{})
	if first.Error == nil {
		return false
	}
	return true
}

// ListConfigs ...
func ListConfigs(app int, env int, filter string, limit int, offset int) ([]*Config, error) {
	configs := make([]*Config, 0)
	fields := []string{"id", "config_name", "description", "config_key", "config_value", "config_schema", "deploy_time", "deploy_tag", "create_time", "update_time"}
	find := db.Select(fields).Where("app_id = ? and env_id = ?  and (config_name LIKE ? or description LIKE ?)", app, env, "%"+filter+"%", "%"+filter+"%").Order("update_time desc").Limit(limit).Offset(offset).Find(&configs)
	if find.Error != nil {
		return nil, find.Error
	}
	return configs, nil
}

// DeleteConfig ...
func DeleteConfig(id int) error {
	config := &Config{}
	config.Id = id
	d := db.Delete(config)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

// UpdateConfig ...
func UpdateConfig(id int, name string, desc string) error {
	config := &Config{
		Id:          id,
		Name:        name,
		Description: desc,
		UpdateTime:  time.Now().Unix(),
	}
	update := db.Model(config).Update(config)
	if update.Error != nil {
		return update.Error
	}
	return nil
}

// UpdateConfigVal ...
func UpdateConfigVal(id int, val string) error {
	config := &Config{
		Id:         id,
		Val:        val,
		UpdateTime: time.Now().Unix(),
	}
	update := db.Model(config).Update(config)
	if update.Error != nil {
		return update.Error
	}
	return nil
}

// UpdateConfigSchema ...
func UpdateConfigSchema(id int, schema string) error {
	config := &Config{
		Id:         id,
		Schema:     schema,
		UpdateTime: time.Now().Unix(),
	}
	update := db.Model(config).Update(config)
	if update.Error != nil {
		return update.Error
	}
	return nil
}

// UpdateConfigValAndConfig ...
func UpdateConfigValAndConfig(id int, val string, schema string) error {
	config := &Config{
		Id:         id,
		Val:        val,
		Schema:     schema,
		UpdateTime: time.Now().Unix(),
	}
	update := db.Model(config).Update(config)
	if update.Error != nil {
		return update.Error
	}
	return nil
}

// GetConfig ...
func GetConfig(id int) (*Config, error) {
	f := &Config{
		Id: id,
	}
	data := db.Where("id = ?", f.Id).Find(f)
	if data.Error != nil {
		return nil, data.Error
	}
	return f, nil
}

// CountConfig ...
func CountConfig() (interface{}, error) {
	var count int
	c := db.Model(&Config{}).Count(&count)
	if c.Error != nil {
		return nil, c.Error
	}
	return count, nil
}

// UpdateConfigDeployData ...
func UpdateConfigDeployData(id int, tag int) error {
	env := &Config{
		Id:         id,
		DeployTime: time.Now().Unix(),
		DeployTag:  tag,
	}
	update := db.Model(env).Update(env)
	if update.Error != nil {
		return update.Error
	}
	return nil
}
