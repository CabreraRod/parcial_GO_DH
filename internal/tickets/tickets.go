package tickets

import (
	"fmt"
	"os"
	"encoding/csv"
	"strings"
	"log"
	"strconv"
)

type Ticket struct {
	id int
	nombre string
	email string
	pais_destino string
	hora_vuelo string
	precio int
}

// ejemplo 1
// Este metodo recibe por parametros un destino y nos indica cuantos tickets de vuelo hay vendidos hacia ese destination

func GetTotalTickets(destination string) (int, error) {
	aux := 0
	file, err := os.Open("./tickets.csv")
	if(err != nil){
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	allTickets, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	//var tickets []Ticket
	for _, ticket := range allTickets {
		pais_destino := ticket[3]
		if pais_destino == destination {
			aux++
		}
	}
	return aux, nil

}

// ejemplo 2
// Este metodo recibe un momento del dia, puede ser (madruagada 0-6; mañana 7-12; tarde 13-19; noche 20-23)
// nos permite saber en que momento del dia se realizan mas vuelos, siendo este actualmente en la "madrugada"
func GetCountByPeriod(time string) (int, error) {
	file, err := os.Open("./tickets.csv")
	if(err != nil){
		log.Fatalf("error al abrir el archivo de tickects.csv")
	}
	defer file.Close()

	reader := csv.NewReader(file)

	allTickets, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	aux := 0
	for _, ticket := range allTickets {
		hour := ticket[4]
		splitHour := strings.Split(hour, ":")
		intHour, err := strconv.Atoi(splitHour[0])
		if err != nil{
			log.Printf("error convirtiendo hora(string) a int: %s", err)
			continue
		}

		switch {
			case intHour <= 6:
				if time == "madrugada" {
					aux++
				}
			case  intHour <= 12:
				if time == "mañana" {
					aux++
				}
		
			case intHour <= 19:
				if time == "tarde" {
					aux++
				}
		
			case intHour <= 23:
				if time == "noche" {
					aux++
				}
		
			default:
				fmt.Printf("Unexpected hour value: %d\n", intHour)
	}
}
	return aux, nil
}

// ejemplo 3
// Este metodo recibira un destination, que puede ser cualquier pais a donde volemos 
// y un total de tickets cargados en la BD. Actualmente 11/03/23 poseemos 1000 tickets de vuelo
// este metodo nos permite saber que paises reciben mas % de vuelos.

func AverageDestination(destination string, total int) (float64, error) {
	allTickets, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	promedio := (float64(allTickets)/float64(total))
	return promedio, err
}
