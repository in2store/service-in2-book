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
	Router.Register(courier.NewRouter(SetBookTag{}))
}

type SetBookTagBody struct {
	// 分类标识
	TagID uint64 `json:"tagID,string"`
}

// 设置文档分类
type SetBookTag struct {
	httpx.MethodPost
	// 文档ID
	BookID uint64         `name:"bookID,string" in:"path"`
	Body   SetBookTagBody `name:"body" in:"body"`
}

func (req SetBookTag) Path() string {
	return "/:bookID/tags"
}

func (req SetBookTag) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	err = modules.SetBookTag(req.BookID, req.Body.TagID, db)
	if err != nil {
		logrus.Errorf("[SetBookTag] modules.SetBookTag err: %v, request: %+v", err, req)
	}
	return
}
