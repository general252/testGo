package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/general252/testGo/pk1/pk2"
)

func main() {
	pk2.MyAdd()

	logs.GetBeeLogger().Info("hello")
}
