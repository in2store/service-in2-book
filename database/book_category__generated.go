package database

import (
	fmt "fmt"
	time "time"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
	github_com_johnnyeven_libtools_sqlx "github.com/johnnyeven/libtools/sqlx"
	github_com_johnnyeven_libtools_sqlx_builder "github.com/johnnyeven/libtools/sqlx/builder"
	github_com_johnnyeven_libtools_timelib "github.com/johnnyeven/libtools/timelib"
)

var BookCategoryTable *github_com_johnnyeven_libtools_sqlx_builder.Table

func init() {
	BookCategoryTable = DBIn2Book.Register(&BookCategory{})
}

func (bookCategory *BookCategory) D() *github_com_johnnyeven_libtools_sqlx.Database {
	return DBIn2Book
}

func (bookCategory *BookCategory) T() *github_com_johnnyeven_libtools_sqlx_builder.Table {
	return BookCategoryTable
}

func (bookCategory *BookCategory) TableName() string {
	return "t_book_category"
}

type BookCategoryFields struct {
	ID          *github_com_johnnyeven_libtools_sqlx_builder.Column
	CategoryKey *github_com_johnnyeven_libtools_sqlx_builder.Column
	BookID      *github_com_johnnyeven_libtools_sqlx_builder.Column
	CreateTime  *github_com_johnnyeven_libtools_sqlx_builder.Column
	UpdateTime  *github_com_johnnyeven_libtools_sqlx_builder.Column
	Enabled     *github_com_johnnyeven_libtools_sqlx_builder.Column
}

var BookCategoryField = struct {
	ID          string
	CategoryKey string
	BookID      string
	CreateTime  string
	UpdateTime  string
	Enabled     string
}{
	ID:          "ID",
	CategoryKey: "CategoryKey",
	BookID:      "BookID",
	CreateTime:  "CreateTime",
	UpdateTime:  "UpdateTime",
	Enabled:     "Enabled",
}

func (bookCategory *BookCategory) Fields() *BookCategoryFields {
	table := bookCategory.T()

	return &BookCategoryFields{
		ID:          table.F(BookCategoryField.ID),
		CategoryKey: table.F(BookCategoryField.CategoryKey),
		BookID:      table.F(BookCategoryField.BookID),
		CreateTime:  table.F(BookCategoryField.CreateTime),
		UpdateTime:  table.F(BookCategoryField.UpdateTime),
		Enabled:     table.F(BookCategoryField.Enabled),
	}
}

func (bookCategory *BookCategory) IndexFieldNames() []string {
	return []string{"BookID", "CategoryKey", "ID"}
}

func (bookCategory *BookCategory) ConditionByStruct() *github_com_johnnyeven_libtools_sqlx_builder.Condition {
	table := bookCategory.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookCategory)

	conditions := []*github_com_johnnyeven_libtools_sqlx_builder.Condition{}

	for _, fieldName := range bookCategory.IndexFieldNames() {
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

func (bookCategory *BookCategory) PrimaryKey() github_com_johnnyeven_libtools_sqlx.FieldNames {
	return github_com_johnnyeven_libtools_sqlx.FieldNames{"ID"}
}
func (bookCategory *BookCategory) Indexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{"I_category": github_com_johnnyeven_libtools_sqlx.FieldNames{"CategoryKey"}}
}
func (bookCategory *BookCategory) UniqueIndexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{"U_book_category": github_com_johnnyeven_libtools_sqlx.FieldNames{"CategoryKey", "BookID", "Enabled"}}
}
func (bookCategory *BookCategory) Comments() map[string]string {
	return map[string]string{
		"BookID":      "文档ID",
		"CategoryKey": "分类标识",
		"CreateTime":  "",
		"Enabled":     "",
		"ID":          "",
		"UpdateTime":  "",
	}
}

func (bookCategory *BookCategory) Create(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookCategory.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if bookCategory.CreateTime.IsZero() {
		bookCategory.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	bookCategory.UpdateTime = bookCategory.CreateTime

	stmt := bookCategory.D().
		Insert(bookCategory).
		Comment("BookCategory.Create")

	dbRet := db.Do(stmt)
	err := dbRet.Err()

	if err == nil {
		lastInsertID, _ := dbRet.LastInsertId()
		bookCategory.ID = uint64(lastInsertID)
	}

	return err
}

func (bookCategory *BookCategory) DeleteByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (err error) {
	table := bookCategory.T()

	stmt := table.Delete().
		Comment("BookCategory.DeleteByStruct").
		Where(bookCategory.ConditionByStruct())

	err = db.Do(stmt).Err()
	return
}

func (bookCategory *BookCategory) CreateOnDuplicateWithUpdateFields(db *github_com_johnnyeven_libtools_sqlx.DB, updateFields []string) error {
	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	bookCategory.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if bookCategory.CreateTime.IsZero() {
		bookCategory.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	bookCategory.UpdateTime = bookCategory.CreateTime

	table := bookCategory.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookCategory, updateFields...)

	delete(fieldValues, "ID")

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	m := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		m[field] = true
	}

	// fields of unique index can not update
	delete(m, "CreateTime")

	for _, fieldNames := range bookCategory.UniqueIndexes() {
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
		Comment("BookCategory.CreateOnDuplicateWithUpdateFields")

	return db.Do(stmt).Err()
}

func (bookCategory *BookCategory) FetchByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookCategory.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookCategory.T()
	stmt := table.Select().
		Comment("BookCategory.FetchByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookCategory.ID),
			table.F("Enabled").Eq(bookCategory.Enabled),
		))

	return db.Do(stmt).Scan(bookCategory).Err()
}

func (bookCategory *BookCategory) FetchByIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookCategory.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookCategory.T()
	stmt := table.Select().
		Comment("BookCategory.FetchByIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookCategory.ID),
			table.F("Enabled").Eq(bookCategory.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(bookCategory).Err()
}

func (bookCategory *BookCategory) DeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookCategory.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookCategory.T()
	stmt := table.Delete().
		Comment("BookCategory.DeleteByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookCategory.ID),
			table.F("Enabled").Eq(bookCategory.Enabled),
		))

	return db.Do(stmt).Scan(bookCategory).Err()
}

func (bookCategory *BookCategory) UpdateByIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	bookCategory.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookCategory.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("BookCategory.UpdateByIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookCategory.ID),
			table.F("Enabled").Eq(bookCategory.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookCategory)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return bookCategory.FetchByID(db)
	}
	return nil
}

func (bookCategory *BookCategory) UpdateByIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookCategory, zeroFields...)
	return bookCategory.UpdateByIDWithMap(db, fieldValues)
}

func (bookCategory *BookCategory) SoftDeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookCategory.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookCategory.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("BookCategory.SoftDeleteByID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookCategory.ID),
			table.F("Enabled").Eq(bookCategory.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookCategory)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return bookCategory.DeleteByID(db)
		}
		return err
	}
	return nil
}

func (bookCategory *BookCategory) FetchByCategoryKeyAndBookID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookCategory.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookCategory.T()
	stmt := table.Select().
		Comment("BookCategory.FetchByCategoryKeyAndBookID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("CategoryKey").Eq(bookCategory.CategoryKey),
			table.F("BookID").Eq(bookCategory.BookID),
			table.F("Enabled").Eq(bookCategory.Enabled),
		))

	return db.Do(stmt).Scan(bookCategory).Err()
}

func (bookCategory *BookCategory) FetchByCategoryKeyAndBookIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookCategory.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookCategory.T()
	stmt := table.Select().
		Comment("BookCategory.FetchByCategoryKeyAndBookIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("CategoryKey").Eq(bookCategory.CategoryKey),
			table.F("BookID").Eq(bookCategory.BookID),
			table.F("Enabled").Eq(bookCategory.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(bookCategory).Err()
}

func (bookCategory *BookCategory) DeleteByCategoryKeyAndBookID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookCategory.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookCategory.T()
	stmt := table.Delete().
		Comment("BookCategory.DeleteByCategoryKeyAndBookID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("CategoryKey").Eq(bookCategory.CategoryKey),
			table.F("BookID").Eq(bookCategory.BookID),
			table.F("Enabled").Eq(bookCategory.Enabled),
		))

	return db.Do(stmt).Scan(bookCategory).Err()
}

func (bookCategory *BookCategory) UpdateByCategoryKeyAndBookIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	bookCategory.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookCategory.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("BookCategory.UpdateByCategoryKeyAndBookIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("CategoryKey").Eq(bookCategory.CategoryKey),
			table.F("BookID").Eq(bookCategory.BookID),
			table.F("Enabled").Eq(bookCategory.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookCategory)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return bookCategory.FetchByCategoryKeyAndBookID(db)
	}
	return nil
}

func (bookCategory *BookCategory) UpdateByCategoryKeyAndBookIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookCategory, zeroFields...)
	return bookCategory.UpdateByCategoryKeyAndBookIDWithMap(db, fieldValues)
}

func (bookCategory *BookCategory) SoftDeleteByCategoryKeyAndBookID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookCategory.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookCategory.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("BookCategory.SoftDeleteByCategoryKeyAndBookID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("CategoryKey").Eq(bookCategory.CategoryKey),
			table.F("BookID").Eq(bookCategory.BookID),
			table.F("Enabled").Eq(bookCategory.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookCategory)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return bookCategory.DeleteByCategoryKeyAndBookID(db)
		}
		return err
	}
	return nil
}

type BookCategoryList []BookCategory

// deprecated
func (bookCategoryList *BookCategoryList) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (count int32, err error) {
	*bookCategoryList, count, err = (&BookCategory{}).FetchList(db, size, offset, conditions...)
	return
}

func (bookCategory *BookCategory) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (bookCategoryList BookCategoryList, count int32, err error) {
	bookCategoryList = BookCategoryList{}

	table := bookCategory.T()

	condition := github_com_johnnyeven_libtools_sqlx_builder.And(conditions...)

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookCategory.FetchList").
		Where(condition)

	errForCount := db.Do(stmt.For(github_com_johnnyeven_libtools_sqlx_builder.Count(github_com_johnnyeven_libtools_sqlx_builder.Star()))).Scan(&count).Err()
	if errForCount != nil {
		err = errForCount
		return
	}

	stmt = stmt.Limit(size).Offset(offset)

	stmt = stmt.OrderDescBy(table.F("CreateTime"))

	err = db.Do(stmt).Scan(&bookCategoryList).Err()

	return
}

func (bookCategory *BookCategory) List(db *github_com_johnnyeven_libtools_sqlx.DB, condition *github_com_johnnyeven_libtools_sqlx_builder.Condition) (bookCategoryList BookCategoryList, err error) {
	bookCategoryList = BookCategoryList{}

	table := bookCategory.T()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookCategory.List").
		Where(condition)

	err = db.Do(stmt).Scan(&bookCategoryList).Err()

	return
}

func (bookCategory *BookCategory) ListByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (bookCategoryList BookCategoryList, err error) {
	bookCategoryList = BookCategoryList{}

	table := bookCategory.T()

	condition := bookCategory.ConditionByStruct()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookCategory.ListByStruct").
		Where(condition)

	err = db.Do(stmt).Scan(&bookCategoryList).Err()

	return
}

// deprecated
func (bookCategoryList *BookCategoryList) BatchFetchByBookIDList(db *github_com_johnnyeven_libtools_sqlx.DB, bookIDList []uint64) (err error) {
	*bookCategoryList, err = (&BookCategory{}).BatchFetchByBookIDList(db, bookIDList)
	return
}

func (bookCategory *BookCategory) BatchFetchByBookIDList(db *github_com_johnnyeven_libtools_sqlx.DB, bookIDList []uint64) (bookCategoryList BookCategoryList, err error) {
	if len(bookIDList) == 0 {
		return BookCategoryList{}, nil
	}

	table := bookCategory.T()

	condition := table.F("BookID").In(bookIDList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookCategory.BatchFetchByBookIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookCategoryList).Err()

	return
}

// deprecated
func (bookCategoryList *BookCategoryList) BatchFetchByCategoryKeyList(db *github_com_johnnyeven_libtools_sqlx.DB, categoryKeyList []string) (err error) {
	*bookCategoryList, err = (&BookCategory{}).BatchFetchByCategoryKeyList(db, categoryKeyList)
	return
}

func (bookCategory *BookCategory) BatchFetchByCategoryKeyList(db *github_com_johnnyeven_libtools_sqlx.DB, categoryKeyList []string) (bookCategoryList BookCategoryList, err error) {
	if len(categoryKeyList) == 0 {
		return BookCategoryList{}, nil
	}

	table := bookCategory.T()

	condition := table.F("CategoryKey").In(categoryKeyList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookCategory.BatchFetchByCategoryKeyList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookCategoryList).Err()

	return
}

// deprecated
func (bookCategoryList *BookCategoryList) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (err error) {
	*bookCategoryList, err = (&BookCategory{}).BatchFetchByIDList(db, idList)
	return
}

func (bookCategory *BookCategory) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (bookCategoryList BookCategoryList, err error) {
	if len(idList) == 0 {
		return BookCategoryList{}, nil
	}

	table := bookCategory.T()

	condition := table.F("ID").In(idList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookCategory.BatchFetchByIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookCategoryList).Err()

	return
}
