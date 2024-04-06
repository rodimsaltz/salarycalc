package main

import "fmt"

func calcularSalario(horasTrabalhadas int, valorPorHora float64) float64 {
	return float64(horasTrabalhadas) * float64(valorPorHora)
}

func main() {
	var idFuncionario, horasTrabalhadas int
	var valorPorHora float64

	fmt.Printf("Digite o seu ID: ")
	fmt.Scan(&idFuncionario)

	fmt.Printf("Digite quantas horas você trabalhou: ")
	fmt.Scan(&horasTrabalhadas)

	fmt.Printf("Digite o valor por horas trabalhadas: ")
	fmt.Scan(&valorPorHora)

	salario := calcularSalario(horasTrabalhadas, valorPorHora)

	fmt.Printf("ID do funcionário: %d\n", idFuncionario)
	fmt.Printf("O valor do salário: %.2f\n", salario)

}
