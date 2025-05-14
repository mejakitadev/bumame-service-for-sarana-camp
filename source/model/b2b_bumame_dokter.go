package model

import "time"

type B2BDokter struct {
	Id         uint64    `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	DokterName string    `gorm:"column:dokter_name;type:varchar(255)" json:"dokter_name"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	IsDeleted  uint      `gorm:"column:is_deleted;default:0" json:"is_deleted"`
}

func (B2BDokter) TableName() string {
	return "b2b_dokter"
}
