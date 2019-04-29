package categories

import "github.com/johnnyeven/libtools/courier"

var Router = courier.NewRouter(CategoriesGroup{})

type CategoriesGroup struct {
	courier.EmptyOperator
}

func (CategoriesGroup) Path() string {
	return "/categories"
}
