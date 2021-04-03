package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("Exercicio.xlsx")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rows, err := f.GetRows("DePara")
	fmt.Println(rows)
}
