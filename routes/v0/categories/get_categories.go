package categories

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
	Router.Register(courier.NewRouter(GetCategories{}))
}

// 获取分类列表
type GetCategories struct {
	httpx.MethodGet
	httplib.Pager
}

func (req GetCategories) Path() string {
	return ""
}

type GetCategoriesResult struct {
	Data  database.CategoryList `json:"data"`
	Total int32                 `json:"total"`
}

func (req GetCategories) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.SlaveDB.Get()
	resp, count, err := modules.GetCategoriesSortAsc(req.Size, req.Offset, db)
	if err != nil {
		logrus.Errorf("[GetCategories] modules.GetCategoriesSortAsc err: %v", err)
		return nil, err
	}
	return GetCategoriesResult{
		resp,
		count,
	}, nil
}
