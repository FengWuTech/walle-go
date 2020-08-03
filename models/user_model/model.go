package user_model

import (
	"fmt"
	"time"
	"walle-go/models"
	util "walle-go/pkg"
	"walle-go/pkg/security"
)

const (
	AvatarPath    = "/avatar/"
	DefaultAvatar = "default.jpg"
)

type Users struct {
	Id              *int       `gorm:"id;default:null" json:"id,omitempty"`
	Username        *string    `gorm:"username;default:null" json:"username,omitempty"`
	IsEmailVerified *int       `gorm:"is_email_verified;default:0" json:"is_email_verified,omitempty"`
	Email           *string    `gorm:"email;default:null" json:"email,omitempty"`
	Password        *string    `gorm:"password;default:null" json:"password,omitempty"`
	PasswordHash    *string    `gorm:"password_hash;default:null" json:"password_hash,omitempty"`
	Avatar          *string    `gorm:"avatar;default:default.jpg" json:"avatar,omitempty"`
	Role            *string    `gorm:"role;default:''" json:"role,omitempty"`
	Status          *int       `gorm:"status;default:1" json:"status,omitempty"`
	LastSpace       *int       `gorm:"last_space;default:0" json:"last_space,omitempty"`
	CreatedAt       *time.Time `gorm:"created_at;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt       *time.Time `gorm:"updated_at;default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func avatarUrl(avatar string) string {
	if avatar == "" {
		avatar = DefaultAvatar
	}
	return AvatarPath + avatar
}

func statusMapping(status int) string {
	mapping := map[int]string{
		-1: "删除",
		0:  "新建",
		1:  "正常",
		2:  "冻结",
	}
	return mapping[status]
}

func (item *Users) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"avatar":            avatarUrl(*item.Avatar),
		"created_at":        item.CreatedAt.Format("2006-01-02 15:04:05"),
		"email":             item.Email,
		"enable_audit":      false,
		"enable_block":      false,
		"enable_create":     false,
		"enable_online":     false,
		"enable_update":     true,
		"enable_delete":     true,
		"enable_view":       true,
		"id":                item.Id,
		"is_email_verified": item.IsEmailVerified,
		"last_space":        item.LastSpace,
		"status":            statusMapping(*item.Status),
		"updated_at":        item.UpdatedAt.Format("2006-01-02 15:04:05"),
		"user_id":           item.Id,
		"username":          item.Username,
	}
}

func (item *Users) GetPassword(password string) string {
	return security.GeneratePasswordHash(password)
}

func (item *Users) VerifyPassword(password string) bool {
	if item.Password == nil || *item.Password == "" {
		return false
	}
	return security.CheckPasswordHash(*item.Password, password)
}

func Item(uid int) map[string]interface{} {
	var item Users
	models.GetDB().Where("id = ?", uid).First(&item)
	if item.Id == nil {
		return nil
	}
	return item.ToMap()
}

func GetByEmail(email string) *Users {
	var item Users
	models.GetDB().Where("email = ?", email).First(&item)
	if item.Id == nil {
		return nil
	}
	return &item
}

func Add(payload map[string]interface{}) *int {
	var item Users
	util.StructCopyUseJson(payload, &item)
	models.GetDB().Create(&item)
	return item.Id
}

func Update(id int, payload map[string]interface{}) error {
	var item Users
	util.StructCopyUseJson(payload, &item)
	err := models.GetDB().Model(Users{}).Where("id = ?", id).Updates(item)
	return err.Error
}

func Remove(id int) bool {
	status := models.StatusRemove
	env := Users{
		Status: &status,
	}
	err := models.GetDB().Model(Users{}).Where("id = ?", id).Updates(env)
	return err.Error == nil
}

func List(page int, pageSize int, name string) ([]map[string]interface{}, int) {
	var (
		items []Users
		count int
		ret   []map[string]interface{}
	)
	query := models.GetDB().Model(Users{}).Where("status != ?", models.StatusRemove)
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
