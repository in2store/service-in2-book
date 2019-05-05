package database

import (
	fmt "fmt"
	time "time"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
	github_com_johnnyeven_libtools_sqlx "github.com/johnnyeven/libtools/sqlx"
	github_com_johnnyeven_libtools_sqlx_builder "github.com/johnnyeven/libtools/sqlx/builder"
	github_com_johnnyeven_libtools_timelib "github.com/johnnyeven/libtools/timelib"
)

var BookTagTable *github_com_johnnyeven_libtools_sqlx_builder.Table

func init() {
	BookTagTable = DBIn2Book.Register(&BookTag{})
}

func (bookTag *BookTag) D() *github_com_johnnyeven_libtools_sqlx.Database {
	return DBIn2Book
}

func (bookTag *BookTag) T() *github_com_johnnyeven_libtools_sqlx_builder.Table {
	return BookTagTable
}

func (bookTag *BookTag) TableName() string {
	return "t_book_tag"
}

type BookTagFields struct {
	ID         *github_com_johnnyeven_libtools_sqlx_builder.Column
	TagID      *github_com_johnnyeven_libtools_sqlx_builder.Column
	BookID     *github_com_johnnyeven_libtools_sqlx_builder.Column
	CreateTime *github_com_johnnyeven_libtools_sqlx_builder.Column
	UpdateTime *github_com_johnnyeven_libtools_sqlx_builder.Column
	Enabled    *github_com_johnnyeven_libtools_sqlx_builder.Column
}

var BookTagField = struct {
	ID         string
	TagID      string
	BookID     string
	CreateTime string
	UpdateTime string
	Enabled    string
}{
	ID:         "ID",
	TagID:      "TagID",
	BookID:     "BookID",
	CreateTime: "CreateTime",
	UpdateTime: "UpdateTime",
	Enabled:    "Enabled",
}

func (bookTag *BookTag) Fields() *BookTagFields {
	table := bookTag.T()

	return &BookTagFields{
		ID:         table.F(BookTagField.ID),
		TagID:      table.F(BookTagField.TagID),
		BookID:     table.F(BookTagField.BookID),
		CreateTime: table.F(BookTagField.CreateTime),
		UpdateTime: table.F(BookTagField.UpdateTime),
		Enabled:    table.F(BookTagField.Enabled),
	}
}

func (bookTag *BookTag) IndexFieldNames() []string {
	return []string{"BookID", "ID", "TagID"}
}

func (bookTag *BookTag) ConditionByStruct() *github_com_johnnyeven_libtools_sqlx_builder.Condition {
	table := bookTag.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookTag)

	conditions := []*github_com_johnnyeven_libtools_sqlx_builder.Condition{}

	for _, fieldName := range bookTag.IndexFieldNames() {
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

func (bookTag *BookTag) PrimaryKey() github_com_johnnyeven_libtools_sqlx.FieldNames {
	return github_com_johnnyeven_libtools_sqlx.FieldNames{"ID"}
}
func (bookTag *BookTag) Indexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{"I_tag": github_com_johnnyeven_libtools_sqlx.FieldNames{"TagID"}}
}
func (bookTag *BookTag) UniqueIndexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{"U_book_tag": github_com_johnnyeven_libtools_sqlx.FieldNames{"TagID", "BookID", "Enabled"}}
}
func (bookTag *BookTag) Comments() map[string]string {
	return map[string]string{
		"BookID":     "文档ID",
		"CreateTime": "",
		"Enabled":    "",
		"ID":         "",
		"TagID":      "标签ID",
		"UpdateTime": "",
	}
}

func (bookTag *BookTag) Create(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookTag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if bookTag.CreateTime.IsZero() {
		bookTag.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	bookTag.UpdateTime = bookTag.CreateTime

	stmt := bookTag.D().
		Insert(bookTag).
		Comment("BookTag.Create")

	dbRet := db.Do(stmt)
	err := dbRet.Err()

	if err == nil {
		lastInsertID, _ := dbRet.LastInsertId()
		bookTag.ID = uint64(lastInsertID)
	}

	return err
}

func (bookTag *BookTag) DeleteByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (err error) {
	table := bookTag.T()

	stmt := table.Delete().
		Comment("BookTag.DeleteByStruct").
		Where(bookTag.ConditionByStruct())

	err = db.Do(stmt).Err()
	return
}

func (bookTag *BookTag) CreateOnDuplicateWithUpdateFields(db *github_com_johnnyeven_libtools_sqlx.DB, updateFields []string) error {
	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	bookTag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if bookTag.CreateTime.IsZero() {
		bookTag.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	bookTag.UpdateTime = bookTag.CreateTime

	table := bookTag.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookTag, updateFields...)

	delete(fieldValues, "ID")

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	m := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		m[field] = true
	}

	// fields of unique index can not update
	delete(m, "CreateTime")

	for _, fieldNames := range bookTag.UniqueIndexes() {
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
		Comment("BookTag.CreateOnDuplicateWithUpdateFields")

	return db.Do(stmt).Err()
}

func (bookTag *BookTag) FetchByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookTag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookTag.T()
	stmt := table.Select().
		Comment("BookTag.FetchByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookTag.ID),
			table.F("Enabled").Eq(bookTag.Enabled),
		))

	return db.Do(stmt).Scan(bookTag).Err()
}

func (bookTag *BookTag) FetchByIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookTag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookTag.T()
	stmt := table.Select().
		Comment("BookTag.FetchByIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookTag.ID),
			table.F("Enabled").Eq(bookTag.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(bookTag).Err()
}

func (bookTag *BookTag) DeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookTag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookTag.T()
	stmt := table.Delete().
		Comment("BookTag.DeleteByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookTag.ID),
			table.F("Enabled").Eq(bookTag.Enabled),
		))

	return db.Do(stmt).Scan(bookTag).Err()
}

func (bookTag *BookTag) UpdateByIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	bookTag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookTag.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("BookTag.UpdateByIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookTag.ID),
			table.F("Enabled").Eq(bookTag.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookTag)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return bookTag.FetchByID(db)
	}
	return nil
}

func (bookTag *BookTag) UpdateByIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookTag, zeroFields...)
	return bookTag.UpdateByIDWithMap(db, fieldValues)
}

func (bookTag *BookTag) SoftDeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookTag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookTag.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("BookTag.SoftDeleteByID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookTag.ID),
			table.F("Enabled").Eq(bookTag.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookTag)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return bookTag.DeleteByID(db)
		}
		return err
	}
	return nil
}

func (bookTag *BookTag) FetchByTagIDAndBookID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookTag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookTag.T()
	stmt := table.Select().
		Comment("BookTag.FetchByTagIDAndBookID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("TagID").Eq(bookTag.TagID),
			table.F("BookID").Eq(bookTag.BookID),
			table.F("Enabled").Eq(bookTag.Enabled),
		))

	return db.Do(stmt).Scan(bookTag).Err()
}

func (bookTag *BookTag) FetchByTagIDAndBookIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookTag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookTag.T()
	stmt := table.Select().
		Comment("BookTag.FetchByTagIDAndBookIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("TagID").Eq(bookTag.TagID),
			table.F("BookID").Eq(bookTag.BookID),
			table.F("Enabled").Eq(bookTag.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(bookTag).Err()
}

func (bookTag *BookTag) DeleteByTagIDAndBookID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookTag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookTag.T()
	stmt := table.Delete().
		Comment("BookTag.DeleteByTagIDAndBookID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("TagID").Eq(bookTag.TagID),
			table.F("BookID").Eq(bookTag.BookID),
			table.F("Enabled").Eq(bookTag.Enabled),
		))

	return db.Do(stmt).Scan(bookTag).Err()
}

func (bookTag *BookTag) UpdateByTagIDAndBookIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	bookTag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookTag.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("BookTag.UpdateByTagIDAndBookIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("TagID").Eq(bookTag.TagID),
			table.F("BookID").Eq(bookTag.BookID),
			table.F("Enabled").Eq(bookTag.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookTag)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return bookTag.FetchByTagIDAndBookID(db)
	}
	return nil
}

func (bookTag *BookTag) UpdateByTagIDAndBookIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookTag, zeroFields...)
	return bookTag.UpdateByTagIDAndBookIDWithMap(db, fieldValues)
}

func (bookTag *BookTag) SoftDeleteByTagIDAndBookID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookTag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookTag.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("BookTag.SoftDeleteByTagIDAndBookID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("TagID").Eq(bookTag.TagID),
			table.F("BookID").Eq(bookTag.BookID),
			table.F("Enabled").Eq(bookTag.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookTag)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return bookTag.DeleteByTagIDAndBookID(db)
		}
		return err
	}
	return nil
}

type BookTagList []BookTag

// deprecated
func (bookTagList *BookTagList) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (count int32, err error) {
	*bookTagList, count, err = (&BookTag{}).FetchList(db, size, offset, conditions...)
	return
}

func (bookTag *BookTag) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (bookTagList BookTagList, count int32, err error) {
	bookTagList = BookTagList{}

	table := bookTag.T()

	condition := github_com_johnnyeven_libtools_sqlx_builder.And(conditions...)

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookTag.FetchList").
		Where(condition)

	errForCount := db.Do(stmt.For(github_com_johnnyeven_libtools_sqlx_builder.Count(github_com_johnnyeven_libtools_sqlx_builder.Star()))).Scan(&count).Err()
	if errForCount != nil {
		err = errForCount
		return
	}

	stmt = stmt.Limit(size).Offset(offset)

	stmt = stmt.OrderDescBy(table.F("CreateTime"))

	err = db.Do(stmt).Scan(&bookTagList).Err()

	return
}

func (bookTag *BookTag) List(db *github_com_johnnyeven_libtools_sqlx.DB, condition *github_com_johnnyeven_libtools_sqlx_builder.Condition) (bookTagList BookTagList, err error) {
	bookTagList = BookTagList{}

	table := bookTag.T()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookTag.List").
		Where(condition)

	err = db.Do(stmt).Scan(&bookTagList).Err()

	return
}

func (bookTag *BookTag) ListByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (bookTagList BookTagList, err error) {
	bookTagList = BookTagList{}

	table := bookTag.T()

	condition := bookTag.ConditionByStruct()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookTag.ListByStruct").
		Where(condition)

	err = db.Do(stmt).Scan(&bookTagList).Err()

	return
}

// deprecated
func (bookTagList *BookTagList) BatchFetchByBookIDList(db *github_com_johnnyeven_libtools_sqlx.DB, bookIDList []uint64) (err error) {
	*bookTagList, err = (&BookTag{}).BatchFetchByBookIDList(db, bookIDList)
	return
}

func (bookTag *BookTag) BatchFetchByBookIDList(db *github_com_johnnyeven_libtools_sqlx.DB, bookIDList []uint64) (bookTagList BookTagList, err error) {
	if len(bookIDList) == 0 {
		return BookTagList{}, nil
	}

	table := bookTag.T()

	condition := table.F("BookID").In(bookIDList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookTag.BatchFetchByBookIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookTagList).Err()

	return
}

// deprecated
func (bookTagList *BookTagList) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (err error) {
	*bookTagList, err = (&BookTag{}).BatchFetchByIDList(db, idList)
	return
}

func (bookTag *BookTag) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (bookTagList BookTagList, err error) {
	if len(idList) == 0 {
		return BookTagList{}, nil
	}

	table := bookTag.T()

	condition := table.F("ID").In(idList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookTag.BatchFetchByIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookTagList).Err()

	return
}

// deprecated
func (bookTagList *BookTagList) BatchFetchByTagIDList(db *github_com_johnnyeven_libtools_sqlx.DB, tagIDList []uint64) (err error) {
	*bookTagList, err = (&BookTag{}).BatchFetchByTagIDList(db, tagIDList)
	return
}

func (bookTag *BookTag) BatchFetchByTagIDList(db *github_com_johnnyeven_libtools_sqlx.DB, tagIDList []uint64) (bookTagList BookTagList, err error) {
	if len(tagIDList) == 0 {
		return BookTagList{}, nil
	}

	table := bookTag.T()

	condition := table.F("TagID").In(tagIDList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookTag.BatchFetchByTagIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookTagList).Err()

	return
}
