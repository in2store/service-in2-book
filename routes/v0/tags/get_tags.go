package tags

import (
	"context"
	"github.com/in2store/service-in2-book/global"
	"github.com/in2store/service-in2-book/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/enumeration"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetTags{}))
}

// 获取标签列表
type GetTags struct {
	httpx.MethodGet
	// 是否过滤零热度项
	FilterZeroHeat enumeration.Bool `name:"filterZeroHeat" in:"query"`
	// 是否依照热度排序
	OrderByHeat enumeration.Bool `name:"orderByHeat" in:"query"`
}

func (req GetTags) Path() string {
	return ""
}

func (req GetTags) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.SlaveDB.Get()
	result, err = modules.GetTagsOrderByHeat(req.FilterZeroHeat.True(), req.OrderByHeat.True(), db)
	if err != nil {
		logrus.Errorf("[GetTags] modules.GetTagsOrderByHeat err: %v, request: %+v", err, req)
	}
	return
}
