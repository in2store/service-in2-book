package books

import (
	"context"
	"github.com/in2store/service-in2-book/database"
	"github.com/in2store/service-in2-book/global"
	"github.com/in2store/service-in2-book/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/httplib"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetBooksMeta{}))
}

// 获取书籍元数据列表
type GetBooksMeta struct {
	httpx.MethodGet
	modules.GetBooksMetaRequest
	httplib.Pager
}

func (req GetBooksMeta) Path() string {
	return ""
}

type GetBooksMetaResult struct {
	Data  database.BookMetaList `json:"data"`
	Total int32                 `json:"total"`
}

func (req GetBooksMeta) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.SlaveDB.Get()
	resp, count, err := modules.GetBooksMeta(req.GetBooksMetaRequest, req.Size, req.Offset, db)
	if err != nil {
		logrus.Errorf("[GetBooksMeta] modules.GetBooksMeta err: %v, request: %+v", err, req)
		return nil, err
	}

	return GetBooksMetaResult{
		resp,
		count,
	}, nil
}
