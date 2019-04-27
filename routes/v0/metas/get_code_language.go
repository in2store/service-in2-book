package metas

import (
	"context"
	"github.com/in2store/service-in2-book/constants/types"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
)

func init() {
	Router.Register(courier.NewRouter(GetCodeLanguage{}))
}

// 获取代码语言配置
type GetCodeLanguage struct {
	httpx.MethodGet
}

func (req GetCodeLanguage) Path() string {
	return "/code-language"
}

func (req GetCodeLanguage) Output(ctx context.Context) (result interface{}, err error) {
	return types.CodeLanguage(0).Enums(), nil
}
