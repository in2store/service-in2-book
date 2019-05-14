package tags

import (
	"context"
	"github.com/in2store/service-in2-book/constants/errors"
	"github.com/in2store/service-in2-book/database"
	"github.com/in2store/service-in2-book/global"
	"github.com/in2store/service-in2-book/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/httplib"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetBooksByTag{}))
}

// 通过tag获取书籍列表
type GetBooksByTag struct {
	httpx.MethodGet
	// 标签ID（优先级高）
	TagID uint64 `name:"tagID,string" in:"path" default:"0"`
	// 标签名称
	Name string `name:"name" in:"query" default:""`
	httplib.Pager
}

type GetBooksByTagResult struct {
	Data  database.BookMetaList `json:"data"`
	Total int32                 `json:"total"`
}

func (req GetBooksByTag) Path() string {
	return "/:tagID/books"
}

func (req GetBooksByTag) Output(ctx context.Context) (result interface{}, err error) {
	err = req.Validate()
	if err != nil {
		return nil, err
	}
	db := global.Config.SlaveDB.Get()
	if req.TagID == 0 {
		tag, err := modules.GetTagByName(req.Name, db)
		if err != nil {
			if err == errors.TagNotFound {
				return nil, err
			}
			logrus.Errorf("[GetBooksByTag] modules.GetTagByName err: %v, request: %+v", err, req)
			return nil, err
		}
		req.TagID = tag.TagID
	}

	books, count, err := modules.GetBooksByTagID(req.TagID, req.Size, req.Offset, db)
	if err != nil {
		logrus.Errorf("[GetBooksByTag] modules.GetBooksByTagID err: %v, request: %+v", err, req)
		return
	}

	return GetBooksByTagResult{
		Data:  books,
		Total: count,
	}, nil
}

func (req GetBooksByTag) Validate() error {
	if req.TagID == 0 && req.Name == "" {
		return errors.BadRequest.StatusError().WithDesc("请求参数错误，标签ID与标签名称不能同时为空").WithErrTalk()
	}
	return nil
}
