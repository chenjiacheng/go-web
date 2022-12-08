package user

import (
	"go-web/app/models"
	"go-web/pkg/database"
	"go-web/pkg/hash"
	"go-web/pkg/logger"
)

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `gorm:"default:null" json:"-"`
	Phone    string `gorm:"default:null" json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() (err error) {
	if err = database.DB.Create(&userModel).Error; err != nil {
		logger.LogIf(err)
		return err
	}

	return nil
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
