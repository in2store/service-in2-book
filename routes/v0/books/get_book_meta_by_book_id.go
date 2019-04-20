package books

import (
	"context"
	"github.com/in2store/service-in2-book/global"
	"github.com/in2store/service-in2-book/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetBookMetaByBookID{}))
}

// 根据书籍ID获取书籍元数据
type GetBookMetaByBookID struct {
	httpx.MethodGet
	// 书籍ID
	BookID uint64 `name:"bookID,string" in:"path"`
}

func (req GetBookMetaByBookID) Path() string {
	return "/:bookID/meta"
}

func (req GetBookMetaByBookID) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.SlaveDB.Get()
	result, err = modules.GetBookMetaByBookID(req.BookID, db, false)
	if err != nil {
		if !sqlx.DBErr(err).IsNotFound() {
			logrus.Errorf("[GetBookMetaByBookID] modules.GetBookMetaByBookID err: %v, request: %d", err, req.BookID)
		}
	}
	return
}
