package modules

import (
	"github.com/in2store/service-in2-book/database"
	"github.com/johnnyeven/libtools/sqlx"
)

type CreateBookRepoParams struct {
	// 通道ID
	ChannelID uint64 `json:"channelID,string"`
	// 入口地址
	EntryURL string `json:"entryURL"`
	// 代码库全名
	RepoFullName string `json:"repoFullName"`
	// 代码库分支
	RepoBranchName string `json:"repoBranchName"`
	// Summary文件相对地址
	SummaryPath string `json:"summaryPath"`
}

func CreateBookRepo(bookID uint64, req CreateBookRepoParams, db *sqlx.DB) (repo *database.BookRepo, err error) {
	repo = &database.BookRepo{
		BookID:         bookID,
		ChannelID:      req.ChannelID,
		EntryURL:       req.EntryURL,
		RepoFullName:   req.RepoFullName,
		RepoBranchName: req.RepoBranchName,
		SummaryPath:    req.SummaryPath,
	}
	err = repo.Create(db)
	if err != nil {
		return nil, err
	}
	return
}

func GetBookRepoByBookID(bookID uint64, db *sqlx.DB, withLock bool) (repo *database.BookRepo, err error) {
	repo = &database.BookRepo{
		BookID: bookID,
	}
	if withLock {
		err = repo.FetchByBookIDForUpdate(db)
	} else {
		err = repo.FetchByBookID(db)
	}
	if err != nil {
		return nil, err
	}
	return repo, nil
}
