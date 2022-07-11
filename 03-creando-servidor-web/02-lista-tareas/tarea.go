package main

import "fmt"

func main() {
	task1 := &task{
		nombre:      "Completar el curso de GO",
		descripcion: "Completar mi curso de GO de Platzi en esta semana",
	}
	fmt.Println(task1)
	fmt.Printf("nombre: %s, descripcion: %s, completada: %v\n", task1.nombre, task1.descripcion, task1.completado)
	fmt.Printf("%+v\n", task1)

	task1.SetNombre("Nuevo nombre")
	task1.SetDescripcion("Una nueva descripcion")
	task1.MarcarCompleta()

	fmt.Printf("%+v\n", task1)

	fmt.Println("----------------------------------------------")

	task2 := &task{
		nombre:      "Completar el curso de Typescript",
		descripcion: "Completar mi curso de Typescript de Platzi en esta semana",
	}

	task3 := &task{
		nombre:      "Completar el curso de Python",
		descripcion: "Completar mi curso de Python de Platzi en esta semana",
	}

	lista := &TaskList{
		tasks: []*task{
			task1,
			task2,
		},
	}

	lista.addTask(task3)
	fmt.Printf("%v\n", lista.tasks[0])
	fmt.Printf("%v\n", lista.tasks[1])
	fmt.Printf("%v\n", len(lista.tasks))
	lista.deleteTask(0)
	fmt.Println(len(lista.tasks))
	lista.iterarLista()
	lista.iterarListaRange()

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	mapaTareas := make(map[string]*TaskList)
	mapaTareas["Hector"] = lista

	task4 := &task{
		nombre:      "Completar el curso de Typescript",
		descripcion: "Completar mi curso de Typescript de Platzi en esta semana",
	}

	task5 := &task{
		nombre:      "Completar el curso de Python",
		descripcion: "Completar mi curso de Python de Platzi en esta semana",
	}

	lista2 := &TaskList{
		tasks: []*task{
			task4, task5,
		},
	}
	mapaTareas["xd"] = lista2

	fmt.Println("Tareas de Hector")
	mapaTareas["Hector"].iterarListaRange()

	fmt.Println("Tareas de XD")
	mapaTareas["xd"].iterarListaRange()

}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) SetNombre(nombre string) {
	t.nombre = nombre
}

func (t *task) SetDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) MarcarCompleta() {
	t.completado = true
}

type TaskList struct {
	tasks []*task
}

func (t *TaskList) addTask(task *task) {
	t.tasks = append(t.tasks, task)
}

func (t *TaskList) deleteTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *TaskList) iterarLista() {
	for i := 0; i < len(t.tasks); i++ {
		fmt.Println("Index:", i, "nombre:", t.tasks[i].nombre)
	}
}

func (t *TaskList) iterarListaRange() {
	for index, tarea := range t.tasks {
		fmt.Println("Index:", index, "nombre:", tarea.nombre)
	}
}
