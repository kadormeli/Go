package main

import "fmt"

// definimos los structs

type task struct {
	nombre      string
	description string
	completado  bool
}

// funciones auxiliare o como los metodos de los struct
// con el asterisco haces que apunte al sitio en memoria
func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizaDescription(description string) {
	t.description = description
}

func (t *task) actulizarNombre(nombre string) {
	t.nombre = nombre
}

type TaskList struct {
	// hay que agregar el apuntador

	// es similar a Java una lista dinamica de task solo que necesita el apuntador
	// y seria como una lista de tareas
	tasks []*task
}

func (T *TaskList) addToList(t *task) {
	T.tasks = append(T.tasks, t)
}

func (T *TaskList) eliminarDeLista(index int) {
	T.tasks = append(T.tasks[:index], T.tasks[index+1:]...)
}

func (t *TaskList) imprimirValores() {

	for i := 0; i < len(t.tasks); i++ {
		fmt.Println(t.tasks[i])
	}

}

func (t *TaskList) imprimirWihtRange() {
	for index, tarea := range t.tasks {
		fmt.Println("Index", index, "nombre ", tarea.nombre)
	}
	// tiene break  y continue al igual que python y apuntadores de c , ademas de slice de python
}

func (t *TaskList) imprimirImcompletos() {
	for index, tarea := range t.tasks {
		if !tarea.completado {
			fmt.Println("Index", index, "nombre ", tarea.nombre)
		}

	}
	// tiene break  y continue al igual que python y apuntadores de c , ademas de slice de python
}

func main() {
	// el & crea un apuntador al struct
	// instanciar unas tareas
	t1 := &task{
		nombre:      "Completar curso go:",
		description: "completar go platzi en esta semana uno ",
	}

	t2 := &task{
		nombre:      "Completar curso python",
		description: "completar go platzi en esta semana dos",
	}

	t3 := &task{
		nombre:      "Completar curso nodejs",
		description: "completar go platzi en esta semana tres",
	}

	t4 := &task{
		nombre:      "Completar curso nodejs",
		description: "completar go platzi en esta semana tres",
	}
	// instancia el struct
	listaKevin := &TaskList{
		tasks: []*task{
			t1, t2,
		},
	}

	listaNestor := &TaskList{
		tasks: []*task{
			t4,
		},
	}

	listaKevin.addToList(t3)

	fmt.Println("Listar Todos")
	listaKevin.imprimirWihtRange()
	fmt.Println("Listar imcompletados")
	listaKevin.tasks[0].marcarCompleta()
	listaKevin.imprimirImcompletos()

	mapaTareas := make(map[string]*TaskList)

	mapaTareas["Kevin"] = listaKevin
	mapaTareas["Nestor"] = listaNestor

	mapaTareas["Nestor"].imprimirWihtRange()
	mapaTareas["Kevin"].imprimirImcompletos()

}
