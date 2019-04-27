package metas

import "github.com/johnnyeven/libtools/courier"

var Router = courier.NewRouter(MetasGroup{})

type MetasGroup struct {
	courier.EmptyOperator
}

func (MetasGroup) Path() string {
	return "/metas"
}
