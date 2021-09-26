package main

import "fmt"

// GO necesita definir de esta forma la funcion
func main() {

	variablesConstantes()
	operadoresAritmeticos()

}

func variablesConstantes() {
	// clase 1 iniciales

	const pi float64 = 3.14
	const pi2 = 3.1416
	fmt.Println("Hola mundo")
	fmt.Println(pi)
	fmt.Println(pi2)

	// declaracion de variables declara y define;
	base := 12

	var altura int = 14

	// go no se ejecuta hasta que no se usen las varibels
	area := base * altura
	fmt.Println(base, altura, area)

	// zeros values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// importante los valores por defectos de GO

	const baseCuadrado = 10

	areaCuadrado := baseCuadrado * baseCuadrado

	// se concatena con el ,
	fmt.Println("Area cuadrado es iguala a:", areaCuadrado)
}

func operadoresAritmeticos() {
	// clase 2 operaciones aritmeticas

	x := 10
	y := 50

	fmt.Println("La sumae es ", x+y)

	fmt.Println("La restas es ", x-y)

	fmt.Print("la multiplicacion es ", x*y)

}
