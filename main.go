package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/general252/testGo/pk1/pk2"
	"github.com/general252/testGoMod/mypkt"
	"time"
)

func main() {
	pk2.MyAdd()
	mypkt.MySum()
	
	for i:=0; i < 10; i ++ {
		fmt.Printf("time %v\n", time.Now().Unix())
		time.Sleep(time.Second)
	}
	logs.GetBeeLogger().Info("hello")
}
