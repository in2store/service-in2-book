package modules

import (
	"github.com/in2store/service-in2-book/constants/errors"
	"github.com/in2store/service-in2-book/database"
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/johnnyeven/libtools/sqlx/builder"
)

type CreateCategoryBody struct {
	// 分类Key
	CategoryKey string `json:"categoryKey"`
	// 分类名
	Name string `json:"name"`
	// 排序
	Sort int32 `json:"sort" default:"0"`
}

func CreateCategory(req CreateCategoryBody, db *sqlx.DB) (result *database.Category, err error) {
	result = &database.Category{
		CategoryKey: req.CategoryKey,
		Name:        req.Name,
		Sort:        req.Sort,
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
	// 排序
	Sort int32 `json:"sort" default:"1"`
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
