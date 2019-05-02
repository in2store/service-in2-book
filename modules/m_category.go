package modules

import (
	"github.com/in2store/service-in2-book/constants/errors"
	"github.com/in2store/service-in2-book/database"
	"github.com/johnnyeven/libtools/courier/enumeration"
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/johnnyeven/libtools/sqlx/builder"
)

type CreateCategoryBody struct {
	// 分类Key
	CategoryKey string `json:"categoryKey"`
	// 分类名
	Name string `json:"name"`
	// 图标类名
	IconClassName string `json:"iconClassName"`
	// 排序
	Sort int32 `json:"sort" default:"0"`
	// 是否保留为系统预设
	Reserved enumeration.Bool `json:"reserved" default:"0"`
}

func CreateCategory(req CreateCategoryBody, db *sqlx.DB) (result *database.Category, err error) {
	result = &database.Category{
		CategoryKey:   req.CategoryKey,
		Name:          req.Name,
		IconClassName: req.IconClassName,
		Sort:          req.Sort,
		Reserved:      req.Reserved,
	}
	err = result.Create(db)
	if err != nil {
		if sqlx.DBErr(err).IsConflict() {
			return nil, errors.CategoryKeyConflict
		}
		return nil, err
	}
	return
}

type UpdateCategoryBody struct {
	// 分类名
	Name string `json:"name"`
	// 图标类名
	IconClassName string `json:"iconClassName"`
	// 排序
	Sort int32 `json:"sort" default:"0"`
	// 是否保留为系统预设
	Reserved enumeration.Bool `json:"reserved" default:"0"`
}

func UpdateCategory(categoryKey string, req UpdateCategoryBody, db *sqlx.DB, withLock bool) error {
	var err error
	c := &database.Category{
		CategoryKey: categoryKey,
	}
	if withLock {
		err = c.FetchByCategoryKeyForUpdate(db)
	} else {
		err = c.FetchByCategoryKey(db)
	}

	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return errors.CategoryKeyNotFound
		}
		return err
	}
	c.IconClassName = req.IconClassName
	c.Sort = req.Sort
	c.Name = req.Name
	c.Reserved = req.Reserved
	err = c.UpdateByCategoryKeyWithStruct(db)
	if err != nil {
		return err
	}
	return nil
}

func GetCategoriesSortAsc(size, offset int32, db *sqlx.DB) (result database.CategoryList, count int32, err error) {
	c := &database.Category{}
	table := c.T()
	stmt := table.
		Select().
		Comment("Category.GetCategoriesSort").
		OrderAscBy(table.F("Sort")).
		Limit(size).
		Offset(offset)

	errForCount := db.Do(stmt.For(builder.Count(builder.Star()))).Scan(&count).Err()
	if errForCount != nil {
		err = errForCount
		return
	}
	err = db.Do(stmt).Scan(&result).Err()
	return
}
