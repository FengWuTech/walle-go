package environment_model

import (
	"fmt"
	"time"
	"walle-go/models"
	util "walle-go/pkg"
)

type Environments struct {
	Id        *int       `gorm:"id;default:null" json:"id,omitempty"`
	Name      *string    `gorm:"name;default:master" json:"name,omitempty"`
	SpaceId   *int       `gorm:"space_id;default:0" json:"space_id,omitempty"`
	Status    *int       `gorm:"status;default:1" json:"status,omitempty"`
	CreatedAt *time.Time `gorm:"created_at;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at;default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func (item *Environments) toJson() map[string]interface{} {
	return map[string]interface{}{
		"created_at":    item.CreatedAt.Format("2006-01-02 15:04:05"),
		"enable_audit":  false,
		"enable_block":  false,
		"enable_create": false,
		"enable_delete": true,
		"enable_online": false,
		"enable_update": true,
		"enable_view":   true,
		"env_name":      item.Name,
		"id":            item.Id,
		"space_id":      item.SpaceId,
		"status":        item.Status,
		"updated_at":    item.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func Item(id int) map[string]interface{} {
	var item Environments
	models.GetDB().Where("id = ?", id).First(&item)
	if item.Id == nil {
		return nil
	}
	return item.toJson()
}

func Add(payload map[string]interface{}) *int {
	var item Environments
	util.StructCopyUseJson(payload, &item)
	models.GetDB().Create(&item)
	return item.Id
}

func Update(id int, payload map[string]interface{}) error {
	var item Environments
	util.StructCopyUseJson(payload, &item)
	err := models.GetDB().Model(Environments{}).Where("id = ?", id).Updates(item)
	return err.Error
}

func Remove(id int) bool {
	status := models.StatusRemove
	item := Environments{
		Status: &status,
	}
	err := models.GetDB().Model(Environments{}).Where("id = ?", id).Updates(item)
	return err.Error == nil
}

func List(page int, pageSize int, name string, spaceId int) ([]map[string]interface{}, int) {
	var (
		items []Environments
		count int
		ret   []map[string]interface{}
	)
	query := models.GetDB().Model(Environments{}).Where("status != ?", models.StatusRemove)
	if name != "" {
		query = query.Where("name like ?", fmt.Sprintf("%s%s%s", "%", name, "%"))
	}
	if spaceId > 0 {
		query = query.Where("space_id = ?", spaceId)
	}
	query.Count(&count)
	query.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&items)
	for _, env := range items {
		item := env.toJson()
		ret = append(ret, item)
	}
	return ret, count
}
