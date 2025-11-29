package main

import (
	"fmt"
)

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func main() {
	var studenti map[string]Student
	studenti = make(map[string]Student)

	studenti["63210001"] = Student{"Ana", "Novak", []int{10, 9, 8}}
	studenti["63210002"] = Student{"Boris", "Kralj", []int{6, 7, 5, 8, 6, 7, 5, 8}}
	studenti["63210003"] = Student{"Janez", "Novak", []int{4, 5, 3}}

	dodajOceno(studenti, "63210004", 5) //Izpise da studenta s to vpisno stevilko ni na seznamu
	dodajOceno(studenti, "63210001", 17)
	fmt.Println(studenti["63210001"]) //Ocena se ni dodala ker ni med 0 in 10
	dodajOceno(studenti, "63210001", 10)
	dodajOceno(studenti, "63210001", 9)
	dodajOceno(studenti, "63210001", 8)
	fmt.Println(studenti["63210001"]) //Ocene so se dodale

	fmt.Println()

	povprecno := povprecje(studenti, "63210001")
	fmt.Println(povprecno) //Izpis povprecja za Ano Novak
	povprecno = povprecje(studenti, "63210004")
	fmt.Println(povprecno) //Povprecje je -1 ker studenta s to vpisno stevilko ni na seznamu
	povprecno = povprecje(studenti, "63210003")
	fmt.Println(povprecno) //Povprecje je 0 ker ima oddanih manj kot 6 nalog

	fmt.Println()

	izpisRedovalnice(studenti)

	fmt.Println()

	izpisiKoncniUspeh(studenti)
	//fmt.Println(studenti)
}
