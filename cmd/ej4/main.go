package main

import "fmt"

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
	juego := NewTennis("Alice", "Bob")
	juego.AnotarPuntoJugador1()
	fmt.Println(juego.ObtenerPuntuacion())
	juego.AnotarPuntoJugador1()
	fmt.Println(juego.ObtenerPuntuacion())
	juego.AnotarPuntoJugador2()
	fmt.Println(juego.ObtenerPuntuacion())
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
