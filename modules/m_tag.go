package modules

import (
	"github.com/in2store/service-in2-book/constants/errors"
	"github.com/in2store/service-in2-book/database"
	"github.com/johnnyeven/libtools/sqlx"
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
