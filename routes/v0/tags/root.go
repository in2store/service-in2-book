package tags

import "github.com/johnnyeven/libtools/courier"

var Router = courier.NewRouter(TagsGroup{})

type TagsGroup struct {
	courier.EmptyOperator
}

func (TagsGroup) Path() string {
	return "/tags"
}
