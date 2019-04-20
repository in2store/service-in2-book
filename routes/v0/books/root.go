package books

import "github.com/johnnyeven/libtools/courier"

var Router = courier.NewRouter(BooksGroup{})

type BooksGroup struct {
	courier.EmptyOperator
}

func (BooksGroup) Path() string {
	return "/books"
}
