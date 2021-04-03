package main

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Cliente struct {
	Cliente          string     `json:"Cliente"`
	CodigoSistemaXYZ string     `json:"CodigoSistemaXYZ"`
	Contas           typeContas `json:"Contas"`
}

type typeContas []int

func main() {

	var clientes []Cliente

	f, err := excelize.OpenFile("Exercicio.xlsx")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rows, err := f.GetRows("DePara")

	for nRow, row := range rows {

		if nRow > 0 {
			var nameClient = row[0]
			var codeCliente = row[1]
			result := searchClientContas(nameClient)
			clientes = append(clientes, Cliente{Cliente: nameClient, CodigoSistemaXYZ: codeCliente, Contas: result})
		}
	}
	fmt.Println(clientes)
}

func searchClientContas(codClient string) typeContas {

	var contas typeContas

	f, err := excelize.OpenFile("Exercicio.xlsx")
	if err != nil {
		fmt.Println("Error:", err)
	}

	rows, err := f.GetRows("Contas")
	for _, row := range rows {

		if row[0] == codClient {
			int1, err := strconv.Atoi(row[1])
			if err != nil {
				fmt.Println("Error:", err)
			}
			contas = append(contas, int(int1))
		}
	}
	return contas
}
