package main

import (
	"github.com/astaxie/beego/logs"
	"mydemo/pk1/pk2"
)

func main() {
	pk2.MyAdd()

	logs.GetBeeLogger().Info("hello")
}
