package main

import "excelFunc"

func main() {
	//创建Excel文件
	// excelFunc.CreateExcelFile()

	//读取指定Excel文件
	excelFunc.ReaderExcelFile()

	//创建图表
	excelFunc.CreateEchart()

}
