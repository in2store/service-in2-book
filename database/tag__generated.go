package database

import (
	fmt "fmt"
	time "time"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
	github_com_johnnyeven_libtools_sqlx "github.com/johnnyeven/libtools/sqlx"
	github_com_johnnyeven_libtools_sqlx_builder "github.com/johnnyeven/libtools/sqlx/builder"
	github_com_johnnyeven_libtools_timelib "github.com/johnnyeven/libtools/timelib"
)

var TagTable *github_com_johnnyeven_libtools_sqlx_builder.Table

func init() {
	TagTable = DBIn2Book.Register(&Tag{})
}

func (tag *Tag) D() *github_com_johnnyeven_libtools_sqlx.Database {
	return DBIn2Book
}

func (tag *Tag) T() *github_com_johnnyeven_libtools_sqlx_builder.Table {
	return TagTable
}

func (tag *Tag) TableName() string {
	return "t_tag"
}

type TagFields struct {
	ID         *github_com_johnnyeven_libtools_sqlx_builder.Column
	TagID      *github_com_johnnyeven_libtools_sqlx_builder.Column
	Name       *github_com_johnnyeven_libtools_sqlx_builder.Column
	Heat       *github_com_johnnyeven_libtools_sqlx_builder.Column
	CreateTime *github_com_johnnyeven_libtools_sqlx_builder.Column
	UpdateTime *github_com_johnnyeven_libtools_sqlx_builder.Column
	Enabled    *github_com_johnnyeven_libtools_sqlx_builder.Column
}

var TagField = struct {
	ID         string
	TagID      string
	Name       string
	Heat       string
	CreateTime string
	UpdateTime string
	Enabled    string
}{
	ID:         "ID",
	TagID:      "TagID",
	Name:       "Name",
	Heat:       "Heat",
	CreateTime: "CreateTime",
	UpdateTime: "UpdateTime",
	Enabled:    "Enabled",
}

func (tag *Tag) Fields() *TagFields {
	table := tag.T()

	return &TagFields{
		ID:         table.F(TagField.ID),
		TagID:      table.F(TagField.TagID),
		Name:       table.F(TagField.Name),
		Heat:       table.F(TagField.Heat),
		CreateTime: table.F(TagField.CreateTime),
		UpdateTime: table.F(TagField.UpdateTime),
		Enabled:    table.F(TagField.Enabled),
	}
}

func (tag *Tag) IndexFieldNames() []string {
	return []string{"Heat", "ID", "TagID"}
}

func (tag *Tag) ConditionByStruct() *github_com_johnnyeven_libtools_sqlx_builder.Condition {
	table := tag.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(tag)

	conditions := []*github_com_johnnyeven_libtools_sqlx_builder.Condition{}

	for _, fieldName := range tag.IndexFieldNames() {
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

func (tag *Tag) PrimaryKey() github_com_johnnyeven_libtools_sqlx.FieldNames {
	return github_com_johnnyeven_libtools_sqlx.FieldNames{"ID"}
}
func (tag *Tag) Indexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{"I_heat": github_com_johnnyeven_libtools_sqlx.FieldNames{"Heat"}}
}
func (tag *Tag) UniqueIndexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{"U_tag_id": github_com_johnnyeven_libtools_sqlx.FieldNames{"TagID", "Enabled"}}
}
func (tag *Tag) Comments() map[string]string {
	return map[string]string{
		"CreateTime": "",
		"Enabled":    "",
		"Heat":       "热度",
		"ID":         "",
		"Name":       "名称",
		"TagID":      "业务ID",
		"UpdateTime": "",
	}
}

func (tag *Tag) Create(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	tag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if tag.CreateTime.IsZero() {
		tag.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	tag.UpdateTime = tag.CreateTime

	stmt := tag.D().
		Insert(tag).
		Comment("Tag.Create")

	dbRet := db.Do(stmt)
	err := dbRet.Err()

	if err == nil {
		lastInsertID, _ := dbRet.LastInsertId()
		tag.ID = uint64(lastInsertID)
	}

	return err
}

func (tag *Tag) DeleteByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (err error) {
	table := tag.T()

	stmt := table.Delete().
		Comment("Tag.DeleteByStruct").
		Where(tag.ConditionByStruct())

	err = db.Do(stmt).Err()
	return
}

func (tag *Tag) CreateOnDuplicateWithUpdateFields(db *github_com_johnnyeven_libtools_sqlx.DB, updateFields []string) error {
	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	tag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if tag.CreateTime.IsZero() {
		tag.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	tag.UpdateTime = tag.CreateTime

	table := tag.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(tag, updateFields...)

	delete(fieldValues, "ID")

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	m := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		m[field] = true
	}

	// fields of unique index can not update
	delete(m, "CreateTime")

	for _, fieldNames := range tag.UniqueIndexes() {
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
		Comment("Tag.CreateOnDuplicateWithUpdateFields")

	return db.Do(stmt).Err()
}

func (tag *Tag) FetchByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	tag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := tag.T()
	stmt := table.Select().
		Comment("Tag.FetchByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(tag.ID),
			table.F("Enabled").Eq(tag.Enabled),
		))

	return db.Do(stmt).Scan(tag).Err()
}

func (tag *Tag) FetchByIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	tag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := tag.T()
	stmt := table.Select().
		Comment("Tag.FetchByIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(tag.ID),
			table.F("Enabled").Eq(tag.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(tag).Err()
}

func (tag *Tag) DeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	tag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := tag.T()
	stmt := table.Delete().
		Comment("Tag.DeleteByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(tag.ID),
			table.F("Enabled").Eq(tag.Enabled),
		))

	return db.Do(stmt).Scan(tag).Err()
}

func (tag *Tag) UpdateByIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	tag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := tag.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("Tag.UpdateByIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(tag.ID),
			table.F("Enabled").Eq(tag.Enabled),
		))

	dbRet := db.Do(stmt).Scan(tag)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return tag.FetchByID(db)
	}
	return nil
}

func (tag *Tag) UpdateByIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(tag, zeroFields...)
	return tag.UpdateByIDWithMap(db, fieldValues)
}

func (tag *Tag) SoftDeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	tag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := tag.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("Tag.SoftDeleteByID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(tag.ID),
			table.F("Enabled").Eq(tag.Enabled),
		))

	dbRet := db.Do(stmt).Scan(tag)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return tag.DeleteByID(db)
		}
		return err
	}
	return nil
}

func (tag *Tag) FetchByTagID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	tag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := tag.T()
	stmt := table.Select().
		Comment("Tag.FetchByTagID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("TagID").Eq(tag.TagID),
			table.F("Enabled").Eq(tag.Enabled),
		))

	return db.Do(stmt).Scan(tag).Err()
}

func (tag *Tag) FetchByTagIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	tag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := tag.T()
	stmt := table.Select().
		Comment("Tag.FetchByTagIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("TagID").Eq(tag.TagID),
			table.F("Enabled").Eq(tag.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(tag).Err()
}

func (tag *Tag) DeleteByTagID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	tag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := tag.T()
	stmt := table.Delete().
		Comment("Tag.DeleteByTagID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("TagID").Eq(tag.TagID),
			table.F("Enabled").Eq(tag.Enabled),
		))

	return db.Do(stmt).Scan(tag).Err()
}

func (tag *Tag) UpdateByTagIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	tag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := tag.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("Tag.UpdateByTagIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("TagID").Eq(tag.TagID),
			table.F("Enabled").Eq(tag.Enabled),
		))

	dbRet := db.Do(stmt).Scan(tag)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return tag.FetchByTagID(db)
	}
	return nil
}

func (tag *Tag) UpdateByTagIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(tag, zeroFields...)
	return tag.UpdateByTagIDWithMap(db, fieldValues)
}

func (tag *Tag) SoftDeleteByTagID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	tag.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := tag.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("Tag.SoftDeleteByTagID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("TagID").Eq(tag.TagID),
			table.F("Enabled").Eq(tag.Enabled),
		))

	dbRet := db.Do(stmt).Scan(tag)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return tag.DeleteByTagID(db)
		}
		return err
	}
	return nil
}

type TagList []Tag

// deprecated
func (tagList *TagList) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (count int32, err error) {
	*tagList, count, err = (&Tag{}).FetchList(db, size, offset, conditions...)
	return
}

func (tag *Tag) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (tagList TagList, count int32, err error) {
	tagList = TagList{}

	table := tag.T()

	condition := github_com_johnnyeven_libtools_sqlx_builder.And(conditions...)

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Tag.FetchList").
		Where(condition)

	errForCount := db.Do(stmt.For(github_com_johnnyeven_libtools_sqlx_builder.Count(github_com_johnnyeven_libtools_sqlx_builder.Star()))).Scan(&count).Err()
	if errForCount != nil {
		err = errForCount
		return
	}

	stmt = stmt.Limit(size).Offset(offset)

	stmt = stmt.OrderDescBy(table.F("CreateTime"))

	err = db.Do(stmt).Scan(&tagList).Err()

	return
}

func (tag *Tag) List(db *github_com_johnnyeven_libtools_sqlx.DB, condition *github_com_johnnyeven_libtools_sqlx_builder.Condition) (tagList TagList, err error) {
	tagList = TagList{}

	table := tag.T()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Tag.List").
		Where(condition)

	err = db.Do(stmt).Scan(&tagList).Err()

	return
}

func (tag *Tag) ListByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (tagList TagList, err error) {
	tagList = TagList{}

	table := tag.T()

	condition := tag.ConditionByStruct()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Tag.ListByStruct").
		Where(condition)

	err = db.Do(stmt).Scan(&tagList).Err()

	return
}

// deprecated
func (tagList *TagList) BatchFetchByHeatList(db *github_com_johnnyeven_libtools_sqlx.DB, heatList []uint32) (err error) {
	*tagList, err = (&Tag{}).BatchFetchByHeatList(db, heatList)
	return
}

func (tag *Tag) BatchFetchByHeatList(db *github_com_johnnyeven_libtools_sqlx.DB, heatList []uint32) (tagList TagList, err error) {
	if len(heatList) == 0 {
		return TagList{}, nil
	}

	table := tag.T()

	condition := table.F("Heat").In(heatList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Tag.BatchFetchByHeatList").
		Where(condition)

	err = db.Do(stmt).Scan(&tagList).Err()

	return
}

// deprecated
func (tagList *TagList) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (err error) {
	*tagList, err = (&Tag{}).BatchFetchByIDList(db, idList)
	return
}

func (tag *Tag) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (tagList TagList, err error) {
	if len(idList) == 0 {
		return TagList{}, nil
	}

	table := tag.T()

	condition := table.F("ID").In(idList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Tag.BatchFetchByIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&tagList).Err()

	return
}

// deprecated
func (tagList *TagList) BatchFetchByTagIDList(db *github_com_johnnyeven_libtools_sqlx.DB, tagIDList []uint64) (err error) {
	*tagList, err = (&Tag{}).BatchFetchByTagIDList(db, tagIDList)
	return
}

func (tag *Tag) BatchFetchByTagIDList(db *github_com_johnnyeven_libtools_sqlx.DB, tagIDList []uint64) (tagList TagList, err error) {
	if len(tagIDList) == 0 {
		return TagList{}, nil
	}

	table := tag.T()

	condition := table.F("TagID").In(tagIDList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Tag.BatchFetchByTagIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&tagList).Err()

	return
}
