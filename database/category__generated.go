package database

import (
	fmt "fmt"
	time "time"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
	github_com_johnnyeven_libtools_sqlx "github.com/johnnyeven/libtools/sqlx"
	github_com_johnnyeven_libtools_sqlx_builder "github.com/johnnyeven/libtools/sqlx/builder"
	github_com_johnnyeven_libtools_timelib "github.com/johnnyeven/libtools/timelib"
)

var CategoryTable *github_com_johnnyeven_libtools_sqlx_builder.Table

func init() {
	CategoryTable = DBIn2Book.Register(&Category{})
}

func (category *Category) D() *github_com_johnnyeven_libtools_sqlx.Database {
	return DBIn2Book
}

func (category *Category) T() *github_com_johnnyeven_libtools_sqlx_builder.Table {
	return CategoryTable
}

func (category *Category) TableName() string {
	return "t_category"
}

type CategoryFields struct {
	ID            *github_com_johnnyeven_libtools_sqlx_builder.Column
	CategoryKey   *github_com_johnnyeven_libtools_sqlx_builder.Column
	Name          *github_com_johnnyeven_libtools_sqlx_builder.Column
	IconClassName *github_com_johnnyeven_libtools_sqlx_builder.Column
	Sort          *github_com_johnnyeven_libtools_sqlx_builder.Column
	Reserved      *github_com_johnnyeven_libtools_sqlx_builder.Column
	CreateTime    *github_com_johnnyeven_libtools_sqlx_builder.Column
	UpdateTime    *github_com_johnnyeven_libtools_sqlx_builder.Column
	Enabled       *github_com_johnnyeven_libtools_sqlx_builder.Column
}

var CategoryField = struct {
	ID            string
	CategoryKey   string
	Name          string
	IconClassName string
	Sort          string
	Reserved      string
	CreateTime    string
	UpdateTime    string
	Enabled       string
}{
	ID:            "ID",
	CategoryKey:   "CategoryKey",
	Name:          "Name",
	IconClassName: "IconClassName",
	Sort:          "Sort",
	Reserved:      "Reserved",
	CreateTime:    "CreateTime",
	UpdateTime:    "UpdateTime",
	Enabled:       "Enabled",
}

func (category *Category) Fields() *CategoryFields {
	table := category.T()

	return &CategoryFields{
		ID:            table.F(CategoryField.ID),
		CategoryKey:   table.F(CategoryField.CategoryKey),
		Name:          table.F(CategoryField.Name),
		IconClassName: table.F(CategoryField.IconClassName),
		Sort:          table.F(CategoryField.Sort),
		Reserved:      table.F(CategoryField.Reserved),
		CreateTime:    table.F(CategoryField.CreateTime),
		UpdateTime:    table.F(CategoryField.UpdateTime),
		Enabled:       table.F(CategoryField.Enabled),
	}
}

func (category *Category) IndexFieldNames() []string {
	return []string{"CategoryKey", "ID"}
}

func (category *Category) ConditionByStruct() *github_com_johnnyeven_libtools_sqlx_builder.Condition {
	table := category.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(category)

	conditions := []*github_com_johnnyeven_libtools_sqlx_builder.Condition{}

	for _, fieldName := range category.IndexFieldNames() {
		if v, exists := fieldValues[fieldName]; exists {
			conditions = append(conditions, table.F(fieldName).Eq(v))
			delete(fieldValues, fieldName)
		}
	}

	if len(conditions) == 0 {
		panic(fmt.Errorf("at least one of field for indexes has value"))
	}

	for fieldName, v := range fieldValues {
		conditions = append(conditions, table.F(fieldName).Eq(v))
	}

	condition := github_com_johnnyeven_libtools_sqlx_builder.And(conditions...)

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	return condition
}

func (category *Category) PrimaryKey() github_com_johnnyeven_libtools_sqlx.FieldNames {
	return github_com_johnnyeven_libtools_sqlx.FieldNames{"ID"}
}
func (category *Category) UniqueIndexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{"U_category": github_com_johnnyeven_libtools_sqlx.FieldNames{"CategoryKey", "Enabled"}}
}
func (category *Category) Comments() map[string]string {
	return map[string]string{
		"CategoryKey":   "业务ID",
		"CreateTime":    "",
		"Enabled":       "",
		"ID":            "",
		"IconClassName": "图标类名",
		"Name":          "分类名",
		"Reserved":      "是否保留为系统预设",
		"Sort":          "排序",
		"UpdateTime":    "",
	}
}

func (category *Category) Create(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	category.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if category.CreateTime.IsZero() {
		category.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	category.UpdateTime = category.CreateTime

	stmt := category.D().
		Insert(category).
		Comment("Category.Create")

	dbRet := db.Do(stmt)
	err := dbRet.Err()

	if err == nil {
		lastInsertID, _ := dbRet.LastInsertId()
		category.ID = uint64(lastInsertID)
	}

	return err
}

func (category *Category) DeleteByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (err error) {
	table := category.T()

	stmt := table.Delete().
		Comment("Category.DeleteByStruct").
		Where(category.ConditionByStruct())

	err = db.Do(stmt).Err()
	return
}

func (category *Category) CreateOnDuplicateWithUpdateFields(db *github_com_johnnyeven_libtools_sqlx.DB, updateFields []string) error {
	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	category.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if category.CreateTime.IsZero() {
		category.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	category.UpdateTime = category.CreateTime

	table := category.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(category, updateFields...)

	delete(fieldValues, "ID")

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	m := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		m[field] = true
	}

	// fields of unique index can not update
	delete(m, "CreateTime")

	for _, fieldNames := range category.UniqueIndexes() {
		for _, field := range fieldNames {
			delete(m, field)
		}
	}

	if len(m) == 0 {
		panic(fmt.Errorf("no fields for updates"))
	}

	for field := range fieldValues {
		if !m[field] {
			delete(fieldValues, field)
		}
	}

	stmt := table.
		Insert().Columns(cols).Values(vals...).
		OnDuplicateKeyUpdate(table.AssignsByFieldValues(fieldValues)...).
		Comment("Category.CreateOnDuplicateWithUpdateFields")

	return db.Do(stmt).Err()
}

func (category *Category) FetchByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	category.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := category.T()
	stmt := table.Select().
		Comment("Category.FetchByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(category.ID),
			table.F("Enabled").Eq(category.Enabled),
		))

	return db.Do(stmt).Scan(category).Err()
}

func (category *Category) FetchByIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	category.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := category.T()
	stmt := table.Select().
		Comment("Category.FetchByIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(category.ID),
			table.F("Enabled").Eq(category.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(category).Err()
}

func (category *Category) DeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	category.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := category.T()
	stmt := table.Delete().
		Comment("Category.DeleteByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(category.ID),
			table.F("Enabled").Eq(category.Enabled),
		))

	return db.Do(stmt).Scan(category).Err()
}

func (category *Category) UpdateByIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	category.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := category.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("Category.UpdateByIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(category.ID),
			table.F("Enabled").Eq(category.Enabled),
		))

	dbRet := db.Do(stmt).Scan(category)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return category.FetchByID(db)
	}
	return nil
}

func (category *Category) UpdateByIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(category, zeroFields...)
	return category.UpdateByIDWithMap(db, fieldValues)
}

func (category *Category) SoftDeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	category.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := category.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("Category.SoftDeleteByID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(category.ID),
			table.F("Enabled").Eq(category.Enabled),
		))

	dbRet := db.Do(stmt).Scan(category)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return category.DeleteByID(db)
		}
		return err
	}
	return nil
}

func (category *Category) FetchByCategoryKey(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	category.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := category.T()
	stmt := table.Select().
		Comment("Category.FetchByCategoryKey").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("CategoryKey").Eq(category.CategoryKey),
			table.F("Enabled").Eq(category.Enabled),
		))

	return db.Do(stmt).Scan(category).Err()
}

func (category *Category) FetchByCategoryKeyForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	category.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := category.T()
	stmt := table.Select().
		Comment("Category.FetchByCategoryKeyForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("CategoryKey").Eq(category.CategoryKey),
			table.F("Enabled").Eq(category.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(category).Err()
}

func (category *Category) DeleteByCategoryKey(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	category.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := category.T()
	stmt := table.Delete().
		Comment("Category.DeleteByCategoryKey").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("CategoryKey").Eq(category.CategoryKey),
			table.F("Enabled").Eq(category.Enabled),
		))

	return db.Do(stmt).Scan(category).Err()
}

func (category *Category) UpdateByCategoryKeyWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	category.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := category.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("Category.UpdateByCategoryKeyWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("CategoryKey").Eq(category.CategoryKey),
			table.F("Enabled").Eq(category.Enabled),
		))

	dbRet := db.Do(stmt).Scan(category)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return category.FetchByCategoryKey(db)
	}
	return nil
}

func (category *Category) UpdateByCategoryKeyWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(category, zeroFields...)
	return category.UpdateByCategoryKeyWithMap(db, fieldValues)
}

func (category *Category) SoftDeleteByCategoryKey(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	category.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := category.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("Category.SoftDeleteByCategoryKey").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("CategoryKey").Eq(category.CategoryKey),
			table.F("Enabled").Eq(category.Enabled),
		))

	dbRet := db.Do(stmt).Scan(category)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return category.DeleteByCategoryKey(db)
		}
		return err
	}
	return nil
}

type CategoryList []Category

// deprecated
func (categoryList *CategoryList) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (count int32, err error) {
	*categoryList, count, err = (&Category{}).FetchList(db, size, offset, conditions...)
	return
}

func (category *Category) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (categoryList CategoryList, count int32, err error) {
	categoryList = CategoryList{}

	table := category.T()

	condition := github_com_johnnyeven_libtools_sqlx_builder.And(conditions...)

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Category.FetchList").
		Where(condition)

	errForCount := db.Do(stmt.For(github_com_johnnyeven_libtools_sqlx_builder.Count(github_com_johnnyeven_libtools_sqlx_builder.Star()))).Scan(&count).Err()
	if errForCount != nil {
		err = errForCount
		return
	}

	stmt = stmt.Limit(size).Offset(offset)

	stmt = stmt.OrderDescBy(table.F("CreateTime"))

	err = db.Do(stmt).Scan(&categoryList).Err()

	return
}

func (category *Category) List(db *github_com_johnnyeven_libtools_sqlx.DB, condition *github_com_johnnyeven_libtools_sqlx_builder.Condition) (categoryList CategoryList, err error) {
	categoryList = CategoryList{}

	table := category.T()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Category.List").
		Where(condition)

	err = db.Do(stmt).Scan(&categoryList).Err()

	return
}

func (category *Category) ListByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (categoryList CategoryList, err error) {
	categoryList = CategoryList{}

	table := category.T()

	condition := category.ConditionByStruct()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Category.ListByStruct").
		Where(condition)

	err = db.Do(stmt).Scan(&categoryList).Err()

	return
}

// deprecated
func (categoryList *CategoryList) BatchFetchByCategoryKeyList(db *github_com_johnnyeven_libtools_sqlx.DB, categoryKeyList []string) (err error) {
	*categoryList, err = (&Category{}).BatchFetchByCategoryKeyList(db, categoryKeyList)
	return
}

func (category *Category) BatchFetchByCategoryKeyList(db *github_com_johnnyeven_libtools_sqlx.DB, categoryKeyList []string) (categoryList CategoryList, err error) {
	if len(categoryKeyList) == 0 {
		return CategoryList{}, nil
	}

	table := category.T()

	condition := table.F("CategoryKey").In(categoryKeyList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Category.BatchFetchByCategoryKeyList").
		Where(condition)

	err = db.Do(stmt).Scan(&categoryList).Err()

	return
}

// deprecated
func (categoryList *CategoryList) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (err error) {
	*categoryList, err = (&Category{}).BatchFetchByIDList(db, idList)
	return
}

func (category *Category) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (categoryList CategoryList, err error) {
	if len(idList) == 0 {
		return CategoryList{}, nil
	}

	table := category.T()

	condition := table.F("ID").In(idList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Category.BatchFetchByIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&categoryList).Err()

	return
}
