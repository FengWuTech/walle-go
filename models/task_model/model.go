package task_model

import (
	"fmt"
	"strings"
	"time"
	"walle-go/models"
	"walle-go/models/project_model"
	"walle-go/models/server_model"
	util "walle-go/pkg"
)

type Tasks struct {
	Id                   *int       `gorm:"id;default:null" json:"id,omitempty"`
	Name                 *string    `gorm:"name;default:null" json:"name,omitempty"`
	UserId               *int       `gorm:"user_id;default:null" json:"user_id,omitempty"`
	ProjectId            *int       `gorm:"project_id;default:null" json:"project_id,omitempty"`
	Action               *int       `gorm:"action;default:0" json:"action,omitempty"`
	Status               *int       `gorm:"status;default:null" json:"status,omitempty"`
	LinkId               *string    `gorm:"link_id;default:''" json:"link_id,omitempty"`
	ExLinkId             *string    `gorm:"ex_link_id;default:''" json:"ex_link_id,omitempty"`
	Servers              *string    `gorm:"servers;default:null" json:"servers,omitempty"`
	CommitId             *string    `gorm:"commit_id;default:''" json:"commit_id,omitempty"`
	Branch               *string    `gorm:"branch;default:master" json:"branch,omitempty"`
	Tag                  *string    `gorm:"tag;default:''" json:"tag,omitempty"`
	FileTransmissionMode *int       `gorm:"file_transmission_mode;default:1" json:"file_transmission_mode,omitempty"`
	FileList             *string    `gorm:"file_list;default:null" json:"file_list,omitempty"`
	IsRollback           *int       `gorm:"is_rollback;default:null" json:"is_rollback,omitempty"`
	CreatedAt            *time.Time `gorm:"created_at;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt            *time.Time `gorm:"updated_at;default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
	UserName             *string    `gorm:"user_name;default:null" json:"user_name,omitempty"`

	Project *project_model.Projects `gorm:"foreignkey:ProjectId" json:"project"` // 项目
}

func (item *Tasks) ServerInfo() []map[string]interface{} {
	var servers []server_model.Servers
	if item.Servers == nil || *item.Servers == "" {
		return nil
	}
	serverIds := strings.Split(*item.Servers, ",")
	models.GetDB().Where("id in (?)", serverIds).Order("id desc").Find(&servers)
	serverInfo := make([]map[string]interface{}, 0)
	for _, server := range servers {
		serverInfo = append(serverInfo, map[string]interface{}{
			"id":            server.Id,
			"name":          server.Name,
			"host":          server.Host,
			"user_model":    server.User,
			"port":          server.Port,
			"created_at":    server.CreatedAt.Format("2006-01-02 15:04:05"),
			"updated_at":    server.UpdatedAt.Format("2006-01-02 15:04:05"),
			"enable_view":   true,
			"enable_update": true,
			"enable_delete": true,
			"enable_create": false,
			"enable_online": false,
			"enable_audit":  true,
			"enable_block":  false,
		})
	}
	return serverInfo
}

func (item *Tasks) ToMap() map[string]interface{} {
	envName := ""
	if item != nil && item.Project != nil && item.Project.Environment != nil && item.Project.Environment.Name != nil {
		envName = *item.Project.Environment.Name
	}
	return map[string]interface{}{
		"created_at":             item.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at":             item.UpdatedAt.Format("2006-01-02 15:04:05"),
		"id":                     item.Id,
		"name":                   item.Name,
		"user_id":                item.UserId,
		"user_name":              item.UserName,
		"project_id":             item.Project.Id,
		"project_name":           item.Project.Name,
		"action":                 item.Action,
		"status":                 item.Status,
		"link_id":                item.LinkId,
		"ex_link_id":             item.ExLinkId,
		"servers":                item.Servers,
		"servers_info":           item.ServerInfo(),
		"commit_id":              item.CommitId,
		"branch":                 item.Branch,
		"tag":                    item.Tag,
		"file_transmission_mode": item.FileTransmissionMode,
		"file_list":              item.FileList,
		"is_rollback":            item.IsRollback,
		"environment_name":       envName,
		// enable,
		"enable_view":     true,
		"enable_update":   true,
		"enable_delete":   true,
		"enable_create":   false,
		"enable_online":   true,
		"enable_audit":    true,
		"enable_rollback": true,
		"project":         item.Project.ToMap(),
	}
}

func Item(uid int) map[string]interface{} {
	var item Tasks
	models.GetDB().Where("id = ?", uid).Preload("Project").Preload("Project.Environment").First(&item)
	if item.Id == nil {
		return nil
	}
	return item.ToMap()
}

func Add(payload map[string]interface{}) *int {
	var item Tasks
	util.StructCopyUseJson(payload, &item)
	models.GetDB().Create(&item)
	return item.Id
}

func Update(id int, payload map[string]interface{}) error {
	var item Tasks
	util.StructCopyUseJson(payload, &item)
	err := models.GetDB().Model(Tasks{}).Where("id = ?", id).Updates(item)
	return err.Error
}

func Remove(id int) bool {
	status := models.StatusRemove
	env := Tasks{
		Status: &status,
	}
	err := models.GetDB().Model(Tasks{}).Where("id = ?", id).Updates(env)
	return err.Error == nil
}

func List(page int, pageSize int, name string) ([]map[string]interface{}, int) {
	var (
		items []Tasks
		count int
		ret   []map[string]interface{}
	)
	query := models.GetDB().Model(Tasks{}).Where("status != ?", models.StatusRemove)
	if name != "" {
		query = query.Where("name like ?", fmt.Sprintf("%s%s%s", "%", name, "%"))
	}
	query.Count(&count)
	query.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Preload("Project").Preload("Project.Environment").Find(&items)
	for _, i := range items {

		item := i.ToMap()
		ret = append(ret, item)
	}
	return ret, count
}
