package main

import (
	"fmt"
)

func main() {
	//Comentario
	usu1 := Usuarios{
		ID:     1,
		Nombre: "Mario",
		Grito:  "41546859"}

	usu2 := Usuarios{
		ID:     2,
		Nombre: "Jose",
		Grito:  "41546859"}

	usu3 := Usuarios{
		ID:     3,
		Nombre: "Katt",
		Grito:  "41546859"}

	usu4 := Usuarios{
		ID:     4,
		Nombre: "Coca",
		Grito:  "41546859"}

	usu5 := Usuarios{
		ID:     5,
		Nombre: "Ken",
		Grito:  "41546859"}

	arbolito := New_Arbol()
	Insertar_Arbol(arbolito, usu1)
	Insertar_Arbol(arbolito, usu2)
	Insertar_Arbol(arbolito, usu3)
	Insertar_Arbol(arbolito, usu4)
	Insertar_Arbol(arbolito, usu5)

	fmt.Println("jalo")
	RecorridoPre(arbolito.Raiz)
	RecorridoPost(arbolito.Raiz)
	RecorrerInord(arbolito.Raiz)
}

type Usuarios struct {
	ID     int
	Nombre string
	Grito  string
}

type Nodo struct {
	User      Usuarios
	Izquierda *Nodo
	Derecha   *Nodo
}

func New_Nodo(user Usuarios) *Nodo {
	return &Nodo{
		user,
		nil,
		nil}
}

type Arbol struct {
	Raiz *Nodo
}

func New_Arbol() *Arbol {
	return &Arbol{nil}
}

func Insertar_Arbol(arbol *Arbol, user Usuarios) {
	Insertar_Nodo(&arbol.Raiz, user)
}

func Insertar_Nodo(nodo **Nodo, user Usuarios) bool {

	if *nodo == nil {
		*nodo = New_Nodo(user)
		return true
	} else if user.ID < (*nodo).User.ID {
		return Insertar_Nodo(&(*nodo).Izquierda, user)
	} else if user.ID > (*nodo).User.ID {
		return Insertar_Nodo(&(*nodo).Derecha, user)
	}
	return false
}

func RecorridoPre(nodo *Nodo) {
	fmt.Println("pre")
	if nodo == nil {
		fmt.Println("No existe arbol")
	}
	if nodo.Izquierda != nil {
		fmt.Println(nodo.Izquierda.User.ID)
		nodo = nodo.Izquierda
		RecorridoPre(nodo)
	} else if nodo.Derecha != nil {
		fmt.Println(nodo.Derecha.User.ID)
		nodo = nodo.Derecha
		RecorridoPre(nodo)
	}
}

func RecorridoPost(nodo *Nodo) {
	fmt.Println("post")
	if nodo == nil {
		fmt.Println("No existe arbol")
	}
	if nodo.Derecha != nil {
		fmt.Println(nodo.Derecha.User.ID)
		nodo = nodo.Derecha
		RecorridoPost(nodo)
	} else if nodo.Izquierda != nil {
		fmt.Println(nodo.Izquierda.User.ID)
		nodo = nodo.Izquierda
		RecorridoPost(nodo)
	}
}

func RecorrerInord(nodo *Nodo) {
	fmt.Println("inord")
	if nodo == nil {
		fmt.Println("No existe arbol")
	}
	if nodo.Izquierda != nil {
		fmt.Println(nodo.Izquierda.User.ID)
		nodo = nodo.Izquierda
		RecorrerInord(nodo)
	} else if nodo.Derecha != nil {
		fmt.Println(nodo.Derecha.User.ID)
		nodo = nodo.Derecha
		RecorrerInord(nodo)
	}
}
