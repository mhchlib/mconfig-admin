package model

import (
	"time"
)

// Cluster ...
type Cluster struct {
	Id          int    `gorm:"primary_key,column:id" json:"id"`
	Namespace   string `gorm:"column:namespace"  json:"namespace"`
	Register    string `gorm:"column:register"  json:"register"`
	Description string `gorm:"column:description" json:"description"`
	CreateUser  int    `gorm:"column:create_user" json:"create_user"`
	UpdateUser  int    `gorm:"column:update_user" json:"update_user"`
	CreateTime  int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime  int64  `gorm:"column:update_time" json:"update_time"`
}

// TableName ...
func (Cluster) TableName() string {
	return "m_cluster"
}

// InsertCluster ...
func InsertCluster(namespace, register, desc string) error {
	create := db.Create(&Cluster{
		Namespace:   namespace,
		Register:    register,
		Description: desc,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	})
	return create.Error
}

// ListClusters ...
func ListClusters(filter string, limit int, offset int) ([]*Cluster, error) {
	clusters := make([]*Cluster, 0)
	fields := []string{"id", "namespace", "description", "register", "create_time", "update_time"}
	find := db.Select(fields).Where("namespace LIKE ? or description LIKE ?", "%"+filter+"%", "%"+filter+"%").Order("update_time desc").Limit(limit).Offset(offset).Find(&clusters)
	if find.Error != nil {
		return nil, find.Error
	}
	return clusters, nil
}

// DeleteCluster ...
func DeleteCluster(id int) error {
	cluster := &Cluster{}
	cluster.Id = id
	d := db.Delete(cluster)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

// UpdateCluster ...
func UpdateCluster(id int, namespace string, register string, desc string) error {
	cluster := &Cluster{
		Id:          id,
		Namespace:   namespace,
		Register:    register,
		Description: desc,
		UpdateTime:  time.Now().Unix(),
	}
	update := db.Model(cluster).Update(cluster)
	if update.Error != nil {
		return update.Error
	}
	return nil
}

// GetCluster ...
func GetCluster(id int) (*Cluster, error) {
	f := &Cluster{
		Id: id,
	}
	data := db.Where("id = ?", f.Id).Find(f)
	if data.Error != nil {
		return nil, data.Error
	}
	return f, nil
}

// CountCluster ...
func CountCluster() (interface{}, error) {
	var count int
	c := db.Model(&Cluster{}).Count(&count)
	if c.Error != nil {
		return nil, c.Error
	}
	return count, nil
}
