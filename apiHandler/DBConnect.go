package apiHandler

import "MySqlSever/model"

//功能：建立数据库连接：输入——用户名、数据库密码、数据库ip地址、端口号、数据库名
func DBConnect(userName string,passWord string,addr string,port int,dbName string) {

	//model.DB.Init("root", "_gTQ699qwX", "100.121.190.3", 15944, "ReportData")
	//本地mysql:	model.DB.Init("root", "qiangwen950424", "127.0.0.1", 3306, "ReportData")
	model.DB.Init(userName, passWord, addr, port, dbName)

	//defer model.DB.Close()


}
