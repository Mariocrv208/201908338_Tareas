package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Struct Mensaje
type Mensajes struct {
	Mensajes []Mensaje `json:"Mensajes"`
}

type Mensaje struct {
	Origen  string `json:Origen`
	Destino string `json:Destino`
	Msg     []Msg  `json:"Msg"`
}

type Msg struct {
	Fecha string `json:"Fecha"`
	Texto string `json:"Texto"`
}

//funciones

func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Si jala we uwu")
}

//variables
var mandar Mensajes

func agregar(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No inserto we :c")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &mandar)
	json.NewEncoder(w).Encode(mandar)
}

func Mostrar(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(mandar.Mensajes); i++ {
		fmt.Println("Origen: " + mandar.Mensajes[i].Origen)
		fmt.Println("Destino: " + mandar.Mensajes[i].Destino)
		for j := 0; j < len(mandar.Mensajes[i].Msg); j++ {
			fmt.Println("Fecha: " + mandar.Mensajes[i].Msg[j].Fecha)
			fmt.Println("Texto: " + mandar.Mensajes[i].Msg[j].Texto)
		}
	}
	json.NewEncoder(w).Encode(mandar)
}

//main

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/agregar", agregar).Methods("POST")
	router.HandleFunc("/Mostrar", Mostrar).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
