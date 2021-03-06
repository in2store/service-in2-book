package database

import (
	"github.com/johnnyeven/libtools/courier/enumeration"
	"github.com/johnnyeven/libtools/sqlx/presets"
)

//go:generate libtools gen model Category --database DBIn2Book --table-name t_category --with-comments
// @def primary ID
// @def unique_index U_category CategoryKey
// @def index I_reserved Reserved
type Category struct {
	presets.PrimaryID
	// 业务ID
	CategoryKey string `json:"categoryKey" db:"F_category_key" sql:"varchar(32) NOT NULL"`
	// 分类名
	Name string `json:"name" db:"F_name" sql:"varchar(32) NOT NULL"`
	// 图标类名
	IconClassName string `json:"iconClassName" db:"F_icon_class_name" sql:"varchar(32) DEFAULT NULL"`
	// 排序
	Sort int32 `json:"sort" db:"F_sort" sql:"int DEFAULT '0'"`
	// 是否保留为系统预设
	Reserved enumeration.Bool `json:"reserved" db:"F_reserved" sql:"tinyint(4) DEFAULT '0'"`

	presets.OperateTime
	presets.SoftDelete
}
