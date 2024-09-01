package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"user-center/dao"
	"user-center/global"
	"user-center/model"
	"user-center/utils"
)

func Register(request *model.RegisterRequest) error {
	if request.Username == "" || request.Password == "" {
		return fmt.Errorf("invalid register input")
	}
	userInDB, err := dao.GetUserByUsername(request.Username)

	if err != nil {
		return errors.New("fail to register new user")
	}

	if userInDB != nil {
		return errors.New("user existed")
	}

	hashedPassword, err := utils.HashPassword(request.Password)

	if err != nil {
		return errors.New("fail to register new user")
	}

	user := model.User{
		Username: request.Username,
		Age:      request.Age,
		Gender:   request.Gender,
		Password: hashedPassword,
		Nickname: request.Nickname,
	}
	if err := dao.CreateOneUser(&user); err != nil {
		return errors.New("fail to register new user")
	}
	return nil
}

func Login(request *model.LoginRequest) error {
	value, err := global.Cache.Get(context.Background(), fmt.Sprintf("login_%s", request.Username)).Result()
	// Check if the user is already in cache
	if err == nil {
		var user model.User
		err = json.Unmarshal([]byte(value), &user)
		if err == nil && user.Username == request.Username {
			return nil
		}
	}

	// if not present in cache, read from DB
	user, err := dao.GetUserByUsername(request.Username)

	if user == nil && err == nil {
		return errors.New("user does not exist")
	}
	if !utils.CheckPasswordHash(request.Password, user.Password) {
		return errors.New("password is incorrect")
	}

	cacheValue, err := json.Marshal(&user)

	if err != nil {
		return err
	}

	if _, err := global.Cache.SetNX(
		context.Background(),
		fmt.Sprintf("login_%s", request.Username),
		cacheValue,
		3*time.Minute,
	).Result(); err != nil {
		return errors.New("failed to set into cache")
	}

	return nil
}

func Logout(request *model.LogoutRequest) error {
	key := fmt.Sprintf("login_%s", request.Username)
	if _, err := global.Cache.Get(context.Background(), key).Result(); err != nil {
		return err
	}
	return global.Cache.Del(context.Background(), key).Err()
}

func GetUserInfo(username string) (*model.UserInfoResponse, error) {
	value, err := global.Cache.Get(context.Background(), fmt.Sprintf("login_%s", username)).Result()

	if err == nil {
		var user model.User
		err = json.Unmarshal([]byte(value), &user)
		if err == nil {
			return &model.UserInfoResponse{
				ID:       user.ID,
				Username: user.Username,
				Age:      user.Age,
				Gender:   user.Gender,
				Nickname: user.Nickname,
			}, nil
		}
	}

	user, err := dao.GetUserByUsername(username)
	if user != nil {
		return &model.UserInfoResponse{
			ID:       user.ID,
			Username: user.Username,
			Age:      user.Age,
			Gender:   user.Gender,
		}, nil
	}
	return nil, err
}

func UpdateUserInfo(request *model.UpdateRequest) error {
	rows, err := dao.UpdateUserInfo(&model.User{
		Username: request.Username,
		Age:      request.Age,
		Gender:   request.Gender,
		Nickname: request.Nickname,
	})

	if err == nil && rows == 1 {
		return nil
	}

	return errors.New("failed to update")
}
