package models

import "time"

type Employee struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null; type:varchar(50)" json:"name"`
	Email     string    `gorm:"not null; type:varchar(50); unique" json:"email"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}
