package main

import (
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"fmt"
	"log"
)
const (
	destination = "Canada"
)
func getTotalTickets() {
	total, err := tickets.GetTotalTickets(destination)
	if err != nil {
		log.Printf("Fallo al obtener el total de tickets: %v", err)
		return
	}
	fmt.Printf("Total de tickets vendido a %s: %d\n", destination, total)
}


func getCountByPeriod() {
	count, err := tickets.GetCountByPeriod("noche")
	if err != nil {
		log.Printf("Fallo al obtener la cantidad de viajeros viajando en este periodo: %v", err)
		return
	}
	fmt.Printf("Total de pasajeros viajando de noche: %d\n", count)
}

func getAverageDestination() {
	const totalTickets = 1000
	promedio, err := tickets.AverageDestination(destination, totalTickets)
	if err != nil {
		log.Printf("Fallo al obetener el promedio de viajes: %v", err)
		return
	}
	fmt.Printf("Promedio de tickets a %s: %.2f\n", destination, promedio)
}

func main() {
	getTotalTickets()
	getCountByPeriod()
	getAverageDestination()
}
