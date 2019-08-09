package Tools

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"strconv"
)

func CalcuSum(tableHead []string, data []interface{}, downloadPath string, goalSum string) (int64, int) {

	////下载数据为excel
	CreateExcel(tableHead, data, downloadPath)

	//打开表
	excl, err := excelize.OpenFile(downloadPath)
	if err != nil {
		log.Fatal("Open file fail !: ", err)
	} else {
		log.Printf("Open file succcess ! ")
	}

	//获取待计算的字段名坐标
	crAxis, _ := excl.SearchSheet("Sheet1", goalSum)
	log.Printf("Goal calculate axis : ", crAxis)

	//单元格坐标转索引
	col, row, _ := excelize.CellNameToCoordinates(crAxis[0])
	log.Printf("ColName col: %d--row: %d", col, row)

	//计算和
	var sum int64
	for i := 0; i < 2; i++ {
		valueTemp, _ := excelize.CoordinatesToCellName(col, row+1)
		valueString, _ := excl.GetCellValue("Sheet1", valueTemp)
		valueCal, _ := strconv.ParseInt(valueString, 10, 64)
		sum += valueCal
		fmt.Printf("value: ", valueCal)
	}
	fmt.Printf("sum: ", sum)

	//条数
	count := len(data)

	return sum, count
}
