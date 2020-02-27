
1. go mod init github.com/general252/testGo
2. go build


go build 依赖文件, 会下载到 GOPATH/pkg/mod/*目录下


go mod 相关命令
go mod download 下载 go.mod 文件中指明的所有依赖

go mod tidy 整理现有的依赖，删除未使用的依赖。

go mod graph 查看现有的依赖结构

go mod init 生成 go.mod 文件 (Go 1.13 中唯一一个可以生成 go.mod 文件的子命令)

go mod edit 编辑 go.mod 文件

go mod vendor 导出现有的所有依赖 (事实上 Go modules 正在淡化 Vendor 的概念)

go mod verify 校验一个模块是否被篡改过

go clean -modcache 清理所有已缓存的模块版本数据。

go mod 查看所有 go mod的使用命令。


————————————————
版权声明：本文为CSDN博主「无风的雨」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/guyan0319/article/details/101783164





------------------------------代码
package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/general252/testGo/pk1/pk2"
)

func main() {
	pk2.MyAdd()

	logs.GetBeeLogger().Info("hello")
}


------------------------------
E:.
│  go.mod
│  go.sum
│  main.go
│  mydemo.exe
│  readme.txt
│
└─pk1
    └─pk2
            Hello.go
