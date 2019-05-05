package database

import (
	fmt "fmt"
	time "time"

	github_com_in2_store_service_in2_book_constants_types "github.com/in2store/service-in2-book/constants/types"
	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
	github_com_johnnyeven_libtools_sqlx "github.com/johnnyeven/libtools/sqlx"
	github_com_johnnyeven_libtools_sqlx_builder "github.com/johnnyeven/libtools/sqlx/builder"
	github_com_johnnyeven_libtools_timelib "github.com/johnnyeven/libtools/timelib"
)

var BookMetaTable *github_com_johnnyeven_libtools_sqlx_builder.Table

func init() {
	BookMetaTable = DBIn2Book.Register(&BookMeta{})
}

func (bookMeta *BookMeta) D() *github_com_johnnyeven_libtools_sqlx.Database {
	return DBIn2Book
}

func (bookMeta *BookMeta) T() *github_com_johnnyeven_libtools_sqlx_builder.Table {
	return BookMetaTable
}

func (bookMeta *BookMeta) TableName() string {
	return "t_book_meta"
}

type BookMetaFields struct {
	ID           *github_com_johnnyeven_libtools_sqlx_builder.Column
	BookID       *github_com_johnnyeven_libtools_sqlx_builder.Column
	CategoryKey  *github_com_johnnyeven_libtools_sqlx_builder.Column
	UserID       *github_com_johnnyeven_libtools_sqlx_builder.Column
	Status       *github_com_johnnyeven_libtools_sqlx_builder.Column
	Selected     *github_com_johnnyeven_libtools_sqlx_builder.Column
	Title        *github_com_johnnyeven_libtools_sqlx_builder.Column
	CoverKey     *github_com_johnnyeven_libtools_sqlx_builder.Column
	Comment      *github_com_johnnyeven_libtools_sqlx_builder.Column
	BookLanguage *github_com_johnnyeven_libtools_sqlx_builder.Column
	CodeLanguage *github_com_johnnyeven_libtools_sqlx_builder.Column
	CreateTime   *github_com_johnnyeven_libtools_sqlx_builder.Column
	UpdateTime   *github_com_johnnyeven_libtools_sqlx_builder.Column
	Enabled      *github_com_johnnyeven_libtools_sqlx_builder.Column
}

var BookMetaField = struct {
	ID           string
	BookID       string
	CategoryKey  string
	UserID       string
	Status       string
	Selected     string
	Title        string
	CoverKey     string
	Comment      string
	BookLanguage string
	CodeLanguage string
	CreateTime   string
	UpdateTime   string
	Enabled      string
}{
	ID:           "ID",
	BookID:       "BookID",
	CategoryKey:  "CategoryKey",
	UserID:       "UserID",
	Status:       "Status",
	Selected:     "Selected",
	Title:        "Title",
	CoverKey:     "CoverKey",
	Comment:      "Comment",
	BookLanguage: "BookLanguage",
	CodeLanguage: "CodeLanguage",
	CreateTime:   "CreateTime",
	UpdateTime:   "UpdateTime",
	Enabled:      "Enabled",
}

func (bookMeta *BookMeta) Fields() *BookMetaFields {
	table := bookMeta.T()

	return &BookMetaFields{
		ID:           table.F(BookMetaField.ID),
		BookID:       table.F(BookMetaField.BookID),
		CategoryKey:  table.F(BookMetaField.CategoryKey),
		UserID:       table.F(BookMetaField.UserID),
		Status:       table.F(BookMetaField.Status),
		Selected:     table.F(BookMetaField.Selected),
		Title:        table.F(BookMetaField.Title),
		CoverKey:     table.F(BookMetaField.CoverKey),
		Comment:      table.F(BookMetaField.Comment),
		BookLanguage: table.F(BookMetaField.BookLanguage),
		CodeLanguage: table.F(BookMetaField.CodeLanguage),
		CreateTime:   table.F(BookMetaField.CreateTime),
		UpdateTime:   table.F(BookMetaField.UpdateTime),
		Enabled:      table.F(BookMetaField.Enabled),
	}
}

func (bookMeta *BookMeta) IndexFieldNames() []string {
	return []string{"BookID", "CategoryKey", "ID", "Selected", "Status", "UserID"}
}

func (bookMeta *BookMeta) ConditionByStruct() *github_com_johnnyeven_libtools_sqlx_builder.Condition {
	table := bookMeta.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookMeta)

	conditions := []*github_com_johnnyeven_libtools_sqlx_builder.Condition{}

	for _, fieldName := range bookMeta.IndexFieldNames() {
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

func (bookMeta *BookMeta) PrimaryKey() github_com_johnnyeven_libtools_sqlx.FieldNames {
	return github_com_johnnyeven_libtools_sqlx.FieldNames{"ID"}
}
func (bookMeta *BookMeta) Indexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{
		"I_author_status": github_com_johnnyeven_libtools_sqlx.FieldNames{"UserID", "Status"},
		"I_category":      github_com_johnnyeven_libtools_sqlx.FieldNames{"CategoryKey", "Status", "Selected"},
	}
}
func (bookMeta *BookMeta) UniqueIndexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{"U_book_id": github_com_johnnyeven_libtools_sqlx.FieldNames{"BookID", "Enabled"}}
}
func (bookMeta *BookMeta) Comments() map[string]string {
	return map[string]string{
		"BookID":       "业务ID",
		"BookLanguage": "文档语言",
		"CategoryKey":  "类别ID",
		"CodeLanguage": "代码语言",
		"Comment":      "简介",
		"CoverKey":     "封面图片key",
		"CreateTime":   "",
		"Enabled":      "",
		"ID":           "",
		"Selected":     "是否精选",
		"Status":       "状态",
		"Title":        "标题",
		"UpdateTime":   "",
		"UserID":       "作者ID",
	}
}

func (bookMeta *BookMeta) Create(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookMeta.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if bookMeta.CreateTime.IsZero() {
		bookMeta.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	bookMeta.UpdateTime = bookMeta.CreateTime

	stmt := bookMeta.D().
		Insert(bookMeta).
		Comment("BookMeta.Create")

	dbRet := db.Do(stmt)
	err := dbRet.Err()

	if err == nil {
		lastInsertID, _ := dbRet.LastInsertId()
		bookMeta.ID = uint64(lastInsertID)
	}

	return err
}

func (bookMeta *BookMeta) DeleteByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (err error) {
	table := bookMeta.T()

	stmt := table.Delete().
		Comment("BookMeta.DeleteByStruct").
		Where(bookMeta.ConditionByStruct())

	err = db.Do(stmt).Err()
	return
}

func (bookMeta *BookMeta) CreateOnDuplicateWithUpdateFields(db *github_com_johnnyeven_libtools_sqlx.DB, updateFields []string) error {
	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	bookMeta.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if bookMeta.CreateTime.IsZero() {
		bookMeta.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	bookMeta.UpdateTime = bookMeta.CreateTime

	table := bookMeta.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookMeta, updateFields...)

	delete(fieldValues, "ID")

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	m := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		m[field] = true
	}

	// fields of unique index can not update
	delete(m, "CreateTime")

	for _, fieldNames := range bookMeta.UniqueIndexes() {
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
		Comment("BookMeta.CreateOnDuplicateWithUpdateFields")

	return db.Do(stmt).Err()
}

func (bookMeta *BookMeta) FetchByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookMeta.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookMeta.T()
	stmt := table.Select().
		Comment("BookMeta.FetchByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookMeta.ID),
			table.F("Enabled").Eq(bookMeta.Enabled),
		))

	return db.Do(stmt).Scan(bookMeta).Err()
}

func (bookMeta *BookMeta) FetchByIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookMeta.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookMeta.T()
	stmt := table.Select().
		Comment("BookMeta.FetchByIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookMeta.ID),
			table.F("Enabled").Eq(bookMeta.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(bookMeta).Err()
}

func (bookMeta *BookMeta) DeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookMeta.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookMeta.T()
	stmt := table.Delete().
		Comment("BookMeta.DeleteByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookMeta.ID),
			table.F("Enabled").Eq(bookMeta.Enabled),
		))

	return db.Do(stmt).Scan(bookMeta).Err()
}

func (bookMeta *BookMeta) UpdateByIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	bookMeta.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookMeta.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("BookMeta.UpdateByIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookMeta.ID),
			table.F("Enabled").Eq(bookMeta.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookMeta)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return bookMeta.FetchByID(db)
	}
	return nil
}

func (bookMeta *BookMeta) UpdateByIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookMeta, zeroFields...)
	return bookMeta.UpdateByIDWithMap(db, fieldValues)
}

func (bookMeta *BookMeta) SoftDeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookMeta.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookMeta.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("BookMeta.SoftDeleteByID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookMeta.ID),
			table.F("Enabled").Eq(bookMeta.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookMeta)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return bookMeta.DeleteByID(db)
		}
		return err
	}
	return nil
}

func (bookMeta *BookMeta) FetchByBookID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookMeta.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookMeta.T()
	stmt := table.Select().
		Comment("BookMeta.FetchByBookID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookID").Eq(bookMeta.BookID),
			table.F("Enabled").Eq(bookMeta.Enabled),
		))

	return db.Do(stmt).Scan(bookMeta).Err()
}

func (bookMeta *BookMeta) FetchByBookIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookMeta.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookMeta.T()
	stmt := table.Select().
		Comment("BookMeta.FetchByBookIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookID").Eq(bookMeta.BookID),
			table.F("Enabled").Eq(bookMeta.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(bookMeta).Err()
}

func (bookMeta *BookMeta) DeleteByBookID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookMeta.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookMeta.T()
	stmt := table.Delete().
		Comment("BookMeta.DeleteByBookID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookID").Eq(bookMeta.BookID),
			table.F("Enabled").Eq(bookMeta.Enabled),
		))

	return db.Do(stmt).Scan(bookMeta).Err()
}

func (bookMeta *BookMeta) UpdateByBookIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	bookMeta.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookMeta.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("BookMeta.UpdateByBookIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookID").Eq(bookMeta.BookID),
			table.F("Enabled").Eq(bookMeta.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookMeta)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return bookMeta.FetchByBookID(db)
	}
	return nil
}

func (bookMeta *BookMeta) UpdateByBookIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookMeta, zeroFields...)
	return bookMeta.UpdateByBookIDWithMap(db, fieldValues)
}

func (bookMeta *BookMeta) SoftDeleteByBookID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookMeta.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookMeta.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("BookMeta.SoftDeleteByBookID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookID").Eq(bookMeta.BookID),
			table.F("Enabled").Eq(bookMeta.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookMeta)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return bookMeta.DeleteByBookID(db)
		}
		return err
	}
	return nil
}

type BookMetaList []BookMeta

// deprecated
func (bookMetaList *BookMetaList) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (count int32, err error) {
	*bookMetaList, count, err = (&BookMeta{}).FetchList(db, size, offset, conditions...)
	return
}

func (bookMeta *BookMeta) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (bookMetaList BookMetaList, count int32, err error) {
	bookMetaList = BookMetaList{}

	table := bookMeta.T()

	condition := github_com_johnnyeven_libtools_sqlx_builder.And(conditions...)

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookMeta.FetchList").
		Where(condition)

	errForCount := db.Do(stmt.For(github_com_johnnyeven_libtools_sqlx_builder.Count(github_com_johnnyeven_libtools_sqlx_builder.Star()))).Scan(&count).Err()
	if errForCount != nil {
		err = errForCount
		return
	}

	stmt = stmt.Limit(size).Offset(offset)

	stmt = stmt.OrderDescBy(table.F("CreateTime"))

	err = db.Do(stmt).Scan(&bookMetaList).Err()

	return
}

func (bookMeta *BookMeta) List(db *github_com_johnnyeven_libtools_sqlx.DB, condition *github_com_johnnyeven_libtools_sqlx_builder.Condition) (bookMetaList BookMetaList, err error) {
	bookMetaList = BookMetaList{}

	table := bookMeta.T()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookMeta.List").
		Where(condition)

	err = db.Do(stmt).Scan(&bookMetaList).Err()

	return
}

func (bookMeta *BookMeta) ListByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (bookMetaList BookMetaList, err error) {
	bookMetaList = BookMetaList{}

	table := bookMeta.T()

	condition := bookMeta.ConditionByStruct()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookMeta.ListByStruct").
		Where(condition)

	err = db.Do(stmt).Scan(&bookMetaList).Err()

	return
}

// deprecated
func (bookMetaList *BookMetaList) BatchFetchByBookIDList(db *github_com_johnnyeven_libtools_sqlx.DB, bookIDList []uint64) (err error) {
	*bookMetaList, err = (&BookMeta{}).BatchFetchByBookIDList(db, bookIDList)
	return
}

func (bookMeta *BookMeta) BatchFetchByBookIDList(db *github_com_johnnyeven_libtools_sqlx.DB, bookIDList []uint64) (bookMetaList BookMetaList, err error) {
	if len(bookIDList) == 0 {
		return BookMetaList{}, nil
	}

	table := bookMeta.T()

	condition := table.F("BookID").In(bookIDList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookMeta.BatchFetchByBookIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookMetaList).Err()

	return
}

// deprecated
func (bookMetaList *BookMetaList) BatchFetchByCategoryKeyList(db *github_com_johnnyeven_libtools_sqlx.DB, categoryKeyList []string) (err error) {
	*bookMetaList, err = (&BookMeta{}).BatchFetchByCategoryKeyList(db, categoryKeyList)
	return
}

func (bookMeta *BookMeta) BatchFetchByCategoryKeyList(db *github_com_johnnyeven_libtools_sqlx.DB, categoryKeyList []string) (bookMetaList BookMetaList, err error) {
	if len(categoryKeyList) == 0 {
		return BookMetaList{}, nil
	}

	table := bookMeta.T()

	condition := table.F("CategoryKey").In(categoryKeyList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookMeta.BatchFetchByCategoryKeyList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookMetaList).Err()

	return
}

// deprecated
func (bookMetaList *BookMetaList) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (err error) {
	*bookMetaList, err = (&BookMeta{}).BatchFetchByIDList(db, idList)
	return
}

func (bookMeta *BookMeta) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (bookMetaList BookMetaList, err error) {
	if len(idList) == 0 {
		return BookMetaList{}, nil
	}

	table := bookMeta.T()

	condition := table.F("ID").In(idList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookMeta.BatchFetchByIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookMetaList).Err()

	return
}

// deprecated
func (bookMetaList *BookMetaList) BatchFetchBySelectedList(db *github_com_johnnyeven_libtools_sqlx.DB, selectedList []github_com_johnnyeven_libtools_courier_enumeration.Bool) (err error) {
	*bookMetaList, err = (&BookMeta{}).BatchFetchBySelectedList(db, selectedList)
	return
}

func (bookMeta *BookMeta) BatchFetchBySelectedList(db *github_com_johnnyeven_libtools_sqlx.DB, selectedList []github_com_johnnyeven_libtools_courier_enumeration.Bool) (bookMetaList BookMetaList, err error) {
	if len(selectedList) == 0 {
		return BookMetaList{}, nil
	}

	table := bookMeta.T()

	condition := table.F("Selected").In(selectedList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookMeta.BatchFetchBySelectedList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookMetaList).Err()

	return
}

// deprecated
func (bookMetaList *BookMetaList) BatchFetchByStatusList(db *github_com_johnnyeven_libtools_sqlx.DB, statusList []github_com_in2_store_service_in2_book_constants_types.BookStatus) (err error) {
	*bookMetaList, err = (&BookMeta{}).BatchFetchByStatusList(db, statusList)
	return
}

func (bookMeta *BookMeta) BatchFetchByStatusList(db *github_com_johnnyeven_libtools_sqlx.DB, statusList []github_com_in2_store_service_in2_book_constants_types.BookStatus) (bookMetaList BookMetaList, err error) {
	if len(statusList) == 0 {
		return BookMetaList{}, nil
	}

	table := bookMeta.T()

	condition := table.F("Status").In(statusList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookMeta.BatchFetchByStatusList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookMetaList).Err()

	return
}

// deprecated
func (bookMetaList *BookMetaList) BatchFetchByUserIDList(db *github_com_johnnyeven_libtools_sqlx.DB, userIDList []uint64) (err error) {
	*bookMetaList, err = (&BookMeta{}).BatchFetchByUserIDList(db, userIDList)
	return
}

func (bookMeta *BookMeta) BatchFetchByUserIDList(db *github_com_johnnyeven_libtools_sqlx.DB, userIDList []uint64) (bookMetaList BookMetaList, err error) {
	if len(userIDList) == 0 {
		return BookMetaList{}, nil
	}

	table := bookMeta.T()

	condition := table.F("UserID").In(userIDList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookMeta.BatchFetchByUserIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookMetaList).Err()

	return
}
