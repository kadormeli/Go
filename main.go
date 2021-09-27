package main

import "fmt"

// GO necesita definir de esta forma la funcion
func main() {

	variablesConstantes()
	operadoresAritmeticos()
	fmtFuncion()

	// call function
	simpleValue := returnValue(3)

	value1, value2 := returnDouble(2)

	// si solo necesitas uno le pones el ( _ ) donde no lo necesitas

	fmt.Println("The value is ", simpleValue)

	fmt.Println("The double value is", value1, value2)

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

	fmt.Print("La division es :", x/y)

	fmt.Println("El modulo de los dos numeros es", x%y)

	x++
	fmt.Println("El incremento eso ", x)

	x--
	fmt.Println("El decremento eso ", x)

	funcinWithArguments("Hola mundo", "Kevin", 28, 1, 1.87)

}

func fmtFuncion() {
	helloMessage := "Hello"
	worldMessage := "World"
	fmt.Println(helloMessage, worldMessage)

	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %d cursos\n", nombre, cursos)

	// sprintf lo que reuslta lo guarda en message
	message := fmt.Sprintf("%s tienes mas de %d cursos\n", nombre, cursos)
	fmt.Println(message)

	// tipos de datos
	fmt.Printf("Hello : %T\n ", helloMessage)

	fmt.Printf("Cursos : %T\n", cursos)

}

func funcinWithArguments(msg string, name string, age, pen int, salary float64) {

	fmt.Printf("El mensaje que te mando es %s , me llamo %s, tengo %d años y mido %f y tengo un %d", msg, name, age, salary, pen)

}

func returnValue(arg int) int {
	return arg
}

func returnDouble(arg int) (c, d int) {
	return arg, 2 * arg
}
