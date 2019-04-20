package modules

import (
	"github.com/in2store/service-in2-book/constants/types"
	"github.com/in2store/service-in2-book/database"
	"github.com/johnnyeven/libtools/sqlx"
)

type CreateBookMetaParams struct {
	// 作者ID
	UserID uint64 `json:"userID,string"`
	// 标题
	Title string `json:"title"`
	// 封面图片key
	CoverKey string `json:"coverKey" default:""`
	// 简介
	Comment string `json:"comment" default:""`
	// 文档语言
	BookLanguage types.BookLanguage `json:"bookLanguage" default:""`
	// 代码语言
	CodeLanguage types.CodeLanguage `json:"codeLanguage" default:""`
}

func CreateBookMeta(bookID uint64, req CreateBookMetaParams, db *sqlx.DB) (meta *database.BookMeta, err error) {
	meta = &database.BookMeta{
		BookID:       bookID,
		UserID:       req.UserID,
		Status:       types.BOOK_STATUS__PENGDING,
		Title:        req.Title,
		CoverKey:     req.CoverKey,
		Comment:      req.Comment,
		BookLanguage: req.BookLanguage,
		CodeLanguage: req.CodeLanguage,
	}
	err = meta.Create(db)
	if err != nil {
		return nil, err
	}
	return
}

func GetBookMetaByBookID(bookID uint64, db *sqlx.DB, withLock bool) (meta *database.BookMeta, err error) {
	meta = &database.BookMeta{
		BookID: bookID,
	}
	if withLock {
		err = meta.FetchByBookIDForUpdate(db)
	} else {
		err = meta.FetchByBookID(db)
	}
	if err != nil {
		return nil, err
	}
	return meta, nil
}
