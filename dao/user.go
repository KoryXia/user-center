package dao

import (
	"errors"
	"gorm.io/gorm"
	"user-center/global"
	"user-center/model"
)

func GetUserByUsername(username string) (*model.User, error) {
	user := model.User{}
	if err := global.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		return nil, errors.New("failed to get user by username")
	}
	return &user, nil
}

func CreateOneUser(user *model.User) error {
	if err := global.DB.Create(user).Error; err != nil {
		return errors.New("failed to create a new user")
	}
	return nil
}

func UpdateUserInfo(user *model.User) (int64, error) {
	tx := global.DB.Where("username = ?", user.Username).Updates(user)
	return tx.RowsAffected, tx.Error
}
