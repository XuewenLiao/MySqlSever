package reportdata

import (
	"MySqlSever/Tools"
	"MySqlSever/model"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
	"time"
)

//插入数据的实际业务
/**
参数：reportMap：要插入的数据；
     standard：以standard为基准增加数据（比如选择CreatedAt作为基准，则表示按时间插入）；
	 num：插入多少条。（比如将CreatedAt作为基准，CreatedAt为"2019-07-01"，num为10，则实际插入2019-07-01到2019-07-10共10天数据）
*/
func ReportDataCreate(reportMap map[string]interface{}, standard string, num int) {
	//插入数据
	//r := model.ReportDataModel{
	//	AppId:        1451008645,
	//	TaskType:     "atomic-trans",
	//	Duration:     10,
	//	CreatedAt:    "2019-06-03 06:00:00",
	//	Mode:         1,
	//	TaskStatus:   3,
	//	UserDefined5: "h264",
	//	UserDefined7: 100,
	//	UserDefined8: 4072,
	//}

	var reportSlice []interface{}
	//取出CreatedAt对应的时间
	standardVal := reportMap[standard].(string)
	//转为时间戳
	formatTime, error := time.Parse("2006-01-02 15:04:05", standardVal)
	var timeTamp int64
	if error == nil {
		timeTamp = formatTime.Unix()
		fmt.Println("CreatedAt转为时间戳：", timeTamp)
	}

	//对时间戳做加法，加数为：86400*num。并生成不同的reportMap。
	reportMapFirstDay := make(map[string]interface{})
	reportMapFirstDay = Tools.CopyMap(reportMap)
	reportSlice = append(reportSlice, reportMapFirstDay)

	for i := 1; i < num; i++ {
		timeTampNew := timeTamp + int64(86400*i)
		//时间戳转为string
		timeNew := time.Unix(timeTampNew, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("第%d天时间为%s\n", i, timeNew)

		//生成不同的reportMap
		reportMapNew := make(map[string]interface{})
		reportMap[standard] = timeNew
		reportMapNew = Tools.CopyMap(reportMap)
		fmt.Printf("第%d天reportMap为: %v\n", i+1, reportMapNew)

		reportSlice = append(reportSlice, reportMapNew)
		fmt.Printf("第%d天reportSlice为: %v\n", i+1, reportSlice)

	}

	//将reportMap转换为结构体
	for i := 0; i < num; i++ {
		r := model.ReportDataModel{}
		err := mapstructure.Decode(reportSlice[i], &r)
		if err != nil {
			log.Fatal("ReportDataModel Fail：%s", err)
		} else {
			log.Printf("ReportDataModel Success")
			log.Printf("get the ReportDataModel: %v", r)
		}

		//rNew := model.ReportDataModel{}
		//rNew = r
		//fmt.Printf("get the rNew: %v",rNew)

		//插入数据
		if err := r.Create(); err != nil {
			log.Fatal("Insert Fail !")
		} else {
			log.Printf("Insert Success !")
		}

	}

	////将reportMap转换为结构体
	//r := model.ReportDataModel{}
	//err := mapstructure.Decode(reportMap,&r)
	//if err != nil {
	//	log.Fatal("ReportDataModel Fail：%s",err)
	//}else {
	//	log.Printf("ReportDataModel Success")
	//	log.Printf("get the ReportDataModel: %v",r)
	//}
	//
	////插入数据
	//if err := r.Create(); err != nil {
	//	log.Fatal("Insert Fail !")
	//} else {
	//	log.Printf("Insert Success !")
	//}
}
