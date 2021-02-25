package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type nodo struct {
	nombre, apellido, apodo, favoritos string
	Siguiente, Anterior                *nodo
}

type lista struct {
	cabeza *nodo
	cola   *nodo
}

func (this *lista) Insertar(nuevo *nodo) {
	if this.cabeza == nil {
		this.cabeza = nuevo
		this.cola = nuevo
	} else {
		this.cola.Siguiente = nuevo
		nuevo.Anterior = this.cola
		this.cola = nuevo
	}
}

func main() {
	li := lista{nil, nil}
	a := nodo{"Marvin", "Martinez", "Marvin25ronal", "Jugar apex", nil, nil}
	b := nodo{"Yaiza", "Pineda", "Bambi", "Patinar", nil, nil}
	c := nodo{"Jonathan", "Lopez", "Pancho", "Comer", nil, nil}
	d := nodo{"usuario1", "bla", "bla", "Jugar apex", nil, nil}
	e := nodo{"usuario2", "bla", "bla", "Jugar apex", nil, nil}
	f := nodo{"usuario3", "sale edd", "vamos con todo", "100 en la fase 1", nil, nil}
	li.Insertar(&a)
	li.Insertar(&b)
	li.Insertar(&c)
	li.Insertar(&d)
	li.Insertar(&e)
	li.Insertar(&f)
	archivo, _ := os.Create("graficoLinealizado.dot")
	_, _ = archivo.WriteString("digraph grafico{" + "\n")
	_, _ = archivo.WriteString("compound=true;" + "\n")
	aux := li.cabeza
	var i int = 0
	for aux != nil {
		_, _ = archivo.WriteString("subgraph cluster0{" + "\n")
		_, _ = archivo.WriteString("edge[minlen=0.1, dir=fordware]" + "\n")
		_, _ = archivo.WriteString("struct" + strconv.Itoa(i) + "[shape=record,label=\"" + aux.nombre + "|" + aux.apellido + "|{ " + aux.apodo + " | " + aux.favoritos + "}\"];" + "\n")
		if i != 0 {
			_, _ = archivo.WriteString("struct" + strconv.Itoa(i-1) + " -> struct" + strconv.Itoa(i) + ";" + "\n")
		}
		_, _ = archivo.WriteString("}" + "\n")
		aux = aux.Siguiente
		i++
	}
	_, _ = archivo.WriteString("}" + "\n")
	archivo.Close()
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tsvg", "./graficoLinealizado.dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("grafica.svg", cmd, os.FileMode(mode))
}
