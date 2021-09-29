package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	entrada := leerEntrada()
	operador := leerEntrada()
	fmt.Println(entrada, operador)

	// instanciar
	c := calculadora{}
	resutl := c.operate(entrada, operador)

	fmt.Println(resutl)
}

type calculadora struct{}

// received funcion
func (calculadora) operate(entrada string, operador string) int {

	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		fmt.Print(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Print(operador1 - operador2)
		return operador1 - operador2

	case "*":
		fmt.Print(operador1 * operador2)
		return operador1 * operador2

	case "/":
		fmt.Print(operador1 / operador2)
		return operador1 / operador2

	default:
		fmt.Print(operador + "no contemplados")
		return 0
	}

}

func parsear(valores string) int {
	operador, _ := strconv.Atoi(valores)
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
