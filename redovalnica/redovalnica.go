package redovalnica

import "fmt"

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študenta ni na seznamu")
		return
	}

	if ocena >= 1 && ocena <= 10 {
		student.Ocene = append(student.Ocene, ocena)
	}

	studenti[vpisnaStevilka] = student
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1.0
	}

	var average float64
	dolzina := len(student.Ocene)

	if dolzina < 6 {
		return 0.0
	}

	for _, ocena := range student.Ocene {
		average += float64(ocena)
	}

	average = average / float64(dolzina)
	return average
}
func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA")
	for vpisnaStevilka, student := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisnaStevilka, student.Ime, student.Priimek, student.Ocene)
	}
}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisnaStevilka, student := range studenti {
		povprecnaOcena := povprecje(studenti, vpisnaStevilka)
		if povprecnaOcena >= 9 {
			fmt.Printf("%s %s: povprečna ocena %f -> Odličen študent!\n", student.Ime, student.Priimek, povprecnaOcena)
		} else if povprecnaOcena >= 6 && povprecnaOcena < 9 {
			fmt.Printf("%s %s: povprečna ocena %f -> Povprečen študent\n", student.Ime, student.Priimek, povprecnaOcena)
		} else {
			fmt.Printf("%s %s: povprečna ocena %f -> Neuspešen študent\n", student.Ime, student.Priimek, povprecnaOcena)
		}
	}
}
