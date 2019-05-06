package tags

import (
	"context"
	"github.com/in2store/service-in2-book/global"
	"github.com/in2store/service-in2-book/modules"
	libModule "github.com/johnnyeven/eden-library/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(CreateTag{}))
}

type CreateTagBody struct {
	// 标签名称
	Name string `json:"name"`
}

// 创建标签
type CreateTag struct {
	httpx.MethodPost
	Body CreateTagBody `name:"body" in:"body"`
}

func (req CreateTag) Path() string {
	return ""
}

func (req CreateTag) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	id, err := libModule.NewUniqueID(global.Config.ClientID)
	if err != nil {
		logrus.Errorf("[CreateTag] libModule.NewUniqueID err: %v", err)
		return nil, err
	}
	result, err = modules.CreateTag(id, req.Body.Name, db)
	if err != nil {
		logrus.Errorf("[CreateTag] modules.CreateTag err: %v, request: %+v", err, req.Body)
	}
	return
}
