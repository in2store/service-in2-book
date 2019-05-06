package database

import (
	"github.com/johnnyeven/libtools/sqlx/presets"
)

//go:generate libtools gen model Tag --database DBIn2Book --table-name t_tag --with-comments
// @def primary ID
// @def unique_index U_tag_id TagID
// @def index I_heat Heat
type Tag struct {
	presets.PrimaryID
	// 业务ID
	TagID uint64 `json:"tagID,string" db:"F_tag_id" sql:"bigint(64) unsigned NOT NULL"`
	// 名称
	Name string `json:"name" db:"F_name" sql:"varchar(16) NOT NULL"`
	// 热度
	Heat uint32 `json:"heat" db:"F_heat" sql:"int(11) NOT NULL DEFAULT '0'"`

	presets.OperateTime
	presets.SoftDelete
}
