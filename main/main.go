package main

import (
	"MySqlSever/handler/reportdata"
	"MySqlSever/model"
	"log"
	"strings"
)

func main() {
	//建立数据库连接
	//model.DB.Init("root", "_gTQ699qwX", "100.121.190.3", 15944, "ReportData")
	//本地mysql
	model.DB.Init("root", "qiangwen950424", "127.0.0.1", 3306, "ReportData")

	defer model.DB.Close()

	///**
	//功能：根据基准值插入数据：输入——表名、要插入的数据、基准字段、插入的条数
	//前端数据处理
	//1、前端传入json，将json转换为map——待实现
	//2、拿到存表名的map，传入GetTableName——待实现
	//3、拿到存数据的map，传入对应的表处理方法
	//*/
	////指定表名
	//tableName := model.GetTableName("report_data_async_2019_06")
	////判断要插入哪一类的表
	//if strings.HasPrefix(tableName, "report_data") {
	//
	//	//假设已解析出map，传入到对应的表处理方法
	//	reportMap := make(map[string]interface{})
	//	reportMap["AppId"] = 1451008645
	//	reportMap["TaskType"] = "atomic-trans"
	//	reportMap["Duration"] = 10
	//	reportMap["CreatedAt"] = "2019-06-05 06:00:00"
	//	reportMap["Mode"] = 1
	//	reportMap["TaskStatus"] = 3
	//	reportMap["UserDefined5"] = "h264"
	//	reportMap["UserDefined7"] = 100
	//	reportMap["UserDefined8"] = 4072
	//	reportMap["ReqTime"] = "00-00-00"
	//	reportMap["ScheduleTime"] = "00-00-00"
	//	reportMap["RealStartTime"] = "00-00-00"
	//	reportMap["FinishTime"] = "00-00-00"
	//
	//	reportdata.ReportDataCreate(reportMap, "CreatedAt", 2)
	//}

	///**
	//功能：根据时间范围查询数据：输入——表名、appid、起始时间、结束时间；输出——数据[]interface{}集合
	//1、前端传入json，内容为：表名、appid、起始时间、结束时间——待实现
	//2、拿到存表名的map，传入GetTableName——待实现
	//3、拿到appid、起始时间、结束时间的map，传入查询的方法
	//*/
	//
	//tableName := model.GetTableName("report_data_async_2019_07")
	//////表头（每个表的字段）
	////var tableHead []string
	////判断要插入哪一类的表
	//if strings.HasPrefix(tableName, "report_data") {
	//	rtableHead := []string{"id", "created_at", "task_type", "vod_task_id", "app_id", "region", "file_id", "biz_id", "task_id", "task_host", "task_status", "task_error_code", "req_time", "schedule_time", "real_start_time", "finish_time", "file_size", "file_num", "duration", "req_para", "task_result", "report_ip", "atomic_task_id", "user_defined1", "user_defined2", "user_defined3", "user_defined7", "user_defined8", "user_defined4", "user_defined5", "user_defined6", "output_file_size", "non_bill", "non_vod", "mode"}
	//	fmt.Printf("tableName: %s", tableName)
	//	rData := reportdata.ReportDataGet(1451008450, "2019-07-01 01:00:00", "2019-07-02 02:00:00")
	//	//tableHead = append(tableHead, rtableHead...)
	//	//下载到excel
	//	Tools.CreateExcel(rtableHead, rData)
	//
	//}

	/**
	功能：根据字段查询并求和: 输入：表名、要查询的字段、要求和的字段、开始时间、结束时间、下载路径
	1、前端传入json，内容为：表名、appid、待查询的字段、查询类型集合、起始时间、结束时间——待实现
	2、拿到待查询的字段的map——模拟
	3、拿到所有参数，传入ReportDataSum
	*/

	//var condition map[string]interface{}
	//condition := make(map[string]interface{})
	//condition["app_id"] = 123

	condition := map[string]interface{}{
		"app_id":    1451008450,
		"task_type": "atomic-trans",
		"mode":      0,
	}

	tableName := model.GetTableName("report_data_async_2019_07")

	goalSum := "duration"

	startTime := "2019-07-01 01:00:00"

	endTime := "2019-07-02 02:00:00"

	path := "reportSum.xlsx"

	////表头（每个表的字段）
	//判断要插入哪一类的表
	if strings.HasPrefix(tableName, "report_data") {
		//rtableHead := []string{"id", "created_at", "task_type", "vod_task_id", "app_id", "region", "file_id", "biz_id", "task_id", "task_host", "task_status", "task_error_code", "req_time", "schedule_time", "real_start_time", "finish_time", "file_size", "file_num", "duration", "req_para", "task_result", "report_ip", "atomic_task_id", "user_defined1", "user_defined2", "user_defined3", "user_defined7", "user_defined8", "user_defined4", "user_defined5", "user_defined6", "output_file_size", "non_bill", "non_vod", "mode"}
		//fmt.Printf("tableName: %s", tableName)
		//rData := reportdata.ReportDataGet(1451008450, "2019-07-01 01:00:00", "2019-07-02 02:00:00")
		//
		////下载到excel
		//Tools.CreateExcel(rtableHead, rData,"reportdata.xlsx")
		sum, count := reportdata.ReportDataSum("report_data_async_2019_07", condition, goalSum, startTime, endTime, path)
		log.Printf("The %s Sum is: %d", goalSum, sum)
		log.Printf("The %s count is: %d", goalSum, count)

	}

}
