package modules

import (
	"github.com/in2store/service-in2-book/constants/errors"
	"github.com/in2store/service-in2-book/database"
	"github.com/johnnyeven/libtools/courier/enumeration"
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/johnnyeven/libtools/sqlx/builder"
)

func CreateTag(tagID uint64, tagName string, db *sqlx.DB) (result *database.Tag, err error) {
	result = &database.Tag{
		TagID: tagID,
		Name:  tagName,
	}
	err = result.Create(db)
	return
}

func SetBookTag(bookID, tagID uint64, db *sqlx.DB) error {
	tag := &database.Tag{
		TagID: tagID,
	}
	err := tag.FetchByTagID(db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return errors.TagNotFound
		}
		return err
	}
	bookTag := &database.BookTag{
		TagID:  tagID,
		BookID: bookID,
	}
	err = bookTag.Create(db)
	if err != nil {
		if sqlx.DBErr(err).IsConflict() {
			return errors.BookTagConflict
		}
		return err
	}
	return nil
}

func GetTagsByBookID(bookID uint64, db *sqlx.DB) (result database.TagList, err error) {
	bookTag := &database.BookTag{}
	bookTags, err := bookTag.BatchFetchByBookIDList(db, []uint64{bookID})
	if err != nil {
		return
	}

	tagIDs := make([]uint64, 0)
	for _, bookTag := range bookTags {
		tagIDs = append(tagIDs, bookTag.TagID)
	}
	tag := &database.Tag{}
	result, err = tag.BatchFetchByTagIDList(db, tagIDs)
	return
}

func GetTagsOrderByHeat(filterZereHeat bool, orderByHeat bool, db *sqlx.DB) (result database.TagList, err error) {
	tag := &database.Tag{}
	table := tag.T()

	condition := builder.And(table.F("Enabled").Eq(enumeration.BOOL__TRUE))

	if filterZereHeat {
		condition = builder.And(table.F("Heat").Neq(0))
	}

	stmt := table.
		Select().
		Comment("Tag.GetTagsOrderByHeat").
		Where(condition)

	if orderByHeat {
		stmt = stmt.OrderDescBy(table.F("Heat"))
	}

	err = db.Do(stmt).Scan(&result).Err()
	return
}
