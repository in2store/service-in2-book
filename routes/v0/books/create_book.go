package books

import (
	"context"
	"github.com/in2store/service-in2-book/constants/errors"
	"github.com/in2store/service-in2-book/database"
	"github.com/in2store/service-in2-book/global"
	"github.com/in2store/service-in2-book/modules"
	libModule "github.com/johnnyeven/eden-library/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(CreateBook{}))
}

type CreateBookBody struct {
	modules.CreateBookMetaParams
	modules.CreateBookRepoParams
}

// 创建书籍
type CreateBook struct {
	httpx.MethodPost
	Body CreateBookBody `name:"body" in:"body"`
}

func (req CreateBook) Path() string {
	return ""
}

func (req CreateBook) Output(ctx context.Context) (result interface{}, err error) {
	bookID, err := libModule.NewUniqueID(global.Config.ClientID)
	if err != nil {
		logrus.Errorf("[CreateBook] libModule.NewUniqueID err: %v", err)
		return nil, err
	}

	db := global.Config.MasterDB.Get()
	tx := sqlx.NewTasks(db)

	var meta *database.BookMeta
	createMeta := func(db *sqlx.DB) error {
		request := modules.CreateBookMetaParams{
			UserID:       req.Body.UserID,
			Title:        req.Body.Title,
			CategoryKey:  req.Body.CategoryKey,
			CoverKey:     req.Body.CoverKey,
			Comment:      req.Body.Comment,
			BookLanguage: req.Body.BookLanguage,
			CodeLanguage: req.Body.CodeLanguage,
		}
		meta, err = modules.CreateBookMeta(bookID, request, db)
		if err != nil {
			logrus.Errorf("[CreateBook] modules.CreateBookMeta err: %v, request: %+v", err, request)
			return err
		}
		return nil
	}
	createRepo := func(db *sqlx.DB) error {
		request := modules.CreateBookRepoParams{
			ChannelID:      req.Body.ChannelID,
			EntryURL:       req.Body.EntryURL,
			RepoFullName:   req.Body.RepoFullName,
			RepoBranchName: req.Body.RepoBranchName,
			SummaryPath:    req.Body.SummaryPath,
		}
		_, err := modules.CreateBookRepo(bookID, request, db)
		if err != nil {
			logrus.Errorf("[CreateBook] modules.CreateBookRepo err: %v, request: %+v", err, request)
			return err
		}
		return nil
	}

	tx = tx.With(createMeta, createRepo)
	err = tx.Do()
	if err != nil {
		if sqlx.DBErr(err).IsConflict() {
			return nil, errors.BookConflict
		}
		logrus.Errorf("[CreateBook] transaction err: %v", err)
		return nil, err
	}

	return meta, nil
}
