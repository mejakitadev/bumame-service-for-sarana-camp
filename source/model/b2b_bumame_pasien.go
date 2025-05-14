package model

import "time"

type B2BPasien struct {
	Id         uint64    `gorm:"primaryKey;column:id;type:serial8;autoIncrement:true" json:"id"`
	PasienName string    `gorm:"type:varchar(255);not null;column:pasien_name" json:"pasien_name"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	IsDeleted  uint      `json:"is_deleted" gorm:"default:0"`
}

func (B2BPasien) TableName() string {
	return "b2b_pasien"
}
