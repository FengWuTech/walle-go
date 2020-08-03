package menu_model

import (
	"time"
	"walle-go/models"
)

const (
	typeModule     = "module"
	typeController = "controller"
)

type Menus struct {
	Id *int `gorm:"id;default:null" json:"id,omitempty"`
	NameCn *string `gorm:"name_cn;default:null" json:"name_cn,omitempty"`
	NameEn *string `gorm:"name_en;default:null" json:"name_en,omitempty"`
	Pid *int `gorm:"pid;default:null" json:"pid,omitempty"`
	Type *string `gorm:"type;default:action" json:"type,omitempty"`
	Sequence *int `gorm:"sequence;default:0" json:"sequence,omitempty"`
	Role *string `gorm:"role;default:''" json:"role,omitempty"`
	Archive *int `gorm:"archive;default:0" json:"archive,omitempty"`
	Icon *string `gorm:"icon;default:''" json:"icon,omitempty"`
	Url *string `gorm:"url;default:''" json:"url,omitempty"`
	Visible *int `gorm:"visible;default:1" json:"visible,omitempty"`
	CreatedAt *time.Time `gorm:"created_at;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at;default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func menuUrl(url string) string {
	if url == "/" {
		return url
	}
	return "/plm" + url
}

func ToList() []map[string]interface{} {
	var menus []Menus
	var retList []map[string]interface{}
	data := make(map[int]interface{})
	models.GetDB().Order("sequence asc").Find(&menus)
	for _, menu := range menus {
		if *menu.Type == typeModule {
			module := map[string]interface{}{
				"title": *menu.NameCn,
				"icon": *menu.Icon,
				"sub_menu": make([]map[string]interface{}, 0),
			}
			if menu.Url != nil && *menu.Url != "" {
				module["url"] = menuUrl(*menu.Url)
			}
			data[*menu.Id] = module
		} else if *menu.Type == typeController {
			module := data[*menu.Pid].(map[string]interface{})["sub_menu"].([]map[string]interface {})
			module = append(module, map[string]interface{}{
				"title": *menu.NameCn,
				"icon": *menu.Icon,
				"url": menuUrl(*menu.Url),
			})
			data[*menu.Pid].(map[string]interface{})["sub_menu"] = module
		}
	}
	for _, item := range data {
		retList = append(retList, item.(map[string]interface{}))
	}
	return retList
}