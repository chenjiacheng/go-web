package user

import "go-web/app/models"

type User struct {
	models.CommonIDField

	Name     string `json:"name,omitempty"`
	Email    string `gorm:"default:null" json:"-"`
	Phone    string `gorm:"default:null" json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
