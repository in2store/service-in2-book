package modules

import (
	"github.com/in2store/service-in2-book/constants/types"
	"github.com/in2store/service-in2-book/database"
	"github.com/johnnyeven/libtools/courier/enumeration"
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
		Selected:     enumeration.BOOL__FALSE,
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

type UpdateBookMetaParams struct {
	// 状态
	Status types.BookStatus `json:"status" default:""`
	// 是否精选
	Selected enumeration.Bool `json:"selected" default:""`
	// 标题
	Title string `json:"title" default:""`
	// 封面图片key
	CoverKey string `json:"coverKey" default:""`
	// 简介
	Comment string `json:"comment" default:""`
	// 文档语言
	BookLanguage types.BookLanguage `json:"bookLanguage" default:""`
	// 代码语言
	CodeLanguage types.CodeLanguage `json:"codeLanguage" default:""`
}

func (req UpdateBookMetaParams) UpdateParamsByRequest(meta *database.BookMeta) {
	if req.Status != types.BOOK_STATUS_UNKNOWN {
		meta.Status = req.Status
	}
	if req.Selected != enumeration.BOOL_UNKNOWN {
		meta.Selected = req.Selected
	}
	if req.Title != "" {
		meta.Title = req.Title
	}
	if req.CoverKey != "" {
		meta.CoverKey = req.CoverKey
	}
	if req.Comment != "" {
		meta.Comment = req.Comment
	}
	if req.BookLanguage != types.BOOK_LANGUAGE_UNKNOWN {
		meta.BookLanguage = req.BookLanguage
	}
	if req.CodeLanguage != types.CODE_LANGUAGE_UNKNOWN {
		meta.CodeLanguage = req.CodeLanguage
	}
}

func UpdateBookMeta(bookID uint64, req UpdateBookMetaParams, db *sqlx.DB, withLock bool) (meta *database.BookMeta, err error) {
	meta, err = GetBookMetaByBookID(bookID, db, withLock)
	if err != nil {
		return
	}

	req.UpdateParamsByRequest(meta)
	err = meta.UpdateByIDWithStruct(db)
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
	// 分类
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

func SetBookTag(bookID, tagID uint64, db *sqlx.DB) error {
	//category := &database.Category{
	//	CategoryKey: categoryKey,
	//}
	//err := category.FetchByCategoryKey(db)
	//if err != nil {
	//	if sqlx.DBErr(err).IsNotFound() {
	//		return errors.CategoryKeyNotFound
	//	}
	//	return err
	//}
	//bookCategory := &database.BookTag{
	//	CategoryKey: categoryKey,
	//	BookID:      bookID,
	//}
	//err = bookCategory.Create(db)
	//if err != nil {
	//	if sqlx.DBErr(err).IsConflict() {
	//		return errors.BookCategoryConflict
	//	}
	//	return err
	//}
	return nil
}
