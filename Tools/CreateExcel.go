package Tools

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"reflect"
)

/**
功能：将数据下载到本地生成excel：
输入——表头数据、存数据struct的slice、生成的文件名；
*/
func CreateExcel(tableHead []string, data []interface{}, fileName string) {

	////将结构体转换为reportMap
	//var rMapSlice []map[string]interface{}
	//for i := 0; i < len(data); i++ {
	//	rMap := make(map[string]interface{})
	//	err := mapstructure.Decode(data[i], &rMap)
	//	if err != nil {
	//		log.Fatal("ReportDataModel Fail：%s", err)
	//	} else {
	//		log.Printf("rMap Success")
	//		log.Printf("get the rMap: %v", rMap)
	//	}
	//
	//	rMapSlice = append(rMapSlice,rMap)
	//
	//}
	//log.Printf("rMapslice: ",rMapSlice)

	excel := excelize.NewFile()
	excel.NewSheet("Sheet1")

	//插入表头
	for c := 0; c < len(tableHead); c++ {
		axis, _ := excelize.CoordinatesToCellName(c+1, 1)
		err := excel.SetCellValue("Sheet1", axis, tableHead[c])
		if err != nil {
			log.Fatal("Insert tableHead Fail", err)
		} else {
			log.Printf("Insert tableHead success !")
		}

	}

	//写入数据
	for r := 0; r < len(data); r++ {
		var dataInterface []interface{} = make([]interface{}, 50) //一定要用make声明，否则会panic
		var dataInter interface{} = data[r]                       //把data拆成一个个数组

		//通过反射读interface{}的值
		value := reflect.ValueOf(dataInter)
		for i := 0; i < value.NumField(); i++ {
			dataInterface[i] = value.Field(i)
			fmt.Printf("Field %d: %v\n", i, dataInterface[i])
		}

		for c := 0; c < len(tableHead); c++ {
			axis, _ := excelize.CoordinatesToCellName(c+1, r+2)
			err := excel.SetCellValue("Sheet1", axis, dataInterface[c])
			if err != nil {
				log.Fatal("Write Fail", err)
			} else {
				log.Printf("Write success !")
			}
		}
	}

	//存excel
	err := excel.SaveAs(fileName)
	if err != nil {
		log.Fatal("Save Fail", err)
	} else {
		log.Printf("Save success !")
	}

}
