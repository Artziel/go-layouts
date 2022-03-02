package main

import (
	"fmt"

	Layouts "gitlab.com/artziel/golang-layouts"
)

type MySampleRow struct {
	Layouts.Row
	ID       int    `excelLayout:"column:A,required,min:1"`
	Username string `excelLayout:"column:B,required,min:6"`
	Password string `excelLayout:"column:C,required,min:8"`
	Avatar   string `excelLayout:"column:D,url"`
	Fullname string `excelLayout:"column:E,required"`
	Age      int    `excelLayout:"column:F,required,min:18,max:50"`
}

func main() {

	l := Layouts.ExcelLayout{}

	err := l.ReadFile(MySampleRow{}, "./sample.xlsx")
	if err != nil {
		for _, e := range l.GetErrors() {
			fmt.Printf("Row %d) %s\n", e.RowIndex, Layouts.ErrToMessage(&e))
		}
	} else {
		rows := l.GetRows().([]MySampleRow)

		for i, r := range rows {
			fmt.Printf("%d) ID:%v, Username: %v\n", i, r.ID, r.Username)
		}
	}
}
