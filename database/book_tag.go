package database

import (
	"github.com/johnnyeven/libtools/sqlx/presets"
)

//go:generate libtools gen model BookTag --database DBIn2Book --table-name t_book_tag --with-comments
// @def primary ID
// @def unique_index U_book_tag TagID BookID
// @def index I_tag TagID
// @def index I_book BookID
type BookTag struct {
	presets.PrimaryID
	// 标签ID
	TagID uint64 `json:"tagID,string" db:"F_tag_id" sql:"bigint(64) unsigned NOT NULL"`
	// 文档ID
	BookID uint64 `json:"bookID,string" db:"F_book_id" sql:"bigint(64) unsigned NOT NULL"`

	presets.OperateTime
	presets.SoftDelete
}
