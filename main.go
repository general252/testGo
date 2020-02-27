package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/general252/testGo/pk1/pk2"
	"github.com/general252/testGoMod/mypkt"
)

func main() {
	pk2.MyAdd()
	mypkt.MySum()

	logs.GetBeeLogger().Info("hello")
}
