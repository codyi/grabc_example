## GRABC 示例
[![GitHub forks](https://img.shields.io/github/forks/codyi/grabc_example.svg?style=social&label=Forks)](https://github.com/codyi/grabc_example/network)
[![GitHub stars](https://img.shields.io/github/stars/codyi/grabc_example.svg?style=social&label=Starss)](https://github.com/codyi/grabc_example/stargazers)
[![GitHub last commit](https://img.shields.io/github/last-commit/codyi/grabc_example.svg)](https://github.com/codyi/grabc_example)
[![Go Report Card](https://goreportcard.com/badge/github.com/codyi/grabc_example)](https://goreportcard.com/report/github.com/codyi/grabc_example)  

GRABC 是一个beego权限管理插件，这个是一个beego应用GRABC的示例~~

### 安装
进入go的src目录，执行以下命令：

    git clone git@github.com:codyi/grabc_example.git
    
然后进入grabc_example目录执行install.sh，根据提示输入数据库的配置。
然后，然后安装完啦~~~~~~是不是很简单抓紧运行下面命令看一下效果吧

<pre>
bee run
</pre>
或者
<pre>
go build
./grabc_example
</pre>

初始化用户  

用户名：18888888888 密码：123456  超级管理员--可以配置路由、权限、角色、用户修改密码  

用户名：17777777777 密码：123456  基础用户--只能登陆进来，不能可以修改密码  


tips:可以根据权限动态控制系统内部的按钮显示，包括导航，链接~~

#### 注意：不能直接运行go run main.go，因为GRABC中在启动的时候会设置默认的视图路径地址，直接运行go run main.go不能获取到正确的目录地址
