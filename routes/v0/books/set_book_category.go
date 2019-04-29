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
	Router.Register(courier.NewRouter(SetBookCategory{}))
}

type SetBookCategoryBody struct {
	// 分类标识
	CategoryKey string `json:"categoryKey"`
}

// 设置文档分类
type SetBookCategory struct {
	httpx.MethodPost
	// 文档ID
	BookID uint64              `name:"bookID,string" in:"path"`
	Body   SetBookCategoryBody `name:"body" in:"body"`
}

func (req SetBookCategory) Path() string {
	return "/:bookID/category"
}

func (req SetBookCategory) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	err = modules.SetBookCategory(req.BookID, req.Body.CategoryKey, db)
	if err != nil {
		logrus.Errorf("[SetBookCategory] modules.SetBookCategory err: %v, request: %+v", err, req)
	}
	return
}
