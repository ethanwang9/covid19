package model

import (
	"github.com/ethanwang9/covid19/server/global"
	"go.uber.org/zap"
)

// name: 用户表
// author: Ethan.Wang
// desc:

type User struct {
	Uid      string `gorm:"column:uid" json:"uid"`
	WxOpenid string `gorm:"column:wx_openid" json:"wx_openid"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
	Level    int    `gorm:"column:level" json:"level"`
	Location string `gorm:"column:location" json:"location"`
	Base
}

var UserApp = new(User)

func (u *User) New(user User) *User {
	return &user
}

// HasUser 用户是否存在
//
// value: true->存在 | false->不存在
func (u *User) HasUser() (bool, error) {
	var temp User
	err := global.DB.Where("wx_openid = ?", u.WxOpenid).Find(&temp).Error
	if err != nil {
		global.LOG.Error("GORM|用户表#用户是否存在获取失败", zap.Error(err), zap.Any("data", u))
		return false, err
	}

	if len(temp.Uid) == 0 {
		return false, nil
	}

	return true, nil
}

// Add 添加用户
func (u *User) Add() error {
	err := global.DB.Create(u).Error
	if err != nil {
		global.LOG.Error("GORM|用户表#添加用户失败", zap.Error(err), zap.Any("data", u))
		return err
	}

	return nil
}

// GetByWxOpenid 通过用户微信OPENID获取用户信息
func (u *User) GetByWxOpenid() (User, error) {
	var t User
	err := global.DB.Where("wx_openid = ?", u.WxOpenid).Find(&t).Error
	if err != nil {
		global.LOG.Error("GORM|用户表#通过用户微信OPENID获取用户信息失败", zap.Error(err), zap.Any("data", u))
		return User{}, err
	}

	return t, nil
}

// Get 获取用户
//
// value:
//
//	limit -> 读取 y 条数据
//	offset -> 跳过 x 条数据
func (u *User) Get(limit, offset int) ([]User, error) {
	var t []User
	err := global.DB.Limit(limit).Offset(offset).Find(&t).Error
	if err != nil {
		global.LOG.Error("GORM|用户表#获取全部用户信息失败", zap.Error(err))
		return t, err
	}

	return t, nil
}

// GetCount 获取用户总数
func (u *User) GetCount() (int, error) {
	var t []User
	err := global.DB.Find(&t).Error
	if err != nil {
		global.LOG.Error("GORM|用户表#获取用户总数失败", zap.Error(err))
		return 0, err
	}

	return len(t), nil
}

// GetUserByMore 复合查询用户
//
// value:
//
//	limit -> 读取 y 条数据
//	offset -> 跳过 x 条数据
func (u *User) GetUserByMore(limit, offset int) ([]User, int, error) {
	var temp []User

	// 查询数据
	var err error
	if u.Level == 0 {
		err = global.DB.Where("(uid like ?) and (nickname like ?)", "%"+u.Uid+"%", "%"+u.Nickname+"%").Limit(limit).Offset(offset).Find(&temp).Error

	} else {
		err = global.DB.Where("(uid like ?) and (nickname like ?) and (level = ?)", "%"+u.Uid+"%", "%"+u.Nickname+"%", u.Level).Limit(limit).Offset(offset).Find(&temp).Error

	}
	if err != nil {
		global.LOG.Error("GORM|用户表#复合查询用户失败", zap.Error(err), zap.Any("data", u))
		return temp, 0, err
	}

	// 返回长度
	var temp2 []User
	if u.Level == 0 {
		err = global.DB.Where("(uid like ?) and (nickname like ?)", "%"+u.Uid+"%", "%"+u.Nickname+"%").Find(&temp2).Error

	} else {
		err = global.DB.Where("(uid like ?) and (nickname like ?) and (level = ?)", "%"+u.Uid+"%", "%"+u.Nickname+"%", u.Level).Find(&temp2).Error

	}
	if err != nil {
		global.LOG.Error("GORM|用户表#复合查询用户失败", zap.Error(err), zap.Any("data", u))
		return temp, 0, err
	}

	return temp, len(temp2), nil
}

// UpdateByLevel 更新用户权限
func (u *User) UpdateByLevel() error {
	err := global.DB.Model(&User{}).Where("uid = ?", u.Uid).Update("level", u.Level).Error
	if err != nil {
		global.LOG.Error("GORM|用户表#更新用户权限失败", zap.Error(err), zap.Any("data", u))
		return err
	}

	return nil
}

// UpdateByLocation 更新用户IP属地
func (u *User) UpdateByLocation() error {
	err := global.DB.Model(&User{}).Where("uid = ?", u.Uid).Update("location", u.Location).Error
	if err != nil {
		global.LOG.Error("GORM|用户表#更新用户IP属地失败", zap.Error(err), zap.Any("data", u))
		return err
	}

	return nil
}

// GetUserCount 获取用户数量
func (u *User) GetUserCount() (int64, error) {
	var count int64
	err := global.DB.Model(u).Count(&count).Error
	if err != nil {
		global.LOG.Error("GORM|用户表#获取用户数量失败", zap.Error(err))
		return 0, err
	}

	return count, nil
}
