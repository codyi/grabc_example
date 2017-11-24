# GoCronJob
GoCronJob 是用于管理本机的定时任务系统。系统可以监听job的执行状态，记录执行的日志。

###### 详情 [gocron.liguosong.com](http://gocron.liguosong.com)

## 快速安装

#### 下载并且安装

    go get github.com/astaxie/beego

#### Create file `hello.go`
```go
package main

import "github.com/astaxie/beego"

func main(){
    beego.Run()
}
```
#### Build and run

    go build hello.go
    ./hello

#### Go to [http://localhost:8080](http://localhost:8080)

Congratulations! You've just built your first **beego** app.

###### Please see [Documentation](http://beego.me/docs) for more.

## Features

* RESTful support
* MVC architecture
* Modularity
* Auto API documents
* Annotation router
* Namespace
* Powerful development tools
* Full stack for Web & API