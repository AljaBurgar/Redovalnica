package main

import (
	"fmt"

	"github.com/AljaBurgar/Redovalnica/redovalnica"
)

func main() {

	studenti["63210001"] = Student{"Ana", "Novak", []int{10, 9, 8}}
	studenti["63210002"] = Student{"Boris", "Kralj", []int{6, 7, 5, 8, 6, 7, 5, 8}}
	studenti["63210003"] = Student{"Janez", "Novak", []int{4, 5, 3}}

	redovalnica.DodajOceno(studenti, "63210004", 5) //Izpise da studenta s to vpisno stevilko ni na seznamu
	redovalnica.DodajOceno(studenti, "63210001", 17)
	fmt.Println(studenti["63210001"]) //Ocena se ni dodala ker ni med 0 in 10
	redovalnica.DodajOceno(studenti, "63210001", 10)
	redovalnica.DodajOceno(studenti, "63210001", 9)
	redovalnica.DodajOceno(studenti, "63210001", 8)
	fmt.Println(studenti["63210001"]) //Ocene so se dodale

	fmt.Println()

	redovalnica.IzpisRedovalnice(studenti)

	fmt.Println()

	redovalnica.IzpisiKoncniUspeh(studenti)
	//fmt.Println(studenti)
}
