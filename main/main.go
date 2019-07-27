package main

import (
	"MySqlSever/handler"
	"MySqlSever/model"
	"strings"
)

func main() {
	//建立数据库连接
	model.DB.Init("root", "_gTQ699qwX", "100.121.190.3", 15944, "ReportData")
	defer model.DB.Close()

	/**
	前端数据处理
	1、前端传入json，将json转换为map——待实现
	2、拿到存表名的map，传入GetTableName——待实现
	3、拿到存数据的map，传入对应的表处理方法
	*/
	//指定表名
	tableName := model.GetTableName("report_data_async_2019_04")
	//判断要插入哪一类的表
	if strings.HasPrefix(tableName, "report_data") {

		//假设已解析出map，传入到对应的表处理方法
		reportMap := make(map[string]interface{})
		reportMap["AppId"] = 1451008645
		reportMap["TaskType"] = "atomic-trans"
		reportMap["Duration"] = 10
		reportMap["CreatedAt"] = "2019-04-01 06:00:00"
		reportMap["Mode"] = 1
		reportMap["TaskStatus"] = 3
		reportMap["UserDefined5"] = "h264"
		reportMap["UserDefined7"] = 100
		reportMap["UserDefined8"] = 4072

		handler.ReportDataCreate(reportMap, "CreatedAt", 2)
	}

}
