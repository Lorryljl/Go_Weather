package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	file, err := xlsx.OpenFile("../docs/AMap_adcode_citycode.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	sheet := file.Sheets[0] //处理第一个工作表

	//a列和c列的索引分别为0和2
	columnA := 0
	columnB := 1

	columnBData := []string{}

	for index, row := range sheet.Rows {
		if index == 1 {
			continue
		}
		if len(row.Cells) > 0 {
			cell := row.Cells[1]
			fmt.Println(cell)
			columnBData = append(columnBData, cell.String())

		}
	}

	for _, row := range sheet.Rows {
		cellA := row.Cells[columnA]
		cellB := row.Cells[columnB]

		if cellA.Value != "" {
			fmt.Println(cellA.Value, cellB.Value)
		}
	}

	for index, row := range sheet.Rows {
		if index == 1 {
			continue
		}
		if len(row.Cells) > 0 {
			fmt.Println(row.Cells[1])
		}
	}

}
