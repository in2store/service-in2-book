package database

import (
	"github.com/johnnyeven/libtools/sqlx/presets"
)

//go:generate libtools gen model BookCategory --database DBIn2Book --table-name t_book_category --with-comments
// @def primary ID
// @def unique_index U_book_category CategoryKey BookID
// @def index I_category CategoryKey
type BookCategory struct {
	presets.PrimaryID
	// 分类标识
	CategoryKey string `json:"categoryKey" db:"F_category_key" sql:"varchar(32) NOT NULL"`
	// 文档ID
	BookID uint64 `json:"bookID,string" db:"F_book_id" sql:"bigint(64) unsigned NOT NULL"`

	presets.OperateTime
	presets.SoftDelete
}
