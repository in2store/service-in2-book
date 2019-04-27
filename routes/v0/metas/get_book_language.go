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

func (req GetBookLanguage) Output(ctx context.Context) (interface{}, error) {
	enums := types.BookLanguage(0).Enums()
	data := make([]MetaItem, 0)
	for _, e := range enums {
		data = append(data, MetaItem{
			Value: e[0],
			Label: e[1],
		})
	}
	return data, nil
}
