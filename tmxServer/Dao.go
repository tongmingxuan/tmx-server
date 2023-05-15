// Package tmxServer /*
package tmxServer

import (
	"github.com/tongmingxuan/tmx-server/plugin/pluginList/DaoPlugin"
	"gorm.io/gorm"
)

type Dao struct {
	GormDb *gorm.DB
}

// CommonGetDao
// @Description: 获取gorm链接
// @param pollName
// @return *Dao
func CommonGetDao(pollName string) *Dao {
	if pollName == "" {
		pollName = "default"
	}

	dao := Dao{
		GormDb: Connection(pollName),
	}

	return &dao
}

type DaoInterface interface {
	GetModel() InterfaceModel
}

type BaseDao struct {
	DaoPlugin.DaoPlugin
}
