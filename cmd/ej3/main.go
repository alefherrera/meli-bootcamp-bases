package main

import "fmt"

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
	jugador1 := NewJugador("Alice", 2)
	jugador2 := NewJugador("Bob", 1)
	fmt.Printf("%v\n%v", jugador1, jugador2)
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
