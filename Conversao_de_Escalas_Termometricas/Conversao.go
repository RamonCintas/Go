package main

import "fmt"

const ebulicaoK = 373.0

func main() {

	fmt.Println("Um professor de ensino médio solicitou aos seus alunos que desenvolvessem um programa em go, para converter o valor do ponto de ebulição da água de Kelvin para graus Celsius.")
	fmt.Println("Dica: C = K - 273")
	fmt.Println("Desenvolvedor: Ramon Cintas Gomes")

	tempK := ebulicaoK
	tempC := tempK - 273.0
	fmt.Printf("A temperatura de ebulição da água em Kelvin = %g , temperatura de ebulição da água em Celsius = %g.", tempK, tempC)

}
