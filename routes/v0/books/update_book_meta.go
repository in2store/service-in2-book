package books

import (
	"context"
	"github.com/in2store/service-in2-book/constants/errors"
	"github.com/in2store/service-in2-book/global"
	"github.com/in2store/service-in2-book/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(UpdateBookMeta{}))
}

// 更新文档元数据
type UpdateBookMeta struct {
	httpx.MethodPatch
	// 书籍ID
	BookID uint64                       `name:"bookID,string" in:"path"`
	Body   modules.UpdateBookMetaParams `name:"body" in:"body"`
}

func (req UpdateBookMeta) Path() string {
	return "/:bookID/meta"
}

func (req UpdateBookMeta) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	result, err = modules.UpdateBookMeta(req.BookID, req.Body, db, false)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, errors.BookNotFound
		}
		logrus.Errorf("[UpdateBookMeta] modules.UpdateBookMeta err: %v, request: %+v", err, req)
	}
	return
}
