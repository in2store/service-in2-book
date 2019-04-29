package categories

import (
	"context"
	"github.com/in2store/service-in2-book/global"
	"github.com/in2store/service-in2-book/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(UpdateCategory{}))
}

// 更新分类
type UpdateCategory struct {
	httpx.MethodPatch
	// 分类标识
	CategoryKey string                     `name:"categoryKey" in:"path"`
	Body        modules.UpdateCategoryBody `name:"body" in:"body"`
}

func (req UpdateCategory) Path() string {
	return "/:categoryKey"
}

func (req UpdateCategory) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	err = modules.UpdateCategory(req.CategoryKey, req.Body, db, false)
	if err != nil {
		logrus.Errorf("[UpdateCategory] modules.UpdateCategory err: %v, request: %+v", err, req)
	}
	return
}
