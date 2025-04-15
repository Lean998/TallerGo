package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

//"errors"
//"go_parte_2/matematica"
//"strings"

func main() {
	var listDispositivos []Dispositivo
	listDispositivos = append(listDispositivos, Dispositivo{Nombre: "Lampara", Estado: false})
	listDispositivos = append(listDispositivos, Dispositivo{Nombre: "Radio", Estado: true})
	listDispositivos = append(listDispositivos, Dispositivo{Nombre: "Televisor", Estado: false})
	listDispositivos = append(listDispositivos, Dispositivo{Nombre: "Cocina", Estado: true})
	listDispositivos = append(listDispositivos, Dispositivo{Nombre: "Telefono", Estado: false})

	for i := 0; i < len(listDispositivos); i++ {
		fmt.Print("Estado Actual: ")
		listDispositivos[i].getEstado()
		if i/2 == 0 {
			err := listDispositivos[i].Encender()
			if err != nil {
				fmt.Println("Error: ", err.Error())
			}
		} else {
			err := listDispositivos[i].Apagar()
			if err != nil {
				fmt.Println("Error: ", err.Error())
			}
		}
		fmt.Print("Nuevo estado: ")
		listDispositivos[i].getEstado()
	}
}

type Dispositivo struct {
	Nombre string
	Estado bool
}

func (d Dispositivo) getEstado() {
	if d.Estado {
		color.Green(d.EstadoActual())
	}
	color.Red(d.EstadoActual())
}

func (d Dispositivo) EstadoActual() string {
	if d.Estado {
		return "Encendido"
	}
	return "Apagado"
}

func (d *Dispositivo) Encender() error {
	if d.Estado {
		return errors.New("el dispositivo ya esta encendido")
	}
	d.Estado = true
	return nil
}

func (d *Dispositivo) Apagar() error {
	if !d.Estado {
		return errors.New("el dispositivo ya esta apagado")
	}
	d.Estado = false
	return nil
}

type Controlabe interface {
	Encender() error
	Apagar() error
	EstadoActual() string
}

/*

func main() {

	litString := "Hola"
	//var puntString = &litString

	modString(&litString)
	fmt.Println(litString)

	err := checkStr("")
	if err != nil {
		fmt.Println("Error", err)
	}
	//fmt.Println(litString, puntString, *puntString)

	persona := Persona{
		Nombre:   "Leandro",
		Apellido: "Aguero",
		Edad:     20,
	}
	var presentable Presentable
	presentable = &persona
	presentable.Presentarse()
	fmt.Println(persona)

}

type Persona struct {
	Nombre   string //publico
	Apellido string
	Edad     int
}

func (p *Persona) Presentarse() string {
	return p.Nombre + " " + p.Apellido
}

type Presentable interface {
	Presentarse() string
}

func (p Persona) getNombre() string {
	return (p.Nombre)
}

func (p *Persona) setNombre(nombre string) {
	p.Nombre = nombre
}

func modString(s *string) {
	*s = "Algo"
}

func checkStr(s string) error {
	if s == "" {
		return errors.New("string vacio")
	}
	return nil
}

*/
