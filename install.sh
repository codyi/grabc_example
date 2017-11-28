#!/bin/sh
echo "GoCronJob 一键安装向导"

echo "安装 beego"
go get -u github.com/astaxie/beego

echo "安装 bee"
go get -u github.com/beego/bee

echo "安装 mysql go-sql-driver"
go get -u github.com/go-sql-driver/mysql

echo "安装 beego/orm"
go get github.com/astaxie/beego/orm

echo 
cp conf/app.conf.example conf/app.conf

read -p "输入监听端口号(默认8080):" listen_port

if [ -z $listen_port ];then
	listen_port="8080"
fi
echo "\nhttpport = "$listen_port >> conf/app.conf
echo 
read -p "确认创建数据库并导入初始化数据(yes or no):" is_install

if [ -z $is_install ];then
	is_install="no"
fi

if [ $is_install = "yes" ]; then
	echo 
	echo "请确认输入的mysql用户是否拥有创建数据库、创建数据表的权限"

	read -p "输入mysql用户名:" user_name
	read -p "输入mysql密码:" passpord
	read -p "输入数据库名称(默认：goCronJob):" database_name
	read -p "输入mysql host(默认：localhost):" host
	read -p "输入mysql port(默认：3306):" host_port

	if [ -z $host_port ];then
		host_port="3306"
	fi

	if [ -z $host ]; then
		host="localhost"
	fi

	if [ -z $database_name ]; then
		database_name="goCronJob"
	fi

	mysql -u $user_name -p$passpord -h $host -P $host_port -e "create database if not exists "$database_name" default charset utf8 collate utf8_general_ci;use "$database_name";source conf/install.sql;"
	echo "\n# 数据库配置\ndb.host = "$host"\ndb.user = "$user_name"\ndb.password = "$passpord"\ndb.port = "$host_port"\ndb.database_name = "$database_name"\n" >> conf/app.conf
else 
	echo "数据库创建取消，请按照如下提示手动安装并在数据库中导入conf目录下的install.sql"
	echo "在conf目录app.conf文件中增加如下配置"
	echo
	printf "# 数据库配置\ndb.host = mysq_host\ndb.user = mysql_user_name\ndb.password = password\ndb.port = port\ndb.name = database_name\n"
fi

echo "\nisInstall = true" >> conf/app.conf
echo 
echo "配置完成"
echo 