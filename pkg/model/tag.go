package model

import (
	"time"
)

type Tag struct {
	Id          int    `gorm:"primary_key,column:id" json:"id"`
	Tag         string `gorm:"column:tag" json:"tag" `
	ConfigId    int    `gorm:"column:config_id" json:"-"`
	Description string `gorm:"column:description" json:"description"`
	Config      string `gorm:"column:config" json:"config"`
	Schema      string `gorm:"column:schema" json:"schema"`
	CreateUser  int    `gorm:"column:create_user" json:"create_user"`
	UpdateUser  int    `gorm:"column:update_user" json:"update_user"`
	CreateTime  int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime  int64  `gorm:"column:update_time" json:"update_time"`
}

func (Tag) TableName() string {
	return "m_tag"
}

func InsertTag(tag, desc string, configId int, config string, schema string) error {
	create := db.Create(&Tag{
		Tag:         tag,
		ConfigId:    configId,
		Config:      config,
		Schema:      schema,
		Description: desc,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	})
	return create.Error
}

func ListTags(configId int, filter string, limit int, offset int) ([]*Tag, error) {
	tags := make([]*Tag, 0)
	fields := []string{"id", "tag", "description", "create_time", "update_time"}
	find := db.Select(fields).Where("config_id = ? and (tag LIKE ? or description LIKE ?)", configId, "%"+filter+"%", "%"+filter+"%").Order("update_time desc").Limit(limit).Offset(offset).Find(&tags)
	if find.Error != nil {
		return nil, find.Error
	}
	return tags, nil
}

func DeleteTag(id int) error {
	tag := &Tag{}
	tag.Id = id
	d := db.Delete(tag)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func GetTag(id int) (*Tag, error) {
	f := &Tag{
		Id: id,
	}
	data := db.Where("id = ?", f.Id).Find(f)
	if data.Error != nil {
		return nil, data.Error
	}
	return f, nil
}
