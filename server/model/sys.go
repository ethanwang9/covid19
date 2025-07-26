package model

import (
	"github.com/ethanwang9/covid19/server/global"
	"go.uber.org/zap"
)

// name: 系统信息表
// author: Ethan.Wang
// desc:

type Sys struct {
	Base
	Id        string `gorm:"colmun:id" json:"id,omitempty" `
	Copyright string `gorm:"colmun:copyright" json:"copyright"`
	GovNo1    string `gorm:"colmun:gov_no1" json:"gov_no1"`
	GovNo2    string `gorm:"colmun:gov_no2" json:"gov_no2"`
	MpUrl     string `gorm:"colmun:mp_url" json:"mp_url"`
	MpImg     string `gorm:"colmun:mp_img" json:"mp_img"`
	Mail      string `gorm:"colmun:mail" json:"mail"`
	Blog      string `gorm:"colmun:blog" json:"blog"`
}

var SysApp = new(Sys)

func (s *Sys) New(sys Sys) *Sys {
	sys.Id = "9"
	return &sys
}

// Get 获取系统信息
func (s *Sys) Get() (Sys, error) {
	var t Sys
	err := global.DB.Where("id = ?", s.Id).Find(&t).Error
	if err != nil {
		global.LOG.Error("GORM|系统信息表#获取系统信息失败", zap.Error(err))
		return Sys{}, err
	}

	return t, nil
}

// Set 设置系统信息
func (s *Sys) Set() error {
	err := global.DB.Model(&Sys{}).Where("id = ?", s.Id).Updates(map[string]interface{}{
		"copyright": s.Copyright,
		"gov_no1":   s.GovNo1,
		"gov_no2":   s.GovNo2,
		"mp_url":    s.MpUrl,
		"mp_img":    s.MpImg,
		"mail":      s.Mail,
		"blog":      s.Blog,
	}).Error
	if err != nil {
		global.LOG.Error("GORM|系统信息表#设置系统信息失败", zap.Error(err), zap.Any("data", s))
		return err
	}

	return nil
}
