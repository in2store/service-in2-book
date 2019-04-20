package database

import (
	fmt "fmt"
	time "time"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
	github_com_johnnyeven_libtools_sqlx "github.com/johnnyeven/libtools/sqlx"
	github_com_johnnyeven_libtools_sqlx_builder "github.com/johnnyeven/libtools/sqlx/builder"
	github_com_johnnyeven_libtools_timelib "github.com/johnnyeven/libtools/timelib"
)

var BookRepoTable *github_com_johnnyeven_libtools_sqlx_builder.Table

func init() {
	BookRepoTable = DBIn2Book.Register(&BookRepo{})
}

func (bookRepo *BookRepo) D() *github_com_johnnyeven_libtools_sqlx.Database {
	return DBIn2Book
}

func (bookRepo *BookRepo) T() *github_com_johnnyeven_libtools_sqlx_builder.Table {
	return BookRepoTable
}

func (bookRepo *BookRepo) TableName() string {
	return "t_book_repo"
}

type BookRepoFields struct {
	ID             *github_com_johnnyeven_libtools_sqlx_builder.Column
	BookRepoID     *github_com_johnnyeven_libtools_sqlx_builder.Column
	BookID         *github_com_johnnyeven_libtools_sqlx_builder.Column
	ChannelID      *github_com_johnnyeven_libtools_sqlx_builder.Column
	EntryURL       *github_com_johnnyeven_libtools_sqlx_builder.Column
	RepoFullName   *github_com_johnnyeven_libtools_sqlx_builder.Column
	RepoBranchName *github_com_johnnyeven_libtools_sqlx_builder.Column
	SummaryPath    *github_com_johnnyeven_libtools_sqlx_builder.Column
	CreateTime     *github_com_johnnyeven_libtools_sqlx_builder.Column
	UpdateTime     *github_com_johnnyeven_libtools_sqlx_builder.Column
	Enabled        *github_com_johnnyeven_libtools_sqlx_builder.Column
}

var BookRepoField = struct {
	ID             string
	BookRepoID     string
	BookID         string
	ChannelID      string
	EntryURL       string
	RepoFullName   string
	RepoBranchName string
	SummaryPath    string
	CreateTime     string
	UpdateTime     string
	Enabled        string
}{
	ID:             "ID",
	BookRepoID:     "BookRepoID",
	BookID:         "BookID",
	ChannelID:      "ChannelID",
	EntryURL:       "EntryURL",
	RepoFullName:   "RepoFullName",
	RepoBranchName: "RepoBranchName",
	SummaryPath:    "SummaryPath",
	CreateTime:     "CreateTime",
	UpdateTime:     "UpdateTime",
	Enabled:        "Enabled",
}

func (bookRepo *BookRepo) Fields() *BookRepoFields {
	table := bookRepo.T()

	return &BookRepoFields{
		ID:             table.F(BookRepoField.ID),
		BookRepoID:     table.F(BookRepoField.BookRepoID),
		BookID:         table.F(BookRepoField.BookID),
		ChannelID:      table.F(BookRepoField.ChannelID),
		EntryURL:       table.F(BookRepoField.EntryURL),
		RepoFullName:   table.F(BookRepoField.RepoFullName),
		RepoBranchName: table.F(BookRepoField.RepoBranchName),
		SummaryPath:    table.F(BookRepoField.SummaryPath),
		CreateTime:     table.F(BookRepoField.CreateTime),
		UpdateTime:     table.F(BookRepoField.UpdateTime),
		Enabled:        table.F(BookRepoField.Enabled),
	}
}

func (bookRepo *BookRepo) IndexFieldNames() []string {
	return []string{"BookID", "BookRepoID", "ID"}
}

func (bookRepo *BookRepo) ConditionByStruct() *github_com_johnnyeven_libtools_sqlx_builder.Condition {
	table := bookRepo.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookRepo)

	conditions := []*github_com_johnnyeven_libtools_sqlx_builder.Condition{}

	for _, fieldName := range bookRepo.IndexFieldNames() {
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

func (bookRepo *BookRepo) PrimaryKey() github_com_johnnyeven_libtools_sqlx.FieldNames {
	return github_com_johnnyeven_libtools_sqlx.FieldNames{"ID"}
}
func (bookRepo *BookRepo) UniqueIndexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{
		"U_book_id":      github_com_johnnyeven_libtools_sqlx.FieldNames{"BookID", "Enabled"},
		"U_book_repo_id": github_com_johnnyeven_libtools_sqlx.FieldNames{"BookRepoID", "Enabled"},
	}
}
func (bookRepo *BookRepo) Comments() map[string]string {
	return map[string]string{
		"BookID":         "书籍ID",
		"BookRepoID":     "业务ID",
		"ChannelID":      "通道ID",
		"CreateTime":     "",
		"Enabled":        "",
		"EntryURL":       "入口地址",
		"ID":             "",
		"RepoBranchName": "代码库分支",
		"RepoFullName":   "代码库全名",
		"SummaryPath":    "Summary文件相对地址",
		"UpdateTime":     "",
	}
}

func (bookRepo *BookRepo) Create(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if bookRepo.CreateTime.IsZero() {
		bookRepo.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	bookRepo.UpdateTime = bookRepo.CreateTime

	stmt := bookRepo.D().
		Insert(bookRepo).
		Comment("BookRepo.Create")

	dbRet := db.Do(stmt)
	err := dbRet.Err()

	if err == nil {
		lastInsertID, _ := dbRet.LastInsertId()
		bookRepo.ID = uint64(lastInsertID)
	}

	return err
}

func (bookRepo *BookRepo) DeleteByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (err error) {
	table := bookRepo.T()

	stmt := table.Delete().
		Comment("BookRepo.DeleteByStruct").
		Where(bookRepo.ConditionByStruct())

	err = db.Do(stmt).Err()
	return
}

func (bookRepo *BookRepo) CreateOnDuplicateWithUpdateFields(db *github_com_johnnyeven_libtools_sqlx.DB, updateFields []string) error {
	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if bookRepo.CreateTime.IsZero() {
		bookRepo.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	bookRepo.UpdateTime = bookRepo.CreateTime

	table := bookRepo.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookRepo, updateFields...)

	delete(fieldValues, "ID")

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	m := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		m[field] = true
	}

	// fields of unique index can not update
	delete(m, "CreateTime")

	for _, fieldNames := range bookRepo.UniqueIndexes() {
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
		Comment("BookRepo.CreateOnDuplicateWithUpdateFields")

	return db.Do(stmt).Err()
}

func (bookRepo *BookRepo) FetchByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()
	stmt := table.Select().
		Comment("BookRepo.FetchByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookRepo.ID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		))

	return db.Do(stmt).Scan(bookRepo).Err()
}

func (bookRepo *BookRepo) FetchByIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()
	stmt := table.Select().
		Comment("BookRepo.FetchByIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookRepo.ID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(bookRepo).Err()
}

func (bookRepo *BookRepo) DeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()
	stmt := table.Delete().
		Comment("BookRepo.DeleteByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookRepo.ID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		))

	return db.Do(stmt).Scan(bookRepo).Err()
}

func (bookRepo *BookRepo) UpdateByIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("BookRepo.UpdateByIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookRepo.ID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookRepo)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return bookRepo.FetchByID(db)
	}
	return nil
}

func (bookRepo *BookRepo) UpdateByIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookRepo, zeroFields...)
	return bookRepo.UpdateByIDWithMap(db, fieldValues)
}

func (bookRepo *BookRepo) SoftDeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("BookRepo.SoftDeleteByID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(bookRepo.ID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookRepo)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return bookRepo.DeleteByID(db)
		}
		return err
	}
	return nil
}

func (bookRepo *BookRepo) FetchByBookID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()
	stmt := table.Select().
		Comment("BookRepo.FetchByBookID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookID").Eq(bookRepo.BookID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		))

	return db.Do(stmt).Scan(bookRepo).Err()
}

func (bookRepo *BookRepo) FetchByBookIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()
	stmt := table.Select().
		Comment("BookRepo.FetchByBookIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookID").Eq(bookRepo.BookID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(bookRepo).Err()
}

func (bookRepo *BookRepo) DeleteByBookID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()
	stmt := table.Delete().
		Comment("BookRepo.DeleteByBookID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookID").Eq(bookRepo.BookID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		))

	return db.Do(stmt).Scan(bookRepo).Err()
}

func (bookRepo *BookRepo) UpdateByBookIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("BookRepo.UpdateByBookIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookID").Eq(bookRepo.BookID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookRepo)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return bookRepo.FetchByBookID(db)
	}
	return nil
}

func (bookRepo *BookRepo) UpdateByBookIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookRepo, zeroFields...)
	return bookRepo.UpdateByBookIDWithMap(db, fieldValues)
}

func (bookRepo *BookRepo) SoftDeleteByBookID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("BookRepo.SoftDeleteByBookID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookID").Eq(bookRepo.BookID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookRepo)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return bookRepo.DeleteByBookID(db)
		}
		return err
	}
	return nil
}

func (bookRepo *BookRepo) FetchByBookRepoID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()
	stmt := table.Select().
		Comment("BookRepo.FetchByBookRepoID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookRepoID").Eq(bookRepo.BookRepoID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		))

	return db.Do(stmt).Scan(bookRepo).Err()
}

func (bookRepo *BookRepo) FetchByBookRepoIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()
	stmt := table.Select().
		Comment("BookRepo.FetchByBookRepoIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookRepoID").Eq(bookRepo.BookRepoID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(bookRepo).Err()
}

func (bookRepo *BookRepo) DeleteByBookRepoID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()
	stmt := table.Delete().
		Comment("BookRepo.DeleteByBookRepoID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookRepoID").Eq(bookRepo.BookRepoID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		))

	return db.Do(stmt).Scan(bookRepo).Err()
}

func (bookRepo *BookRepo) UpdateByBookRepoIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("BookRepo.UpdateByBookRepoIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookRepoID").Eq(bookRepo.BookRepoID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookRepo)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return bookRepo.FetchByBookRepoID(db)
	}
	return nil
}

func (bookRepo *BookRepo) UpdateByBookRepoIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(bookRepo, zeroFields...)
	return bookRepo.UpdateByBookRepoIDWithMap(db, fieldValues)
}

func (bookRepo *BookRepo) SoftDeleteByBookRepoID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	bookRepo.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := bookRepo.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("BookRepo.SoftDeleteByBookRepoID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("BookRepoID").Eq(bookRepo.BookRepoID),
			table.F("Enabled").Eq(bookRepo.Enabled),
		))

	dbRet := db.Do(stmt).Scan(bookRepo)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return bookRepo.DeleteByBookRepoID(db)
		}
		return err
	}
	return nil
}

type BookRepoList []BookRepo

// deprecated
func (bookRepoList *BookRepoList) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (count int32, err error) {
	*bookRepoList, count, err = (&BookRepo{}).FetchList(db, size, offset, conditions...)
	return
}

func (bookRepo *BookRepo) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (bookRepoList BookRepoList, count int32, err error) {
	bookRepoList = BookRepoList{}

	table := bookRepo.T()

	condition := github_com_johnnyeven_libtools_sqlx_builder.And(conditions...)

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookRepo.FetchList").
		Where(condition)

	errForCount := db.Do(stmt.For(github_com_johnnyeven_libtools_sqlx_builder.Count(github_com_johnnyeven_libtools_sqlx_builder.Star()))).Scan(&count).Err()
	if errForCount != nil {
		err = errForCount
		return
	}

	stmt = stmt.Limit(size).Offset(offset)

	stmt = stmt.OrderDescBy(table.F("CreateTime"))

	err = db.Do(stmt).Scan(&bookRepoList).Err()

	return
}

func (bookRepo *BookRepo) List(db *github_com_johnnyeven_libtools_sqlx.DB, condition *github_com_johnnyeven_libtools_sqlx_builder.Condition) (bookRepoList BookRepoList, err error) {
	bookRepoList = BookRepoList{}

	table := bookRepo.T()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookRepo.List").
		Where(condition)

	err = db.Do(stmt).Scan(&bookRepoList).Err()

	return
}

func (bookRepo *BookRepo) ListByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (bookRepoList BookRepoList, err error) {
	bookRepoList = BookRepoList{}

	table := bookRepo.T()

	condition := bookRepo.ConditionByStruct()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookRepo.ListByStruct").
		Where(condition)

	err = db.Do(stmt).Scan(&bookRepoList).Err()

	return
}

// deprecated
func (bookRepoList *BookRepoList) BatchFetchByBookIDList(db *github_com_johnnyeven_libtools_sqlx.DB, bookIDList []uint64) (err error) {
	*bookRepoList, err = (&BookRepo{}).BatchFetchByBookIDList(db, bookIDList)
	return
}

func (bookRepo *BookRepo) BatchFetchByBookIDList(db *github_com_johnnyeven_libtools_sqlx.DB, bookIDList []uint64) (bookRepoList BookRepoList, err error) {
	if len(bookIDList) == 0 {
		return BookRepoList{}, nil
	}

	table := bookRepo.T()

	condition := table.F("BookID").In(bookIDList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookRepo.BatchFetchByBookIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookRepoList).Err()

	return
}

// deprecated
func (bookRepoList *BookRepoList) BatchFetchByBookRepoIDList(db *github_com_johnnyeven_libtools_sqlx.DB, bookRepoIDList []uint64) (err error) {
	*bookRepoList, err = (&BookRepo{}).BatchFetchByBookRepoIDList(db, bookRepoIDList)
	return
}

func (bookRepo *BookRepo) BatchFetchByBookRepoIDList(db *github_com_johnnyeven_libtools_sqlx.DB, bookRepoIDList []uint64) (bookRepoList BookRepoList, err error) {
	if len(bookRepoIDList) == 0 {
		return BookRepoList{}, nil
	}

	table := bookRepo.T()

	condition := table.F("BookRepoID").In(bookRepoIDList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookRepo.BatchFetchByBookRepoIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookRepoList).Err()

	return
}

// deprecated
func (bookRepoList *BookRepoList) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (err error) {
	*bookRepoList, err = (&BookRepo{}).BatchFetchByIDList(db, idList)
	return
}

func (bookRepo *BookRepo) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (bookRepoList BookRepoList, err error) {
	if len(idList) == 0 {
		return BookRepoList{}, nil
	}

	table := bookRepo.T()

	condition := table.F("ID").In(idList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("BookRepo.BatchFetchByIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&bookRepoList).Err()

	return
}
