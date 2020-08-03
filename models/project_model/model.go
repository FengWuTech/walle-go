package project_model

import (
	"fmt"
	"time"
	"walle-go/models"
	"walle-go/models/environment_model"
	util "walle-go/pkg"
)

type Projects struct {
	Id             *int       `gorm:"id;default:null" json:"id,omitempty"`
	UserId         *int       `gorm:"user_id;default:null" json:"user_id,omitempty"`
	Name           *string    `gorm:"name;default:master" json:"name,omitempty"`
	EnvironmentId  *int       `gorm:"environment_id;default:null" json:"environment_id,omitempty"`
	SpaceId        *int       `gorm:"space_id;default:0" json:"space_id,omitempty"`
	Status         *int       `gorm:"status;default:1" json:"status,omitempty"`
	Master         *string    `gorm:"master;default:''" json:"master,omitempty"`
	Version        *string    `gorm:"version;default:''" json:"version,omitempty"`
	Excludes       *string    `gorm:"excludes;default:null" json:"excludes,omitempty"`
	TargetRoot     *string    `gorm:"target_root;default:null" json:"target_root,omitempty"`
	TargetReleases *string    `gorm:"target_releases;default:null" json:"target_releases,omitempty"`
	ServerIds      *string    `gorm:"server_ids;default:null" json:"server_ids,omitempty"`
	TaskVars       *string    `gorm:"task_vars;default:null" json:"task_vars,omitempty"`
	PrevDeploy     *string    `gorm:"prev_deploy;default:null" json:"prev_deploy,omitempty"`
	PostDeploy     *string    `gorm:"post_deploy;default:null" json:"post_deploy,omitempty"`
	PrevRelease    *string    `gorm:"prev_release;default:null" json:"prev_release,omitempty"`
	PostRelease    *string    `gorm:"post_release;default:null" json:"post_release,omitempty"`
	KeepVersionNum *int       `gorm:"keep_version_num;default:20" json:"keep_version_num,omitempty"`
	RepoUrl        *string    `gorm:"repo_url;default:''" json:"repo_url,omitempty"`
	RepoUsername   *string    `gorm:"repo_username;default:''" json:"repo_username,omitempty"`
	RepoPassword   *string    `gorm:"repo_password;default:''" json:"repo_password,omitempty"`
	RepoMode       *string    `gorm:"repo_mode;default:branch" json:"repo_mode,omitempty"`
	RepoType       *string    `gorm:"repo_type;default:git" json:"repo_type,omitempty"`
	NoticeType     *string    `gorm:"notice_type;default:''" json:"notice_type,omitempty"`
	NoticeHook     *string    `gorm:"notice_hook;default:null" json:"notice_hook,omitempty"`
	TaskAudit      *int       `gorm:"task_audit;default:0" json:"task_audit,omitempty"`
	CreatedAt      *time.Time `gorm:"created_at;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt      *time.Time `gorm:"updated_at;default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
	IsInclude      *int       `gorm:"is_include;default:0" json:"is_include,omitempty"`

	Environment *environment_model.Environments `gorm:"foreignkey:EnvironmentId" json:"environment"` // 环境
}

func (item *Projects) ToMap() map[string]interface{} {
	envName := ""
	if item != nil && item.Environment != nil && item.Environment.Name != nil {
		envName = *item.Environment.Name
	}
	return map[string]interface{}{
		"created_at":       item.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at":       item.UpdatedAt.Format("2006-01-02 15:04:05"),
		"enable_audit":     false,
		"enable_block":     false,
		"enable_create":    false,
		"enable_delete":    true,
		"enable_online":    false,
		"enable_update":    true,
		"enable_view":      true,
		"environment_id":   item.EnvironmentId,
		"environment_name": envName,
		"excludes":         item.Excludes,
		"id":               item.Id,
		"is_include":       item.IsInclude,
		"keep_version_num": item.KeepVersionNum,
		"master":           item.Master,
		"name":             item.Name,
		"notice_hook":      item.NoticeHook,
		"notice_type":      item.NoticeType,
		"post_deploy":      item.PostDeploy,
		"post_release":     item.PostRelease,
		"prev_deploy":      item.PrevDeploy,
		"prev_release":     item.PrevRelease,
		"repo_mode":        item.RepoMode,
		"repo_password":    item.RepoPassword,
		"repo_type":        item.RepoType,
		"repo_url":         item.RepoUrl,
		"repo_username":    item.RepoUsername,
		"server_ids":       item.ServerIds,
		"space_id":         item.SpaceId,
		"space_name":       "plm",
		"status":           item.Status,
		"target_releases":  item.TargetReleases,
		"target_root":      item.TargetRoot,
		"task_audit":       item.TaskAudit,
		"task_vars":        item.TaskVars,
		"user_id":          item.UserId,
		"version":          item.Version,
	}
}

func Item(id int) map[string]interface{} {
	var item Projects
	models.GetDB().Where("id = ?", id).Preload("Environment").First(&item)
	if item.Id == nil {
		return nil
	}
	return item.ToMap()
}

func Add(payload map[string]interface{}) *int {
	var item Projects
	util.StructCopyUseJson(payload, &item)
	models.GetDB().Create(&item)
	return item.Id
}

func Update(id int, payload map[string]interface{}) error {
	var item Projects
	util.StructCopyUseJson(payload, &item)
	err := models.GetDB().Model(Projects{}).Where("id = ?", id).Updates(item)
	return err.Error
}

func Remove(id int) bool {
	status := models.StatusRemove
	env := Projects{
		Status: &status,
	}
	err := models.GetDB().Model(Projects{}).Where("id = ?", id).Updates(env)
	return err.Error == nil
}

func List(page int, pageSize int, name string, environmentId int, spaceId int) ([]map[string]interface{}, int) {
	var (
		items []Projects
		count int
		ret   []map[string]interface{}
	)
	query := models.GetDB().Model(Projects{}).Where("status != ?", models.StatusRemove)
	if name != "" {
		query = query.Where("name like ?", fmt.Sprintf("%s%s%s", "%", name, "%"))
	}
	if spaceId > 0 {
		query = query.Where("space_id = ?", spaceId)
	}
	if environmentId > 0 {
		query = query.Where("environment_id = ?", environmentId)
	}
	query.Count(&count)
	query.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Preload("Environment").Find(&items)
	for _, env := range items {
		item := env.ToMap()
		ret = append(ret, item)
	}
	return ret, count
}
