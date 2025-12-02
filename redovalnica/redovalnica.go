// Paket omogoča dodajanje ocen študentom, računanje povprečja, izpis redovalnice in izpis končnega uspeha študenta.
package redovalnica

import "fmt"

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// Funkcija DodajOceno sprejme slovar študentov, vpisno številko študenta in oceno, ki jo je prejel.
// Če vpisna številka ne pripada nobenemu študentu funkcija ne vrne nič.
// Če je ocena ustrezna (med 1 in 10) jo doda učencu.
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

// Funkcija povprecje sprejme slovar študentov in vpisno številko študenta.
// Če vpisna številka ne pripada nobenemu študentu, vrne -1.0.
// Če ima študent ocen manj kot 6, funkcija vrne 0.
// Drugače vrne povprečje študentovih ocen.
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

// Funkcija IzpisRedovalnice sprejme slovar študentov.
// Za vsakega študenta izpiše vpisno številko, njegovo ime in priimek in njegove ocene.
func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA")
	for vpisnaStevilka, student := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisnaStevilka, student.Ime, student.Priimek, student.Ocene)
	}
}

// Funkcija IzpisKoncniUspeh sprejme slovar študentov.
// Sprehodi se po študentih in izračuna njihovo povprečje.
// Če je povprečna ocena večja od 9 izpiše, da je študent odličen.
// Če je povprečna ocena med 6 in 9 izpiše, da je študent povprečen.
// Če je povprečna ocena manjša od 6 izpiše, da je študent neuspešen.
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
