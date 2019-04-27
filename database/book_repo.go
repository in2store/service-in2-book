package database

import (
	"github.com/johnnyeven/libtools/sqlx/presets"
)

//go:generate libtools gen model BookRepo --database DBIn2Book --table-name t_book_repo --with-comments
// @def primary ID
// @def unique_index U_book_id BookID
// @def unique_index U_channel_repo_name ChannelID RepoFullName
type BookRepo struct {
	presets.PrimaryID
	// 书籍ID
	BookID uint64 `json:"bookID,string" db:"F_book_id" sql:"bigint(64) unsigned NOT NULL"`
	// 通道ID
	ChannelID uint64 `json:"channelID,string" db:"F_channel_id" sql:"bigint(64) unsigned NOT NULL"`
	// 入口地址
	EntryURL string `json:"entryURL" db:"F_entry_url" sql:"varchar(255) NOT NULL"`
	// 代码库全名
	RepoFullName string `json:"repoFullName" db:"F_repo_full_name" sql:"varchar(255) NOT NULL"`
	// 代码库分支
	RepoBranchName string `json:"repoBranchName" db:"F_repo_branch_name" sql:"varchar(255) NOT NULL"`
	// Summary文件相对地址
	SummaryPath string `json:"summaryPath" db:"F_summary_path" sql:"varchar(255) NOT NULL"`

	presets.OperateTime
	presets.SoftDelete
}
