package categories

import (
	"context"
	"github.com/in2store/service-in2-book/constants/errors"
	"github.com/in2store/service-in2-book/global"
	"github.com/in2store/service-in2-book/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(CreateCategory{}))
}

// 创建分类
type CreateCategory struct {
	httpx.MethodPost
	Body modules.CreateCategoryBody `name:"body" in:"body"`
}

func (req CreateCategory) Path() string {
	return ""
}

func (req CreateCategory) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	result, err = modules.CreateCategory(req.Body, db)
	if err != nil {
		if err != errors.CategoryKeyConflict {
			logrus.Errorf("[CreateCategory] modules.CreateCategory err: %v, request: %+v", err, req.Body)
		}
		return nil, err
	}
	return
}
