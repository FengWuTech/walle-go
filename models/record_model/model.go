package record_model

import "time"

type Records struct {
	Id        *int       `gorm:"id;default:null" json:"id,omitempty"`
	Stage     *string    `gorm:"stage;default:null" json:"stage,omitempty"`
	Sequence  *int       `gorm:"sequence;default:null" json:"sequence,omitempty"`
	UserId    *int       `gorm:"user_id;default:null" json:"user_id,omitempty"`
	TaskId    *int       `gorm:"task_id;default:null" json:"task_id,omitempty"`
	Status    *int       `gorm:"status;default:null" json:"status,omitempty"`
	Host      *string    `gorm:"host;default:''" json:"host,omitempty"`
	User      *string    `gorm:"user_model;default:''" json:"user_model,omitempty"`
	Command   *string    `gorm:"command;default:null" json:"command,omitempty"`
	Success   *string    `gorm:"success;default:null" json:"success,omitempty"`
	Error     *string    `gorm:"error;default:null" json:"error,omitempty"`
	CreatedAt *time.Time `gorm:"created_at;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at;default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

