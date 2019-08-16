package reportdata

import (
	"MySqlSever/Tools"
	"MySqlSever/model"
	"fmt"
	"log"
	"reflect"
	"strings"
)

//根据字段查询并求和
/**
功能：输入：表名、要查询的字段、要求和的字段、开始时间、结束时间、下载路径
*/
func ReportDataSum(tableName string, condition map[string]interface{}, goalSum string, startTime string, endTime string, downloadPath string) (int64, int) {

	//拼接sql
	var re string
	var tempStr []string

	sqlHead := "select * from " + tableName + " where"
	tempStr = append(tempStr, sqlHead)
	for k, v := range condition {

		conditionStr := k + "=" + Tools.Int2String(v) + " and"
		tempStr = append(tempStr, conditionStr)
		re = strings.Join(tempStr, " ")
	}

	sql := re + " created_at>=" + `"` + startTime + `"` + " and " + "created_at<=" + `"` + endTime + `"` + ";"
	log.Printf("sql: %s", sql)

	//根据待查询参数获取结果集
	//sql := "select * from report_data_async_2019_07 where app_id=1451008450 and task_type='atomic-trans' and mode=0 and created_at >= '2019-07-01 01:00:00' and  created_at <= '2019-07-02 02:00:00';"
	//var  result []model.ReportDataModel
	var resultModel []model.ReportDataModel
	model.DB.Self.Raw(sql).Scan(&resultModel)
	fmt.Printf("value: %v", resultModel)
	fmt.Print("value type is : ", reflect.TypeOf(resultModel))

	//转化为[]interface{}
	var rInterfaceSlice []interface{} = make([]interface{}, len(resultModel))
	for index, data := range resultModel {
		rInterfaceSlice[index] = data
		log.Print("rInterfaceSlice type is : ", reflect.TypeOf(data))
	}

	rtableHead := []string{"id", "created_at", "task_type", "vod_task_id", "app_id", "region", "file_id", "biz_id", "task_id", "task_host", "task_status", "task_error_code", "req_time", "schedule_time", "real_start_time", "finish_time", "file_size", "file_num", "duration", "req_para", "task_result", "report_ip", "atomic_task_id", "user_defined1", "user_defined2", "user_defined3", "user_defined7", "user_defined8", "user_defined4", "user_defined5", "user_defined6", "output_file_size", "non_bill", "non_vod", "mode"}

	//计算和以及条数
	sum, count := Tools.CalcuSum(rtableHead, rInterfaceSlice, downloadPath, goalSum)
	return sum, count

}
