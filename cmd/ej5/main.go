package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Juego interface {
	AnotarPuntoJugador1()
	AnotarPuntoJugador2()
	ObtenerPuntuacion() string
}

type Tennis struct {
	jugador1 Jugador
	jugador2 Jugador
}

func NewTennis(nombreJugador1, nombreJugador2 string) *Tennis {
	return &Tennis{
		jugador1: Jugador{Nombre: nombreJugador1},
		jugador2: Jugador{Nombre: nombreJugador2},
	}
}

func (receiver *Tennis) AnotarPuntoJugador1() {
	receiver.jugador1.Puntaje++
}

func (receiver *Tennis) AnotarPuntoJugador2() {
	receiver.jugador2.Puntaje++
}

func (receiver Tennis) ObtenerPuntuacion() string {
	return fmt.Sprintf("%v\n%v\n=======\n", receiver.jugador1, receiver.jugador2)
}

var _ Juego = (*Tennis)(nil)

type Jugador struct {
	Nombre  string
	Puntaje int
}

func NewJugador(nombre string, puntaje int) *Jugador {
	return &Jugador{Nombre: nombre, Puntaje: puntaje}
}

func (receiver Jugador) String() string {
	return fmt.Sprintf("%v: %v", receiver.Nombre, parsearPuntaje(receiver.Puntaje))
}

func main() {
	contenido, err := ioutil.ReadFile("jugadores.txt")
	if err != nil {
		panic(err)
	}
	jugadores := strings.Split(string(contenido), ",")
	juego := NewTennis(jugadores[0], jugadores[1])
	var resultados []string
	juego.AnotarPuntoJugador1()
	resultados = append(resultados, juego.ObtenerPuntuacion())
	juego.AnotarPuntoJugador1()
	resultados = append(resultados, juego.ObtenerPuntuacion())
	juego.AnotarPuntoJugador2()
	resultados = append(resultados, juego.ObtenerPuntuacion())
	print := fmt.Sprintf("%v", resultados)
	fmt.Println(print)
	err = ioutil.WriteFile("result.txt", []byte(print), 0644)
	if err != nil {
		panic(err)
	}
}

func parsearPuntaje(valor int) string {
	switch valor {
	case 1:
		return "15"
	case 2:
		return "30"
	case 3:
		return "40"
	default:
		return "0"
	}
}
