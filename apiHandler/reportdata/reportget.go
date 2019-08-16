package reportdata

import (
	"MySqlSever/model"
	"log"
	"reflect"
)

//查询数据的实际业务
/**
参数：输入——appid、起始时间、结束时间；输出——数据集合
*/
func ReportDataGet(appid int64, startTime string, endTime string) []interface{} {

	r, err := model.GetData(appid, startTime, endTime)
	if err != nil {
		log.Fatal("Get Fail：", err)
	} else {
		log.Print("Get Success,Message is : ", r)
		log.Print("Message type is : ", reflect.TypeOf(r))
	}

	//[]type{数组} 转换为 []interface{}
	var interfaceSlice []interface{} = make([]interface{}, len(r))
	for index, data := range r {
		interfaceSlice[index] = data
		log.Print("interfaceSlice type is : ", reflect.TypeOf(data))
	}

	return interfaceSlice

}

