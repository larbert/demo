package main

import (
    "log"
    "fmt"

    "github.com/xuri/excelize/v2"
)

func main() {
	readWeiChatBill()
	readAlipayBill()
}

func readWeiChatBill() {
	f, err := excelize.OpenFile("excel1.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "A18")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}
	rows = rows[17:]
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func readAlipayBill() {

}