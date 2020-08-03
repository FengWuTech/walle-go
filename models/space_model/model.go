package space_model

import (
	"fmt"
	"time"
	"walle-go/models"
	util "walle-go/pkg"
)

type Spaces struct {
	Id        *int       `gorm:"id;default:null" json:"id,omitempty"`
	UserId    *int       `gorm:"user_id;default:null" json:"user_id,omitempty"`
	Name      *string    `gorm:"name;default:null" json:"name,omitempty"`
	Status    *int       `gorm:"status;default:1" json:"status,omitempty"`
	CreatedAt *time.Time `gorm:"created_at;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at;default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func (item *Spaces) toJson() map[string]interface{} {
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
		"id":            item.Id,
		"user_id":       item.UserId,
		"user_name":     "",
		"name":          item.Name,
		"status":        item.Status,
	}
}

func Item(id int) map[string]interface{} {
	var item Spaces
	models.GetDB().Where("id = ?", id).First(&item)
	if item.Id == nil {
		return nil
	}
	return item.toJson()
}

func Add(payload map[string]interface{}) *int {
	var item Spaces
	util.StructCopyUseJson(payload, &item)
	models.GetDB().Create(&item)
	return item.Id
}

func Update(id int, payload map[string]interface{}) error {
	var item Spaces
	util.StructCopyUseJson(payload, &item)
	err := models.GetDB().Model(Spaces{}).Where("id = ?", id).Updates(item)
	return err.Error
}

func Remove(id int) bool {
	status := models.StatusRemove
	env := Spaces{
		Status: &status,
	}
	err := models.GetDB().Model(Spaces{}).Where("id = ?", id).Updates(env)
	return err.Error == nil
}

func List(page int, pageSize int, name string) ([]map[string]interface{}, int) {
	var (
		items []Spaces
		count int
		ret   []map[string]interface{}
	)
	query := models.GetDB().Model(Spaces{}).Where("status != ?", models.StatusRemove)
	if name != "" {
		query = query.Where("name like ?", fmt.Sprintf("%s%s%s", "%", name, "%"))
	}
	query.Count(&count)
	query.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&items)
	for _, i := range items {
		item := i.toJson()
		ret = append(ret, item)
	}
	return ret, count
}
