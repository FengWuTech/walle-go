package server_model

import (
	"fmt"
	"time"
	"walle-go/models"
	util "walle-go/pkg"
)

type Servers struct {
	Id        *int       `gorm:"id;default:null" json:"id,omitempty"`
	Name      *string    `gorm:"name;default:''" json:"name,omitempty"`
	User      *string    `gorm:"user;default:null" json:"user,omitempty"`
	Host      *string    `gorm:"host;default:null" json:"host,omitempty"`
	Port      *int       `gorm:"port;default:22" json:"port,omitempty"`
	Status    *int       `gorm:"status;default:1" json:"status,omitempty"`
	CreatedAt *time.Time `gorm:"created_at;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at;default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func (item *Servers) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"created_at":    item.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at":    item.UpdatedAt.Format("2006-01-02 15:04:05"),
		"enable_audit":  false,
		"enable_block":  false,
		"enable_create": false,
		"enable_delete": true,
		"enable_online": false,
		"enable_update": true,
		"enable_view":   true,
		"host":          item.Host,
		"id":            item.Id,
		"name":          item.Name,
		"port":          item.Port,
		"user":          item.User,
	}
}

func Item(id int) map[string]interface{} {
	var item Servers
	models.GetDB().Where("id = ?", id).First(&item)
	if item.Id == nil {
		return nil
	}
	return item.ToMap()
}

func Add(payload map[string]interface{}) *int {
	var item Servers
	util.StructCopyUseJson(payload, &item)
	models.GetDB().Create(&item)
	return item.Id
}

func Update(id int, payload map[string]interface{}) error {
	var item Servers
	util.StructCopyUseJson(payload, &item)
	err := models.GetDB().Model(Servers{}).Where("id = ?", id).Updates(item)
	return err.Error
}

func Remove(id int) bool {
	status := models.StatusRemove
	env := Servers{
		Status: &status,
	}
	err := models.GetDB().Model(Servers{}).Where("id = ?", id).Updates(env)
	return err.Error == nil
}

func List(page int, pageSize int, name string) ([]map[string]interface{}, int) {
	var (
		items []Servers
		count int
		ret   []map[string]interface{}
	)
	query := models.GetDB().Model(Servers{}).Where("status != ?", models.StatusRemove)
	if name != "" {
		query = query.Where("name like ?", fmt.Sprintf("%s%s%s", "%", name, "%"))
	}
	query.Count(&count)
	query.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&items)
	for _, i := range items {
		item := i.ToMap()
		ret = append(ret, item)
	}
	return ret, count
}
