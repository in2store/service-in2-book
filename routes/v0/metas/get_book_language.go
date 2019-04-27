package metas

import (
	"context"
	"github.com/in2store/service-in2-book/constants/types"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
)

func init() {
	Router.Register(courier.NewRouter(GetBookLanguage{}))
}

// 获取书籍语言配置
type GetBookLanguage struct {
	httpx.MethodGet
}

func (req GetBookLanguage) Path() string {
	return "/book-language"
}

func (req GetBookLanguage) Output(ctx context.Context) (result interface{}, err error) {
	return types.BookLanguage(0).Enums(), nil
}
