package modules

import (
	"github.com/in2store/service-in2-book/constants/errors"
	"github.com/in2store/service-in2-book/constants/types"
	"github.com/in2store/service-in2-book/database"
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/johnnyeven/libtools/sqlx/builder"
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
		Status:       types.BOOK_STATUS__READY,
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

type GetBooksMetaRequest struct {
	// 用户ID
	UserID uint64 `name:"userID,string" json:"userID,string" in:"query" default:""`
	// 状态
	Status types.BookStatus `name:"status" json:"status" in:"query" default:""`
}

func GetBooksMeta(req GetBooksMetaRequest, size, offset int32, db *sqlx.DB) (result database.BookMetaList, count int32, err error) {
	meta := &database.BookMeta{}
	table := meta.T()
	var conditions *builder.Condition

	if req.UserID != 0 {
		conditions = builder.And(conditions, table.F("UserID").Eq(req.UserID))
	}
	if req.Status != types.BOOK_STATUS_UNKNOWN {
		conditions = builder.And(conditions, table.F("Status").Eq(req.Status))
	}

	return meta.FetchList(db, size, offset, conditions)
}

func SetBookCategory(bookID uint64, categoryKey string, db *sqlx.DB) error {
	bookCategory := &database.BookCategory{
		CategoryKey: categoryKey,
		BookID:      bookID,
	}
	err := bookCategory.Create(db)
	if err != nil {
		if sqlx.DBErr(err).IsConflict() {
			return errors.BookCategoryConflict
		}
		return err
	}
	return nil
}
