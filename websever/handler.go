package websever

import (
	"MySqlSever/Tools"
	"MySqlSever/apiHandler"
	"MySqlSever/apiHandler/reportdata"
	"MySqlSever/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//建立数据库连接
func ConnectDB(c *gin.Context) {

	data, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("ConnectDB req: %v", string(data))

	//json转化为struct
	dbDataStruct := model.DBConnect{}
	err := json.Unmarshal([]byte(data), &dbDataStruct)
	if err != nil {
		log.Printf("DbConnect Data UnMarshal fail : %v", err)
	} else {
		//连接数据库
		apiHandler.DBConnect(dbDataStruct.UserName, dbDataStruct.PassWord, dbDataStruct.Addr, dbDataStruct.Port, dbDataStruct.DbName)
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Database connection Success",
	})

}

//根据基准值(时间)插入数据
func InsertByStandard(c *gin.Context) {

	data, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("ConnectDB req: %v", string(data))

	//json转化为map
	dataMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &dataMap)
	if err != nil {
		log.Printf("Insert Data UnMarshal fail : %v", err)
	} else {

		tableName := model.GetTableName(dataMap["TableName"].(string))
		//判断是哪一类表
		if strings.HasPrefix(tableName, "report_data") {
			data := dataMap["Data"]
			if v, ok := data.([]interface{})[0].(map[string]interface{}); ok {

				fmt.Printf("v type : %v", v)
				//插入数据
				num := int(dataMap["Num"].(float64)) //大坑：json 解码后，go默认是float64的，若确定为int类型，必须先decode为float64再强转为int，否则报错"panic: interface conversion: interface {} is float64, not int"
				reportdata.ReportDataCreate(v, dataMap["Standard"].(string), num)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Insert Success",
	})

}

//根据时间范围查询数据和下载数据
func FindByTimeAndDownload(c *gin.Context) {

	data, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("ConnectDB req: %v", string(data))

	//json转化为struct
	findStruct := model.FindAndDownload{}
	err := json.Unmarshal([]byte(data), &findStruct)
	if err != nil {
		log.Printf("FindByTime Data UnMarshal fail : %v", err)
	} else {
		tableName := model.GetTableName(findStruct.TableName)
		if strings.HasPrefix(tableName, "report_data") {
			rtableHead := []string{"id", "created_at", "task_type", "vod_task_id", "app_id", "region", "file_id", "biz_id", "task_id", "task_host", "task_status", "task_error_code", "req_time", "schedule_time", "real_start_time", "finish_time", "file_size", "file_num", "duration", "req_para", "task_result", "report_ip", "atomic_task_id", "user_defined1", "user_defined2", "user_defined3", "user_defined7", "user_defined8", "user_defined4", "user_defined5", "user_defined6", "output_file_size", "non_bill", "non_vod", "mode"}
			rData := reportdata.ReportDataGet(findStruct.Appid, findStruct.StartTime, findStruct.EndTime)
			//tableHead = append(tableHead, rtableHead...)
			//下载到excel
			Tools.CreateExcel(rtableHead, rData, findStruct.DownloadPath)

		}

	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "FindAndDownload Success",
	})

}

//根据指定字段查询并求和
func FindAndSum(c *gin.Context) {

	data, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("ConnectDB req: %v", string(data))
	var sum int64
	var count int

	//json转化为map
	dataMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &dataMap)
	if err != nil {
		log.Printf("FindAndSum Data UnMarshal fail : %v", err)
	} else {

		tableName := model.GetTableName(dataMap["TableName"].(string))
		//判断是哪一类表
		if strings.HasPrefix(tableName, "report_data") {
			data := dataMap["Condition"]
			if v, ok := data.([]interface{})[0].(map[string]interface{}); ok {
				sum, count = reportdata.ReportDataSum(tableName, v, dataMap["GoalSum"].(string), dataMap["StartTime"].(string), dataMap["EndTime"].(string), dataMap["DownloadPath"].(string))
				log.Printf("The %s Sum is: %d", dataMap["GoalSum"].(string), sum)
				log.Printf("The %s count is: %d", dataMap["GoalSum"].(string), count)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "FindAndSum Success",
		"Sum": strconv.FormatInt(sum,10),
		"Count": strconv.Itoa(count),
	})

}
