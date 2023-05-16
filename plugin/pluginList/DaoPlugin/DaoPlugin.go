// Package DaoPlugin /*
package DaoPlugin

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"time"
)

const (
	CreatedAt = "created_at"
	UpdatedAt = "updated_at"
	DeletedAt = "deleted_at"
)

var DeletedNullMap = map[string]interface{}{
	DeletedAt: nil,
}

var UpdatedAtMap = map[string]interface{}{
	UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
}

func BuildFindByWhere(db *gorm.DB, where interface{}) (*gorm.DB, error) {
	return BuildWhere(db.Where(DeletedNullMap), where)
}

func BuildUpdateByWhere(db *gorm.DB, where interface{}) (*gorm.DB, error) {
	return BuildWhere(db.Where(UpdatedAtMap), where)
}

func BuildGetListByWhere(db *gorm.DB, where interface{}) (*gorm.DB, error) {
	return BuildWhere(db.Where(DeletedNullMap), where)
}

func BuildWhere(db *gorm.DB, where interface{}) (*gorm.DB, error) {
	var err error

	t := reflect.TypeOf(where).Kind()
	if t == reflect.Struct || t == reflect.Map {
		db = db.Where(where)
	} else if t == reflect.Slice {
		for _, item := range where.([]interface{}) {
			item := item.([]interface{})
			column := item[0]
			if reflect.TypeOf(column).Kind() == reflect.String {
				count := len(item)
				if count == 1 {
					return nil, errors.New("切片长度不能小于2")
				}
				columnStr := column.(string)
				// 拼接参数形式
				if strings.Index(columnStr, "?") > -1 {
					db = db.Where(column, item[1:]...)
				} else {
					cond := "and" //cond
					opt := "="
					_opt := " = "
					var val interface{}
					if count == 2 {
						opt = "="
						val = item[1]
					} else {
						opt = strings.ToLower(item[1].(string))
						_opt = " " + strings.ReplaceAll(opt, " ", "") + " "
						val = item[2]
					}

					if count == 4 {
						cond = strings.ToLower(strings.ReplaceAll(item[3].(string), " ", ""))
					}

					/*
					   '=', '<', '>', '<=', '>=', '<>', '!=', '<=>',
					   'like', 'like binary', 'not like', 'ilike',
					   '&', '|', '^', '<<', '>>',
					   'rlike', 'regexp', 'not regexp',
					   '~', '~*', '!~', '!~*', 'similar to',
					   'not similar to', 'not ilike', '~~*', '!~~*',
					*/

					if strings.Index(" in notin ", _opt) > -1 {
						// val 是数组类型
						column = columnStr + " " + opt + " (?)"
					} else if strings.Index(" = < > <= >= <> != <=> like likebinary notlike ilike rlike regexp notregexp", _opt) > -1 {
						column = columnStr + " " + opt + " ?"
					}

					if cond == "and" {
						db = db.Where(column, val)
					} else {
						db = db.Or(column, val)
					}
				}
			} else if t == reflect.Map /*Map*/ {
				db = db.Where(item)
			} else {
				/*
					// 解决and 与 or 混合查询，但这种写法有问题，会抛出 invalid query condition
					db = db.Where(func(db *gorm.DB) *gorm.DB {
						db, err = BuildWhere(db, item)
						if err != nil {
							panic(err)
						}
						return db
					})*/

				db, err = BuildWhere(db, item)
				if err != nil {
					return nil, err
				}
			}
		}
	} else {
		return nil, errors.New("参数有误")
	}
	return db, nil
}

// DaoPlugin
// @Description: 基础dao方法
type DaoPlugin struct {
}

func (p DaoPlugin) BuildWhere(db *gorm.DB, where interface{}) (*gorm.DB, error) {
	return BuildWhere(db, where)
}

func (p DaoPlugin) BuildGetListByWhere(db *gorm.DB, where interface{}) (*gorm.DB, error) {
	return BuildGetListByWhere(db, where)
}

func (p DaoPlugin) BuildUpdateByWhere(db *gorm.DB, where interface{}) (*gorm.DB, error) {
	return BuildUpdateByWhere(db, where)
}

func (p DaoPlugin) BuildFindByWhere(db *gorm.DB, where interface{}) (*gorm.DB, error) {
	return BuildFindByWhere(db, where)
}
