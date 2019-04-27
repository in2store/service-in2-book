package v0

import (
	"github.com/in2store/service-in2-book/routes/v0/books"
	"github.com/in2store/service-in2-book/routes/v0/metas"
	"github.com/johnnyeven/libtools/courier"
)

var Router = courier.NewRouter(V0Group{})

func init() {
	Router.Register(books.Router)
	Router.Register(metas.Router)
}

type V0Group struct {
	courier.EmptyOperator
}

func (V0Group) Path() string {
	return "/v0"
}
