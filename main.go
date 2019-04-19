package main

import (
	"github.com/johnnyeven/libtools/servicex"

	"github.com/in2store/service-in2-book/global"
	"github.com/in2store/service-in2-book/routes"
)

func main() {
	servicex.Execute()
	global.Config.Server.Serve(routes.RootRouter)
}
