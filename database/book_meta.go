package database

import (
	"github.com/in2store/service-in2-book/constants/types"
	"github.com/johnnyeven/libtools/sqlx/presets"
)

//go:generate libtools gen model BookMeta --database DBIn2Book --table-name t_book_meta --with-comments
// @def primary ID
// @def unique_index U_book_id BookID
// @def index I_author_status UserID Status
// @def index I_category CategoryKey Status
type BookMeta struct {
	presets.PrimaryID
	// 业务ID
	BookID uint64 `json:"bookID,string" db:"F_book_id" sql:"bigint(64) unsigned NOT NULL"`
	// 类别ID
	CategoryKey string `json:"categoryKey" db:"F_category_key" sql:"varchar(32) NOT NULL"`
	// 作者ID
	UserID uint64 `json:"userID,string" db:"F_user_id" sql:"bigint(64) unsigned NOT NULL"`
	// 状态
	Status types.BookStatus `json:"status" db:"F_status" sql:"tinyint(4) unsigned NOT NULL"`
	// 标题
	Title string `json:"title" db:"F_title" sql:"varchar(255) NOT NULL"`
	// 封面图片key
	CoverKey string `json:"coverKey" db:"F_cover_key" sql:"varchar(64) DEFAULT NULL"`
	// 简介
	Comment string `json:"comment" db:"F_comment" sql:"text DEFAULT NULL"`
	// 文档语言
	BookLanguage types.BookLanguage `json:"bookLanguage" db:"F_book_language" sql:"tinyint(4) DEFAULT NULL"`
	// 代码语言
	CodeLanguage types.CodeLanguage `json:"codeLanguage" db:"F_code_language" sql:"tinyint(4) DEFAULT NULL"`

	presets.OperateTime
	presets.SoftDelete
}
