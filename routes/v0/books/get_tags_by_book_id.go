package books

import (
	"context"
	"github.com/in2store/service-in2-book/global"
	"github.com/in2store/service-in2-book/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetTagsByBookID{}))
}

// 通过文档ID获取标签
type GetTagsByBookID struct {
	httpx.MethodGet
	// 文档ID
	BookID uint64 `name:"bookID,string" in:"path"`
}

func (req GetTagsByBookID) Path() string {
	return "/:bookID/tags"
}

func (req GetTagsByBookID) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.SlaveDB.Get()
	result, err = modules.GetTagsByBookID(req.BookID, db)
	if err != nil {
		logrus.Errorf("[GetTagsByBookID] modules.GetTagsByBookID err: %v, request: %+d", err, req.BookID)
	}
	return
}
